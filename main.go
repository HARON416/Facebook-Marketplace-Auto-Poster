package main

import (
	"errors"
	"math"
	"math/rand"
	"os"
	"time"

	l "github.com/charmbracelet/log"
	"github.com/go-rod/rod"
	u "github.com/kwandapchumba/thelazyfbposter/utils"
)

func main() {
	url := "https://web.facebook.com/"
	file := "requirements.txt"
	var username, password, path, expiry string
	var browser *rod.Browser
	var page *rod.Page
	var items []u.MarketplaceItem

	username, password, path, expiry, err := u.RetrieveRequirements(file)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			username, password, path, err = u.GetRequirements()
			if err != nil {
				l.Error(err)
				os.Exit(1)
			}

			expiry, err = u.SaveRequirements(username, password, path, file)
			if err != nil {
				l.Error(err)
				os.Exit(1)
			}
		default:
			l.Error(err)
			os.Exit(1)
		}
	}

	if username == "" || password == "" || path == "" || expiry == "" {
		username, password, path, err = u.GetRequirements()
		if err != nil {
			l.Error(err)
			os.Exit(1)
		}

		expiry, err = u.SaveRequirements(username, password, path, file)
		if err != nil {
			l.Error(err)
			os.Exit(1)
		}
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", expiry)
	if err != nil {
		l.Error("Error parsing time", "err", err)
		return
	}

	elapsedTime := time.Since(parsedTime)

	if elapsedTime > 72*time.Hour {
		l.Warn("Your free trial has ended. Please contact support to upgrade")
		os.Exit(1)
	}

	browser, page = u.Login(username, password, url)

	items, err = u.GetItemsForMarketplace(path)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			l.Warn("Path does not exist")
			_, err = u.UpdatePath(file)
			if err != nil {
				l.Error(err)
				os.Exit(1)
			}

			l.Warn("Rerun the program")
			os.Exit(1)
		default:
			l.Error(err)
			os.Exit(1)
		}
	}

	totalItems := len(items)

	if totalItems < 1 {
		l.Info("No items found in the parent directory")
		os.Exit(1)
	}

	l.Info("Do not close this window")

	start := time.Now()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	l.Info("Items shuffled")

	err = u.PostItemsToMarketplace(browser, page, items)
	if err != nil {
		l.Error(err)
		os.Exit(1)
	}

	duration := math.Round(time.Since(start).Minutes())

	l.Infof("ALL ITEMS LISTED IN %f MINUTES", duration)
}
