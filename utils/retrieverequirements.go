package utils

import (
	"bufio"
	"os"
	"strings"
)

func RetrieveRequirements(f string) (string, string, string, string, error) {
	file, err := os.Open(f)
	if err != nil {
		return "", "", "", "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var username, password, path, expiry string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Username:") {
			username = strings.TrimSpace(strings.TrimPrefix(line, "Username:"))
		} else if strings.HasPrefix(line, "Password:") {
			password = strings.TrimSpace(strings.TrimPrefix(line, "Password:"))
		} else if strings.HasPrefix(line, "Path:") {
			path = strings.TrimSpace(strings.TrimPrefix(line, "Path:"))
		} else if strings.HasPrefix(line, "Expiry:") {
			expiry = strings.TrimSpace(strings.TrimPrefix(line, "Expiry:"))
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", "", "", err
	}

	return username, password, path, expiry, nil
}
