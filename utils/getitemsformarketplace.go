package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MarketplaceItem struct {
	Images      []string
	Title       string
	Price       string
	Category    string
	Condition   string
	Description string
	Tags        []string
}

func GetItemsForMarketplace(path string) ([]MarketplaceItem, error) {
	var items []MarketplaceItem

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading root directory: %w", err)
	}

	for _, entry := range entries {
		subDir := filepath.Join(path, entry.Name())
		detailsFile := filepath.Join(subDir, "details.txt")

		// Read images from the subdirectory
		subEntries, err := os.ReadDir(subDir)
		if err != nil {
			continue // Handle error as needed
		}

		var imageFiles []string
		for _, subEntry := range subEntries {
			if !subEntry.IsDir() && filepath.Ext(subEntry.Name()) != ".txt" {
				filePath := filepath.Join(subDir, subEntry.Name())
				imageFiles = append(imageFiles, filePath)
			}
		}

		// Read details.txt for title, price, description
		file, err := os.Open(detailsFile)
		if err != nil {
			continue // Handle error as needed
		}
		defer file.Close()

		// Initialize variables to hold the extracted fields
		var title, price, category, condition, description, tagsString string

		// Create a new scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			switch {
			case strings.HasPrefix(line, "title:"):
				title = strings.ToUpper(strings.TrimSpace(line[len("title:"):]))
			case strings.HasPrefix(line, "price:"):
				price = strings.TrimSpace(line[len("price:"):])
			case strings.HasPrefix(line, "category"):
				category = strings.ToLower(strings.TrimSpace(line[len("category:"):]))
			case strings.HasPrefix(line, "condition"):
				condition = strings.ToLower(strings.TrimSpace(line[len("condition:"):]))
			case strings.HasPrefix(line, "description:"):
				description = strings.ToUpper(strings.TrimSpace(line[len("description:"):]))
			case strings.HasPrefix(line, "tags:"):
				tagsString = strings.TrimSpace(line[len("tags:"):])
			}
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		// Split the text by "..."
		parts := strings.Split(description, "...")

		// Trim spaces from each part
		for i := range parts {
			parts[i] = "✅ " + strings.TrimSpace(parts[i])
		}

		// Join the parts with "...\n"
		description = strings.Join(parts, "\n\n")

		tags := strings.Split(tagsString, ",")

		for i := len(tags) - 1; i >= 0; i-- {
			if tags[i] == "" {
				tags = tags[:i]
			} else {
				break
			}
		}

		// Create a PostContent instance and append to slice
		items = append(items, MarketplaceItem{
			Images:      imageFiles,
			Title:       title,
			Price:       price,
			Category:    category,
			Condition:   condition,
			Description: description,
			Tags:        tags,
		})
	}

	return items, nil
}
