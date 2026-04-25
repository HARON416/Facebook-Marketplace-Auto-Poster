# MONICAH - Facebook Marketplace Auto Poster / Facebook Marketplace Listing Tool / Facebook Marketplace Listing Software

MONICAH helps you post item-for-sale listings to Facebook Marketplace from product folders on your computer. Each folder represents one item, with photos and a simple `description.txt` file.

This tool is built for local-folder posting. It does not pull listings from websites by default, and it is not for rentals, vehicles, or real estate unless customized.

## What It Does

- Opens Chrome and lets you log in to Facebook.
- Reads items from folders on your computer.
- Uploads up to 6 random photos per item. You can change the number of selected photos in `utils/get_items.go
- Reads title, price, category, condition, description, and tags.
- Rewrites descriptions with Gemini so they are concise and clear.
- Keeps your WhatsApp phone number in the listing description.
- Posts each item to Facebook Marketplace with delays between posts.

## What You Need

- A Facebook account with Marketplace access.
- Google Chrome installed.
- Go installed on your computer.
- A Gemini API key saved in a `.env` file.
- Your product photos and descriptions arranged in folders.

## Setup

Create a `.env` file in the project folder:

```env
GEMINI_API_KEY=your_gemini_api_key_here
```

Install the project dependencies:

```bash
make tidy
```

## Prepare Your Items

Create one folder for each item you want to post.

Example:

```txt
MARKETPLACE/
├── iphone-13/
│   ├── photo1.jpg
│   ├── photo2.jpg
│   └── description.txt
├── samsung-a54/
│   ├── front.jpg
│   ├── back.jpg
│   └── description.txt
└── laptop/
    ├── image1.jpg
    ├── image2.jpg
    └── description.txt
```

Each item folder must contain a `description.txt` file.

Example `description.txt`:

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

Multiline descriptions are supported. The description starts at `description:` and continues until the next field such as `title:`, `price:`, `category:`, `condition:`, or `tags:`.

## Choose Your Product Folder

Open `main.go` and update this line to point to your own folder:

```go
itemsPath := "/home/kibet/Pictures/FACEBOOK/PHONES/MARKETPLACE"
```

Use the folder that contains all your item folders.

## Run The Poster

Start the app:

```bash
make run
```

Chrome will open. If you are not already logged in, log in to Facebook in the browser window. After that, the app will read your folders and begin posting your items.

## Important Notes

- This currently works only for normal Facebook Marketplace items for sale.
- Listings must come from local folders on your computer.
- Facebook Marketplace layouts can change, so the script may need adjustment over time.
- The app uses your existing Chrome session and does not save your Facebook password.
- If Gemini fails to rewrite a description, the original description is used.
- Only up to 6 photos are selected per item, randomly shuffled from the item folder.
- You can change the number of selected photos in `utils/get_items.go`.
- You can change the random sleep time between posts in `main.go`.
- Suggested groups are auto-selected by default. You can disable this in `utils/post_to_marketplace.go`.

## Customization

Need a version made for your business or listing style? Custom scripts are available for:

- Vehicle listings.
- Item listings.
- Real estate listings.
- Listings from local folders.
- Listings from online sources, for example a vehicle dealership website to Facebook Marketplace.

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
