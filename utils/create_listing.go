package utils

import (
	"math/rand/v2"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func CreateListing(page *rod.Page, i Item, listingIndex int) {

	page = page.MustNavigate("https://www.facebook.com/marketplace/create/item").MustWaitLoad()
	Logger.Info("NAVIGATED TO CREATE ITEM PAGE")
	if page.MustHas(`div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xyamay9.x1l90r2v`) {
		elements := page.MustElements(`div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xyamay9.x1l90r2v`)

		for _, element := range elements {
			if strings.Contains(element.MustText(), "Limit reached") {
				Logger.Warn("LIMIT REACHED")
				Logger.Warn("YOU ARE NOT ABLE TO CREATE NEW LISTINGS AT THE MOMENT. WE LIMIT HOW OFTEN SELLERS CAN POST IN AN EFFORT TO BUILD A BETTER COMMUNITY FOR EVERYONE.")
				panic("Limit reached")
			}
		}
	}
	// insert images
	page.MustElement(`input[type="file"][accept="image/*,image/heif,image/heic"]`).MustSetFiles(i.Images...)
	Logger.Info("IMAGES UPLOADED", "COUNT", len(i.Images))
	page.MustElement(`div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x193iq5w.x1l7klhg.x1iyjqo2.xs83m0k.x2lwn1j.xyamay9`).MustClick()

	labels := page.MustElements(`label.x78zum5.xh8yej3`)
	for _, label := range labels {
		text := strings.ToLower(strings.TrimSpace(label.MustText()))
		switch text {
		case "title":
			label.MustInput(i.Title)
			Logger.Info("TITLE INSERTED", "TITLE", i.Title)
		case "price":
			label.MustInput(i.Price)
			Logger.Info("PRICE INSERTED", "PRICE", i.Price)
		case "category":
			selectDrawerOption(page, label, i.Category, "CATEGORY", []string{
				`div[role="button"][tabindex="0"]`,
				`div[role="button"]`,
			})
			Logger.Info("CATEGORY INSERTED", "CATEGORY", i.Category)
		case "condition":
			selectDrawerOption(page, label, i.Condition, "CONDITION", []string{
				`div[role="option"]`,
				`div[role="button"][tabindex="0"]`,
				`div[role="button"]`,
			})
			Logger.Info("CONDITION INSERTED", "CONDITION", i.Condition)
		case "description":
			label.MustInput(i.Description)
		case "product tags":
			allTags := len(i.Tags)
			if allTags > 0 {
				for _, tag := range i.Tags {
					tag = strings.TrimSpace(tag)
					label.MustInput(tag)
					err := page.Keyboard.Press(input.Enter)
					if err != nil {
						Logger.Error("FAILED TO PRESS ENTER FOR TAG", "TAG", tag, "ERR", err)
					}
				}

			}
			Logger.Info("TAGS INSERTED", "COUNT", allTags)
		}
	}

	if strings.Contains(page.MustHTML(), "Please upload at least one photo.") {
		Logger.Warn("PLEASE UPLOAD AT LEAST ONE PHOTO.")
		Logger.Warn("SOMETHING ISN'T WORKING. THIS MAY BE BECAUSE OF A TECHNICAL ERROR WE'RE WORKING TO FIX.")
		panic("Something isn't working. This may be because of a technical error we're working to fix.")
	}

	// go to next page
	page.MustElement(`div[aria-label="Next"]`).MustClick()
	Logger.Info("CLICKED NEXT BUTTON. WAITING FOR AUDIENCE STEP...")
	waitForNextPage(page, 120)

	// select suggested groups
	groups := page.MustElements(`div[role="checkbox"]`)
	Logger.Info("FOUND AUDIENCE GROUPS", "COUNT", len(groups))

	startFromBottom := listingIndex%2 != 0
	selectionDirection := "TOP_TO_BOTTOM"
	if startFromBottom {
		selectionDirection = "BOTTOM_TO_TOP"
	}
	Logger.Info("SELECTING AUDIENCE GROUPS", "DIRECTION", selectionDirection)

	tickedGroups := 0
	for _, groupIndex := range groupSelectionIndexes(len(groups), startFromBottom) {
		groups[groupIndex].MustClick()
		time.Sleep(100 * time.Millisecond)
		tickedGroups++
	}
	Logger.Info("SELECTED AUDIENCE GROUPS", "COUNT", tickedGroups)
	waitForPostButton(page)

	waitForPublishing(page, i.Title)

	Logger.Info("PUBLISHED ITEM", "TITLE", i.Title)
}

func groupSelectionIndexes(groupCount int, startFromBottom bool) []int {
	if groupCount <= 0 {
		return nil
	}

	indexes := make([]int, 0, groupCount)
	if startFromBottom {
		for i := groupCount - 1; i >= 0; i-- {
			indexes = append(indexes, i)
		}
		return indexes
	}

	for i := range groupCount {
		indexes = append(indexes, i)
	}

	return indexes
}

func waitForNextPage(page *rod.Page, timeoutSeconds int) {
	maxRetries := timeoutSeconds / 3 // Convert timeout to retry count

	for i := range maxRetries {
		currentURL := page.MustInfo().URL

		if strings.Contains(currentURL, "/marketplace/create/item?step=audience") {
			Logger.Info("SUCCESSFULLY REACHED AUDIENCE STEP")
			time.Sleep(3 * time.Second)
			return // Exit function immediately
		}

		Logger.Info("WAITING FOR NEXT PAGE", "RETRY", i+1)
		time.Sleep(3 * time.Second)
	}

	Logger.Warn("TIMEOUT REACHED WHILE WAITING FOR THE AUDIENCE STEP")
}

func waitForPostButton(page *rod.Page) {
	maxRetries := 60 // Max retries (2 seconds per retry → 120 seconds)
	var postButton *rod.Element

	for range maxRetries {
		// Re-fetch the button in each iteration
		postButton = page.MustElement(`div[aria-label="Publish"]`)
		if postButton == nil {
			Logger.Info("POST BUTTON NOT FOUND, RETRYING")
			time.Sleep(2 * time.Second)
			continue
		}

		// Check if the button is enabled
		disabled, err := postButton.Attribute("aria-disabled")
		if err == nil && disabled == nil {
			Logger.Info("POST BUTTON IS READY, CLICKING")
			postButton.MustClick()
			return
		}

		Logger.Info("POST BUTTON IS NOT READY YET, RETRYING")
		time.Sleep(2 * time.Second)
	}

	Logger.Warn("TIMEOUT REACHED WHILE WAITING FOR THE POST BUTTON")
}

func waitForPublishing(page *rod.Page, title string) {
	maxRetries := 60 // (2 seconds per retry → 120 seconds total)

	for i := range maxRetries {
		currentURL := page.MustInfo().URL

		// First, check if the item was published
		if strings.Contains(currentURL, "marketplace/you/selling") {
			Logger.Info("ITEM PUBLISHED SUCCESSFULLY", "TITLE", title)
			return
		}

		// Check for the "Something isn't working" error
		somethingIsntWorking := strings.Contains(page.MustHTML(), "Something isn't working. This may be because of a technical error we're working to fix.")
		if somethingIsntWorking {
			panic("Something isn't working. This may be because of a technical error we're working to fix.")
		}

		// If neither condition is met, continue waiting
		Logger.Info("WAITING FOR PUBLISHING TO COMPLETE", "TITLE", title, "RETRY", i+1, "MAX_RETRIES", maxRetries)
		time.Sleep(2 * time.Second)
	}

	Logger.Warn("TIMEOUT REACHED WHILE PUBLISHING ITEM")
}

func selectDrawerOption(page *rod.Page, trigger *rod.Element, value string, fieldName string, selectors []string) {
	normalizedValue := strings.ToLower(strings.TrimSpace(value))
	trigger.MustClick()
	time.Sleep(1 * time.Second)

	for attempt := 0; attempt < 5; attempt++ {
		for _, selector := range selectors {
			if clickMatchingOption(page, selector, normalizedValue, value, fieldName) {
				return
			}
		}

		Logger.Info("WAITING FOR DRAWER OPTIONS", "FIELD", fieldName, "TARGET", value, "ATTEMPT", attempt+1)
		time.Sleep(500 * time.Millisecond)
	}

	panic("failed to select " + fieldName + ": " + value)
}

func clickMatchingOption(page *rod.Page, selector string, normalizedValue string, rawValue string, fieldName string) bool {
	options := page.MustElements(selector)
	if len(options) == 0 {
		return false
	}

	for _, option := range options {
		optionText := normalizeOptionText(option.MustText())
		if optionText == "" {
			continue
		}

		if optionText == normalizedValue {
			Logger.Info("MATCHED DRAWER OPTION", "FIELD", fieldName, "TARGET", rawValue, "SELECTOR", selector, "OPTION_TEXT", option.MustText())
			option.MustClick()
			time.Sleep(500 * time.Millisecond)
			return true
		}
	}

	for _, option := range options {
		optionText := normalizeOptionText(option.MustText())
		if optionText == "" || !strings.Contains(optionText, normalizedValue) {
			continue
		}

		Logger.Info("MATCHED DRAWER OPTION", "FIELD", fieldName, "TARGET", rawValue, "SELECTOR", selector, "OPTION_TEXT", option.MustText())
		option.MustClick()
		time.Sleep(500 * time.Millisecond)
		return true
	}

	return false
}

func normalizeOptionText(text string) string {
	text = strings.ToLower(strings.TrimSpace(text))
	if text == "" {
		return ""
	}

	lines := strings.Split(text, "\n")
	primary := strings.TrimSpace(lines[0])
	if primary != "" {
		return primary
	}

	return text
}

func returnRandomFloat(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return min + rand.Float64()*(max-min)
}
