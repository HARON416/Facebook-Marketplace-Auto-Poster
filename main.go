package main

import (
	"context"
	"facebook-marketplace-listing-tool/utils"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	browser, page := utils.Login()
	defer browser.MustClose()
	defer page.MustClose()

	itemsPath := "/home/kibet/Pictures/FACEBOOK/PHONES/MARKETPLACE"

	fmt.Printf("Looking for items in %s\n", itemsPath)

	items, err := utils.GetItems(itemsPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d items to post\n", len(items))

	ctx := context.Background()
	rewriter, err := utils.NewGeminiDescriptionRewriter(ctx)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Printf("Title: %s\nPrice: %s\nCategory: %s\nCondition: %s\nDescription: %s\nTags: %v\nImages: %v\n\n",
			item.Title, item.Price, item.Category, item.Condition, item.Description, item.Tags, item.Images)

		description, err := rewriter.Rewrite(ctx, item.Description)
		if err != nil {
			fmt.Printf("Gemini rewrite failed for %s, using original description: %v\n", item.Title, err)
			description = item.Description
		}

		fmt.Println("============================================")
		fmt.Printf("Rewritten Description: %s\n\n", description)
		fmt.Println("============================================")

		item.Description = description

		utils.CreateListing(page, item)

		// Sleep for a random duration between 5 and 15 seconds to mimic human behavior
		time.Sleep(time.Duration(rand.IntN(10)+5) * time.Second)
	}
}
