package utils

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Login(username, password, url string) (*rod.Browser, *rod.Page) {
	dir := "~/.config/google-chrome"

	u := launcher.New().UserDataDir(dir).Leakless(true).NoSandbox(true).Headless(false).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect().NoDefaultDevice()

	page := browser.MustPage(url).MustWaitLoad().MustWindowMaximize()

	if page.MustHas(`button[name="login"]`) {
		page.MustElement(`input[name="email"]`).MustInput(username)
		page.MustElement(`input[name="pass"]`).MustInput(password)
		page.MustElement(`button[type="submit"]`).MustClick()
		time.Sleep(1 * time.Minute)
	}

	return browser, page
}
