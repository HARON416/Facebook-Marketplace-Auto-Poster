package main

import (
	"bufio"
	"context"
	"facebook-marketplace-listing-tool/utils"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	browser, page := utils.Login()
	defer browser.MustClose()
	defer page.MustClose()

	itemsPath := "/home/kibet/Pictures/FACEBOOK/PHONES/MARKETPLACE"

	utils.Logger.Info("LOOKING FOR ITEMS", "PATH", itemsPath)

	items, err := utils.GetItems(itemsPath)
	if err != nil {
		utils.Logger.Fatal("FAILED TO LOAD ITEMS", "PATH", itemsPath, "ERR", err)
	}

	utils.Logger.Info("LOADED ITEMS", "COUNT", len(items))

	ctx := context.Background()
	rewriter, err := utils.NewGeminiDescriptionRewriter(ctx)
	if err != nil {
		utils.Logger.Warn("FAILED TO INITIALIZE GEMINI DESCRIPTION REWRITER", "ERR", err)
		utils.Logger.Warn("PRESS ENTER TO CONTINUE WITHOUT GEMINI REWRITES OR PRESS CTRL+C TO EXIT")
		if _, readErr := bufio.NewReader(os.Stdin).ReadString('\n'); readErr != nil {
			utils.Logger.Warn("FAILED TO READ CONTINUE CONFIRMATION", "ERR", readErr)
		}
	}

	startedAt := time.Now()
	createdListings := 0
	defer func() {
		duration := time.Since(startedAt).Round(time.Second)
		if recovered := recover(); recovered != nil {
			utils.Logger.Error(
				fmt.Sprintf("TOOK %.2f MINUTES TO CREATE %d LISTINGS BEFORE FAILURE", duration.Minutes(), createdListings),
				"COUNT", createdListings,
				"DURATION", duration,
				"ERR", recovered,
			)
			panic(recovered)
		}

		utils.Logger.Info(
			fmt.Sprintf("TOOK %.2f MINUTES TO CREATE %d LISTINGS", duration.Minutes(), createdListings),
			"COUNT", createdListings,
			"DURATION", duration,
		)
	}()

	for i, item := range items {
		utils.Logger.Info("PREPARING LISTING",
			"INDEX", i+1,
			"TOTAL", len(items),
			"TITLE", item.Title,
			"PRICE", item.Price,
			"CATEGORY", item.Category,
			"CONDITION", item.Condition,
			"TAGS", item.Tags,
			"IMAGES", len(item.Images),
		)

		description := item.Description
		if rewriter != nil {
			rewrittenDescription, err := rewriter.Rewrite(ctx, item.Description)
			if err != nil {
				utils.Logger.Warn("GEMINI REWRITE FAILED, USING ORIGINAL DESCRIPTION",
					"TITLE", item.Title,
					"ERR", err,
				)
			} else {
				description = rewrittenDescription
				utils.Logger.Info("GEMINI DESCRIPTION REWRITTEN",
					"TITLE", item.Title,
					"DESCRIPTION", description,
				)
			}
		} else {
			utils.Logger.Warn("SKIPPING GEMINI REWRITE, USING ORIGINAL DESCRIPTION", "TITLE", item.Title)
		}

		item.Description = description

		utils.Logger.Info("CREATING LISTING", "TITLE", item.Title)
		utils.CreateListing(page, item, i)
		createdListings++
		utils.Logger.Info("LISTING CREATED", "TITLE", item.Title)

		if i == len(items)-1 {
			continue
		}

		// Sleep for a random duration between 5 and 15 seconds to mimic human behavior
		duration := time.Duration(rand.IntN(10)+5) * time.Second
		utils.Logger.Info("SLEEPING BEFORE NEXT ITEM", "DURATION", duration)
		time.Sleep(duration)
	}
}
