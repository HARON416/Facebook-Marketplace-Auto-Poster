# MONICAH — Facebook Marketplace Listing Tool

MONICAH posts Facebook Marketplace items from folders on your computer. Add each product's photos and listing details to its own folder, run the tool, and MONICAH opens Chrome, completes the listing form, uploads the images, and publishes it.

Built for general Facebook Marketplace **items for sale** from **local folders**—including furniture, electronics, phones, and other one-off listings.

> **Looking for something else? Check out my other tools:**
>
> - 🚗 **Auto dealerships and sales teams** — Tired of manually re-listing vehicle inventory? I built a Chrome extension that posts vehicles directly from your dealership website in one click—one unit or your whole lot. Trusted by Steve Marshall Group, Mainland Motors, Applewood, Tricity Mitsubishi, and more. [Message me on WhatsApp](https://wa.me/254718448461) or [send me an email](mailto:haronkibetrutoh@gmail.com) for a quick demo.
> - 📋 **[KundiPost](https://chromewebstore.google.com/detail/kundipost-%E2%80%93-facebook-grou/lphjckpophfkedacijahclaoenhikcgo)** — Save and reuse content, publish to Facebook Groups and Marketplace with paced queues, rewrite content with Gemini, and clean up old Marketplace listings—all from a Chrome side panel.

## What MONICAH Does

- Creates multiple item-for-sale listings from local product folders.
- Fills in the title, price, category, condition, description, and product tags.
- Randomly selects and uploads up to six images per item.
- Optionally rewrites descriptions with Gemini while preserving the configured WhatsApp number.
- Selects suggested audience groups before publishing.
- Keeps your Facebook login in a dedicated local Chrome profile for future runs.

## Requirements

- A Facebook account with Marketplace access.
- Google Chrome.
- Go 1.24 or later.
- A Gemini API key if you want automatic description rewriting.

## Quick Start

1. Install the dependencies:

```bash
make tidy
```

2. In `main.go`, set `itemsPath` to the folder containing your products:

```go
itemsPath := "/path/to/your/marketplace/items"
```

3. Optional: create a `.env` file to enable Gemini rewriting:

```env
GEMINI_API_KEY=your_gemini_api_key
```

4. Run MONICAH:

```bash
make run
```

Chrome will open. Log in to Facebook when prompted, and MONICAH will begin creating listings from your folders. If Gemini is not configured, press Enter when prompted to continue with the original descriptions.

## Prepare Your Listings

Create one folder per item. Each folder needs a `description.txt` file and at least one image:

```text
MARKETPLACE/
├── samsung-note-20-ultra/
│   ├── front.jpg
│   ├── back.jpg
│   └── description.txt
└── office-chair/
    ├── chair.jpg
    └── description.txt
```

Format each `description.txt` like this:

```text
description: Clean Samsung Galaxy Note 20 Ultra in excellent condition.
Includes the original charger. Delivery available within Nairobi.
title: Samsung Galaxy Note 20 Ultra
price: 45000
category: Mobile phones
condition: Used - like new
tags: samsung, smartphone, note 20 ultra
```

The description may span multiple lines. MONICAH treats everything after `description:` as part of the description until it reaches `title:`, `price:`, `category:`, `condition:`, or `tags:`.

## Before You Run It

- MONICAH supports standard Marketplace **item-for-sale** listings—not vehicles, property, rentals, or inventory pulled from websites.
- Facebook changes its interface regularly, so selectors may occasionally need updating.
- Category and condition values must match options available on Facebook Marketplace.
- MONICAH publishes listings automatically. Review your source files before running it and use it responsibly.
- Your Facebook session is stored locally in the project's `browser_profile` directory. MONICAH does not ask for or store your Facebook password directly.

## Need a Custom Version?

I build custom Marketplace automation for vehicle dealerships, real estate, online inventory, and other specialized workflows.

[Request a demo on WhatsApp](https://wa.me/254718448461) or email [haronkibetrutoh@gmail.com](mailto:haronkibetrutoh@gmail.com).

## License

Proprietary software. All rights reserved.
