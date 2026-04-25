package utils

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
)

type Item struct {
	Images      []string
	Title       string
	Price       string
	Category    string
	Condition   string
	Description string
	Tags        []string
}

func GetItems(path string) ([]Item, error) {
	var items []Item

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading root directory: %w", err)
	}

	for _, entry := range entries {
		subDir := filepath.Join(path, entry.Name())
		detailsFile := filepath.Join(subDir, "description.txt")

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

		rand.Shuffle(len(imageFiles), func(i, j int) {
			imageFiles[i], imageFiles[j] = imageFiles[j], imageFiles[i]
		})
		if len(imageFiles) > 6 {
			imageFiles = imageFiles[:6]
		}

		// Read description.txt for title, price, description
		file, err := os.Open(detailsFile)
		if err != nil {
			continue // Handle error as needed
		}
		defer file.Close()

		// Initialize variables to hold the extracted fields
		var title, price, category, condition, description, tagsString string
		var descriptionLines []string
		inDescription := false

		// Create a new scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			if inDescription {
				switch {
				case strings.HasPrefix(line, "tags:"):
					tagsString = strings.TrimSpace(line[len("tags:"):])
					inDescription = false
					continue
				case strings.HasPrefix(line, "title:"),
					strings.HasPrefix(line, "price:"),
					strings.HasPrefix(line, "category:"),
					strings.HasPrefix(line, "condition:"):
					inDescription = false
				default:
					descriptionLines = append(descriptionLines, line)
					continue
				}
			}

			switch {
			case strings.HasPrefix(line, "title:"):
				title = strings.ToUpper(strings.TrimSpace(line[len("title:"):]))
			case strings.HasPrefix(line, "price:"):
				price = strings.TrimSpace(line[len("price:"):])
			case strings.HasPrefix(line, "category:"):
				category = strings.ToLower(strings.TrimSpace(line[len("category:"):]))
			case strings.HasPrefix(line, "condition:"):
				condition = strings.ToLower(strings.TrimSpace(line[len("condition:"):]))
			case strings.HasPrefix(line, "description:"):
				descriptionLines = []string{strings.TrimSpace(line[len("description:"):])}
				inDescription = true
			case strings.HasPrefix(line, "tags:"):
				tagsString = strings.TrimSpace(line[len("tags:"):])
			}
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		description = strings.Join(descriptionLines, "\n")

		tags := strings.Split(tagsString, ",")

		for i := len(tags) - 1; i >= 0; i-- {
			if tags[i] == "" {
				tags = tags[:i]
			} else {
				break
			}
		}

		// Create a PostContent instance and append to slice
		items = append(items, Item{
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
