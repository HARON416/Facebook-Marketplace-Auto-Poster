package utils

import (
	"fmt"
	"hash/fnv"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Login() (*rod.Browser, *rod.Page) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	accountName := filepath.Base(cwd)
	profileDir := filepath.Join(cwd, "browser_profile")
	if err := os.MkdirAll(profileDir, 0o755); err != nil {
		panic(err)
	}

	port, err := accountDebugPort(cwd)
	if err != nil {
		panic(err)
	}

	url := "https://web.facebook.com/"

	fmt.Printf("Starting browser - Account: %s, Profile: %s, Port: %d\n", accountName, profileDir, port)

	u := launcher.NewUserMode().
		Leakless(true).
		NoSandbox(true).
		Headless(false).
		Devtools(false).
		UserDataDir(profileDir).
		Set("remote-debugging-port", fmt.Sprintf("%d", port)).
		Set("disable-notifications").
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect().NoDefaultDevice()

	page := browser.MustPage(url).MustWindowMaximize().MustWaitLoad().MustWaitDOMStable()

	fmt.Println("⏳ Initializing Facebook login")

	if !page.MustHas(`div[aria-label="Your profile"]`) {
		fmt.Println("⏳ Login to Facebook in the browser window opened by this tool. You have 3 minutes to complete the login.")

		for remaining := range 180 {
			i := 180 - remaining
			mins := i / 60
			secs := i % 60
			fmt.Printf("\rTime remaining: %02d:%02d", mins, secs)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("\rTime remaining: 00:00")
	}

	fmt.Println("✅ Login successful")

	return browser, page
}

func accountDebugPort(accountPath string) (int, error) {
	const (
		minPort = 9222
		maxPort = 65535
	)

	portRange := maxPort - minPort + 1
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(accountPath))
	start := minPort + int(hash.Sum32()%uint32(portRange))

	for i := range portRange {
		port := minPort + ((start - minPort + i) % portRange)
		if portAvailable(port) {
			return port, nil
		}
	}

	return 0, fmt.Errorf("no available browser debugging ports in range %d-%d", minPort, maxPort)
}

func portAvailable(port int) bool {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return false
	}
	return listener.Close() == nil
}
