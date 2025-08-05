package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func GetRequirements() (string, string, string, error) {
	// Create a reader for standard input
	reader := bufio.NewReader(os.Stdin)

	// Ask for the username
	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", "", err
	}
	username = strings.TrimSpace(username) // Remove newline character

	// Ask for the password
	fmt.Print("Enter password: ")
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", "", "", nil
	}
	fmt.Println() // Add a newline after password input
	password := string(passwordBytes)

	// Ask for the username
	fmt.Print("Enter the path to folder containing your products: ")
	path, err := reader.ReadString('\n')
	if err != nil {
		return "", "", "", err
	}

	path = strings.TrimSpace(path)

	return username, password, path, nil
}
