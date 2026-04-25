package utils

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func PostItemsToMarketplace(browser *rod.Browser, page *rod.Page, items []Item) error {
	defer browser.MustClose()

	for _, i := range items {
		page = page.MustNavigate("https://www.facebook.com/marketplace/create/item").MustWaitLoad()
		log.Info("NAVIGATED TO CREATE ITEM PAGE")
		if page.MustHas(`div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xyamay9.x1l90r2v`) {
			elements := page.MustElements(`div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x2lah0s.x193iq5w.xyamay9.x1l90r2v`)

			for _, element := range elements {
				if strings.Contains(element.MustText(), "Limit reached") {
					log.Warn("Limit reached")
					log.Warn("You are not able to create new listings at the moment. We limit how often sellers can post in an effort to build a better community for everyone.")
					os.Exit(1)
				}
			}
		}
		// insert images
		page.MustElement(`input[type="file"][accept="image/*,image/heif,image/heic"]`).MustSetFiles(i.Images...)
		log.Info("IMAGES INSERTED")
		page.MustElement(`div.x1n2onr6.x1ja2u2z.x9f619.x78zum5.xdt5ytf.x193iq5w.x1l7klhg.x1iyjqo2.xs83m0k.x2lwn1j.xyamay9`).MustClick()

		labels := page.MustElements(`label.x78zum5.xh8yej3`)
		for _, label := range labels {
			text := strings.ToLower(strings.TrimSpace(label.MustText()))
			switch {
			case text == "title":
				label.MustInput(i.Title)
				log.Info("TITLE INSERTED")
			case text == "price":
				label.MustInput(i.Price)
				log.Info("PRICE INSERTED")
			case text == "category":
				label.MustClick()
				time.Sleep(1 * time.Second)
				options := page.MustElements(`div[data-visualcompletion="ignore-dynamic"]`)
				for _, option := range options {
					if strings.Contains(strings.ToLower(option.MustText()), i.Category) {
						option.MustClick()
						break
					}
				}
				log.Info("CATEGORY INSERTED")
			case text == "condition":
				label.MustClick()
				time.Sleep(1 * time.Second)
				options := page.MustElements(`div[role="option"]`)
				for _, option := range options {
					if strings.ToLower(option.MustText()) == i.Condition {
						option.MustClick()
						break
					}
				}
				log.Info("CONDITION INSERTED")
			case text == "description":
				label.MustInput(i.Description)
			case text == "product tags":
				allTags := len(i.Tags)
				if allTags > 0 {
					for _, tag := range i.Tags {
						tag = strings.TrimSpace(tag)
						label.MustInput(tag)
						err := page.Keyboard.Press(input.Enter)
						if err != nil {
							return fmt.Errorf("ERROR PRESSING ENTER KEY: %v", err)
						}
					}

				}
				log.Info("TAGS INSERTED")
			}
		}

		if strings.Contains(page.MustHTML(), "Please upload at least one photo.") {
			log.Warn("Please upload at least one photo.")
			log.Warn("Something isn't working. This may be because of a technical error we're working to fix.")
			os.Exit(1)
		}

		// go to next page
		page.MustElement(`div[aria-label="Next"]`).MustClick()
		log.Info("CLICKED ON THE NEXT BUTTON")
		waitForNextPage(page, 120)
		// select suggested groups
		groups := page.MustElements(`div[role="checkbox"]`)
		log.Infof("ALL GROUPS: %v", len(groups))
		tickedGroups := 0
		for _, group := range groups {
			group.MustClick()
			time.Sleep(100 * time.Millisecond)
			tickedGroups++
		}
		log.Infof("TICKED GROUPS: %v", tickedGroups)
		waitForPostButton(page)
		waitForPublishing(page, i.Title)
		log.Infof("PUBLISHED ITEM: %s", i.Title)
		// wait for 30-60 seconds before posting the next item
		waitTime := returnRandomFloat(30, 60)
		log.Infof("WAITING FOR %.2f SECONDS BEFORE POSTING THE NEXT ITEM", waitTime)
		time.Sleep(time.Duration(waitTime) * time.Second)
		log.Info("WAITING COMPLETED, MOVING TO THE NEXT ITEM")
	}

	return nil
}

func waitForNextPage(page *rod.Page, timeoutSeconds int) {
	maxRetries := timeoutSeconds / 3 // Convert timeout to retry count

	for i := range maxRetries {
		currentURL := page.MustInfo().URL

		if strings.Contains(currentURL, "/marketplace/create/item?step=audience") {
			log.Info("SUCCESSFULLY REACHED AUDIENCE STEP")
			time.Sleep(3 * time.Second)
			return // Exit function immediately
		}

		log.Infof("GOING TO NEXT PAGE, retry: %d", i+1)
		time.Sleep(3 * time.Second)
	}

	log.Warn("TIMEOUT REACHED! PAGE DID NOT LOAD THE EXPECTED URL.")
}

func waitForPostButton(page *rod.Page) {
	maxRetries := 60 // Max retries (2 seconds per retry → 120 seconds)
	var postButton *rod.Element

	for range maxRetries {
		// Re-fetch the button in each iteration
		postButton = page.MustElement(`div[aria-label="Publish"]`)
		if postButton == nil {
			log.Info("POST BUTTON NOT FOUND! RETRYING...")
			time.Sleep(2 * time.Second)
			continue
		}

		// Check if the button is enabled
		disabled, err := postButton.Attribute("aria-disabled")
		if err == nil && disabled == nil {
			log.Info("POST BUTTON IS READY! CLICKING...")
			postButton.MustClick()
			return
		}

		log.Info("POST BUTTON IS NOT YET READY! RETRYING...")
		time.Sleep(2 * time.Second)
	}

	log.Warn("TIMEOUT REACHED! POST BUTTON DID NOT BECOME READY.")
}

func waitForPublishing(page *rod.Page, title string) {
	maxRetries := 60 // (2 seconds per retry → 120 seconds total)

	for i := range maxRetries {
		currentURL := page.MustInfo().URL

		// First, check if the item was published
		if strings.Contains(currentURL, "marketplace/you/selling") {
			log.Infof("✅ %s PUBLISHED SUCCESSFULLY!", title)
			return
		}

		// Check for the "Something isn't working" error
		somethingIsntWorking := strings.Contains(page.MustHTML(), "Something isn't working. This may be because of a technical error we're working to fix.")
		if somethingIsntWorking {
			log.Error("Something isn't working. This may be because of a technical error we're working to fix.")
			os.Exit(1)
		}

		// If neither condition is met, continue waiting
		log.Infof("⏳ PUBLISHING %s... (%d/%d RETRIES)", title, i+1, maxRetries)
		time.Sleep(2 * time.Second)
	}

	log.Warn("⏰ TIMEOUT! PUBLISHING TOOK TOO LONG.")
}

func returnRandomFloat(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return min + rand.Float64()*(max-min)
}
