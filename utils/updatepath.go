package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UpdatePath(f string) (string, error) {
	// Create a reader for standard input
	reader := bufio.NewReader(os.Stdin)

	// Ask for the path
	fmt.Print("Enter correct path to folder containing your items: ")
	path, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	path = path[:len(path)-1] // Remove newline character

	// Read the file
	file, err := os.Open(f)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Check if the line starts with "Path:"
		if strings.HasPrefix(line, "Path:") {
			// Update the path value
			newPath := fmt.Sprintf("Path: %s", path)
			line = newPath
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Write updated lines back to the file
	file, err = os.Create(f)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return "", err
		}
	}
	writer.Flush()

	return path, nil
}
