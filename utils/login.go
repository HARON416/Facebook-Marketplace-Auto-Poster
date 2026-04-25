package utils

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Login() (*rod.Browser, *rod.Page) {

	url := "https://web.facebook.com/"

	u := launcher.NewUserMode().
		Leakless(true).
		NoSandbox(true).
		Headless(false).
		Devtools(false).
		UserDataDir("~/.config/google-chrome").
		Set("disable-notifications").
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect().NoDefaultDevice()

	page := browser.MustPage(url).MustWindowMaximize().MustWaitLoad().MustWaitDOMStable()

	fmt.Println("⏳ Initializing Facebook login")

	if !page.MustHas(`div[aria-label="Your profile"]`) {
		fmt.Println("⏳ Login to Facebook in the browser window opened by this tool. You have 3 minutes to complete the login.")

		for i := 180; i > 0; i-- {
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
