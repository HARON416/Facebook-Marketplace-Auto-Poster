# 🚀 MONICAH - Facebook Marketplace Auto Poster 

A powerful Go-based automation tool that streamlines the process of posting items to Facebook Marketplace. Built with modern web automation using Rod browser library and featuring a beautiful CLI interface.

## ✨ Features

- **🔐 Secure Authentication**: Secure Facebook login with credential management
- **📦 Batch Processing**: Automatically process multiple items from organized folders
- **🖼️ Image Upload**: Seamless image upload support for product listings
- **📝 Smart Content Parsing**: Automatic parsing of product details from structured files
- **⏱️ Intelligent Timing**: Random delays between posts to avoid rate limiting
- **🎯 Group Targeting**: Automatic selection of relevant Facebook groups
- **🔄 Trial Management**: Built-in trial period management system
- **📊 Progress Tracking**: Real-time logging and progress monitoring

## 🏗️ Architecture

### Core Components

```
MONICAH/
├── main.go                 # Application entry point
├── makefile               # Build and run commands
├── go.mod                 # Go module dependencies
├── utils/                 # Core functionality modules
│   ├── login.go          # Facebook authentication
│   ├── getrequirements.go # User input collection
│   ├── saverequirements.go # Credential persistence
│   ├── retrieverequirements.go # Credential retrieval
│   ├── getitemsformarketplace.go # Product data parsing
│   ├── posttomarketplace.go # Marketplace posting automation
│   └── updatepath.go     # Path management
└── .gitignore            # Git ignore rules
```

## 🚀 Quick Start

### Prerequisites

- Go 1.23.4 or higher
- Google Chrome browser
- Facebook account

### Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd MONICAH
   ```

2. **Install dependencies**

   ```bash
   make tidy
   ```

3. **Run the application**
   ```bash
   make run
   ```

## 📁 Project Structure Setup

### Product Organization

Organize your products in the following structure:

```
products/
├── product1/
│   ├── image1.jpg
│   ├── image2.jpg
│   └── details.txt
├── product2/
│   ├── image1.jpg
│   └── details.txt
└── product3/
    ├── image1.jpg
    ├── image2.jpg
    └── details.txt
```

### Details.txt Format

Each product folder should contain a `details.txt` file with the following format:

```txt
title: Your Product Title
price: $99.99
category: electronics
condition: new
description: This is a detailed description of your product... You can use multiple sentences... Each sentence will be formatted with checkmarks...
tags: electronics, gadget, tech, wireless
```

## 🔧 Configuration

### First Run Setup

On first run, the application will prompt for:

1. **Facebook Username/Email**
2. **Facebook Password** (hidden input)
3. **Path to products folder**

Credentials are securely stored in `requirements.txt` for future use.

### Trial Period

The application includes a 72-hour trial period. After expiration, contact support for upgrade options.

## 🛠️ Usage

### Running the Application

```bash
make run
```

### Available Commands

- `make run` - Run the application
- `make tidy` - Clean and update Go dependencies

## 🔒 Security Features

- **Credential Encryption**: Passwords are stored securely
- **Session Management**: Automatic browser session handling
- **Rate Limiting**: Intelligent delays between posts
- **Error Handling**: Comprehensive error management

## 📊 How It Works

### 1. Authentication Flow

```go
// Secure Facebook login with Rod browser automation
browser, page = u.Login(username, password, url)
```

### 2. Product Processing

```go
// Parse product details from organized folders
items, err = u.GetItemsForMarketplace(path)
```

### 3. Marketplace Automation

```go
// Automated posting with intelligent timing
err = u.PostItemsToMarketplace(browser, page, items)
```

### 4. Smart Features

- **Random Shuffling**: Products are randomly shuffled before posting
- **Intelligent Delays**: 30-60 second random delays between posts
- **Group Selection**: Automatic selection of relevant Facebook groups
- **Progress Monitoring**: Real-time logging of all operations

## 🎯 Key Features Explained

### Product Parsing

- Automatically reads product images from folders
- Parses structured `details.txt` files
- Formats descriptions with checkmarks for better presentation
- Handles multiple images per product

### Marketplace Integration

- Navigates to Facebook Marketplace create item page
- Uploads multiple images per product
- Fills all required fields (title, price, category, condition, description, tags)
- Handles category and condition selection from dropdowns
- Manages product tags with Enter key simulation

### Error Handling

- Detects rate limiting and posting limits
- Handles network errors gracefully
- Validates file paths and product data
- Provides clear error messages and recovery options

## 🔧 Technical Details

### Dependencies

- **github.com/go-rod/rod**: Modern browser automation
- **github.com/charmbracelet/log**: Beautiful CLI logging
- **golang.org/x/term**: Secure password input

### Browser Automation

Uses Rod browser library for:

- Chrome browser automation
- Headless/headed mode support
- Element interaction and form filling
- File upload handling
- Navigation and page waiting

### Performance Features

- **Concurrent Processing**: Efficient handling of multiple products
- **Memory Management**: Proper browser cleanup and resource management
- **Timeout Handling**: Intelligent retry mechanisms
- **Progress Tracking**: Real-time operation logging

## 🚨 Important Notes

### Geographic Limitations

- **Marketplace Layout Variations**: Facebook Marketplace layout and functionality vary significantly by country
- **Current Compatibility**: This tool is optimized for Kenya and similar markets
- **US/Other Countries**: May not work properly in the United States or other countries due to different marketplace layouts
- **Customization Available**: Contact the developer for country-specific customizations and adaptations

### Rate Limiting

- Facebook has posting limits to prevent spam
- The tool includes intelligent delays to avoid detection
- Monitor for "Limit reached" messages

### Browser Requirements

- Requires Google Chrome browser
- Uses Chrome user data directory for session persistence
- Supports both headless and headed modes

### File Structure Requirements

- Products must be organized in folders
- Each folder must contain a `details.txt` file
- Images should be in common formats (jpg, png, etc.)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is proprietary software. All rights reserved.

## 🆘 Support

For support, licensing inquiries, and country-specific customizations, please contact:

- **Email**: haronkibetrutoh@gmail.com
- **WhatsApp**: +254718448461

We provide customizations for different countries where Facebook Marketplace layouts may vary.

---

**⚠️ Disclaimer**: This tool is for educational and legitimate business use only. Users are responsible for complying with Facebook's Terms of Service and applicable laws.
