package utils

import (
	"fmt"
	"os"
	"time"
)

func SaveRequirements(username, password, path, f string) (string, error) {
	expiry := time.Now().Local().String()

	file, err := os.Create(f)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Username: %s\nPassword: %s\nPath: %s\nExpiry: %s\n", username, password, path, expiry))
	if err != nil {
		return "", err
	}

	return expiry, nil
}
