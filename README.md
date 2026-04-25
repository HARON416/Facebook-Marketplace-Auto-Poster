# MONICAH

**Facebook Marketplace Auto Poster for item-for-sale listings.**

MONICAH helps you post Facebook Marketplace items from folders on your computer. Put each product in its own folder, add photos plus a `description.txt` file, and the tool opens Chrome, fills the Marketplace form, uploads images, and posts the item.

> Built for normal Facebook Marketplace **items for sale** from **local folders**. Vehicles, real estate, rentals, online inventory, and dealership websites need a custom version.

## At A Glance

| Feature | What it means |
| --- | --- |
| Local folders | Each product lives in its own folder on your computer. |
| Photo upload | Up to 6 random photos are selected per item by default. |
| Description cleanup | Gemini can rewrite descriptions to be concise and clear. |
| Phone retention | The WhatsApp number is kept in emoji format. |
| Chrome session | Uses your existing Chrome login session. |
| Group posting | Suggested groups are auto-selected by default. |

## What You Need

- A Facebook account with Marketplace access.
- Google Chrome installed.
- Go installed.
- A Gemini API key.
- Product folders containing photos and `description.txt`.

## Quick Start

1. Create a `.env` file in this project folder:

```env
GEMINI_API_KEY=your_gemini_api_key_here
```

2. Install dependencies:

```bash
make tidy
```

3. Open `main.go` and set your products folder:

```go
itemsPath := "/home/kibet/Pictures/FACEBOOK/PHONES/MARKETPLACE"
```

4. Run the poster:

```bash
make run
```

Chrome will open. Log in to Facebook if needed, then the tool will read your folders and start posting.

## Folder Format

Create one folder per item:

```txt
MARKETPLACE/
├── samsung-note-20-ultra/
│   ├── front.jpg
│   ├── back.jpg
│   ├── side.jpg
│   └── description.txt
├── iphone-13/
│   ├── photo1.jpg
│   ├── photo2.jpg
│   └── description.txt
└── laptop/
    ├── image1.jpg
    ├── image2.jpg
    └── description.txt
```

Images can use normal formats such as `.jpg`, `.jpeg`, `.png`, or `.webp`. Text files are ignored as images.

## Description File

Each item folder must contain `description.txt`.

Use this format:

```txt
description: ✅ LIPA MDOGO MDOGO
✅ SAMSUNG GALAXY NOTE 20 ULTRA
✅ DEPOSIT: KSH. 11,199
✅ WEEKLY: KSH. 1,370
✅ FIRST WEEK FREE
✅ LOCATION: PIONEER BUILDING, KIMATHI STREET, 5TH FLOOR, SHOP #2
✅ NAIROBI AND ITS ENVIRONS ONLY. FREE DELIVERY. PAY AFTER DELIVERY
✅ AGENT CODE: 514 SHULEM PHONES AND ACCESSORIES
✅ CALL/WHATSAPP 0️⃣7️⃣1️⃣8️⃣4️⃣4️⃣8️⃣4️⃣6️⃣1️⃣
title: SAMSUNG GALAXY NOTE 20 ULTRA LIPA MDOGO MDOGO
price: 11199
category: Mobile phones
condition: New
tags: lipa mdogo mdogo smartphones, samsung lipa mdogo mdogo, iphone lipa mdogo mdogo, lipa pole pole smartphones, smartphones nairobi
```

The description can be multiple lines. It starts at `description:` and continues until the next field: `title:`, `price:`, `category:`, `condition:`, or `tags:`.

## What Can Be Changed

| Setting | Where to change it |
| --- | --- |
| Product folder path | `main.go` |
| Random sleep time between posts | `main.go` |
| Number of selected photos | `utils/get_items.go` |
| Suggested group auto-selection | `utils/post_to_marketplace.go` |
| Gemini rewrite prompt and timeout | `utils/gemini.go` |

## Important Notes

- This tool currently works only for normal Marketplace items for sale.
- Listings must come from local folders on your computer.
- Facebook Marketplace changes often, so selectors may need updates over time.
- The app uses your existing Chrome session and does not save your Facebook password.
- If Gemini fails to rewrite a description, the original description is used.
- Up to 6 photos are selected per item by default, randomly shuffled from the folder.

## Customization

Custom versions are available for:

- Vehicle listings.
- Item listings.
- Real estate listings.
- Local folder inventory.
- Online inventory, for example a vehicle dealership website to Facebook Marketplace.

For customization, contact:

- Email: haronkibetrutoh@gmail.com
- WhatsApp: +254718448461

## Commands

```bash
make run
make tidy
```

## License

This project is proprietary software. All rights reserved.
