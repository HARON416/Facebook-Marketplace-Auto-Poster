# 📁 Example Product Structure

This document shows you exactly how to organize your products for the MONICAH tool.

## 📂 Folder Structure

```
your-products/
├── laptop-dell-xps/
│   ├── dell-xps-1.jpg
│   ├── dell-xps-2.jpg
│   ├── dell-xps-3.jpg
│   └── details.txt
├── iphone-14-pro/
│   ├── iphone-front.jpg
│   ├── iphone-back.jpg
│   └── details.txt
└── gaming-chair/
    ├── chair-1.jpg
    ├── chair-2.jpg
    ├── chair-3.jpg
    └── details.txt
```

## 📝 Details.txt Examples

### Example 1: Electronics

```txt
title: Dell XPS 13 Laptop - Mint Condition
price: $899.99
category: electronics
condition: used - like new
description: This Dell XPS 13 laptop is in excellent condition... Barely used for 6 months... Comes with original charger and box... Perfect for work or school... No scratches or damage...
tags: laptop, dell, xps, computer, electronics, portable
```

### Example 2: Mobile Phones

```txt
title: iPhone 14 Pro - 256GB - Space Black
price: $999.00
category: mobile phones
condition: new
description: Brand new iPhone 14 Pro... Still sealed in original box... 256GB storage... Space Black color... Includes all original accessories... Apple warranty included...
tags: iphone, apple, smartphone, mobile, 5g, camera
```

### Example 3: Furniture

```txt
title: Ergonomic Gaming Chair - Black & Red
price: $299.99
category: furniture
condition: used - good
description: High-quality gaming chair... Excellent lumbar support... Adjustable armrests... Breathable mesh back... Perfect for long gaming sessions... Minor wear on armrests...
tags: chair, gaming, ergonomic, office, furniture, comfortable
```

## 🎯 Tips for Best Results

### Image Requirements

- **Format**: JPG, PNG, HEIC, HEIF
- **Size**: Facebook recommends high-quality images
- **Quantity**: 1-10 images per product
- **Content**: Clear, well-lit photos showing all angles

### Description Formatting

- Use `...` to separate sentences
- Each sentence will be formatted with ✅ checkmarks
- Keep descriptions honest and detailed
- Include key features and condition details

### Category Selection

Common categories include:

- `electronics`
- `mobile phones`
- `furniture`
- `clothing`
- `books`
- `sports`
- `automotive`
- `home & garden`

### Condition Options

- `new`
- `used - like new`
- `used - good`
- `used - fair`
- `used - poor`

### Price Formatting

- Use standard currency format: `$99.99`
- Include cents for professional appearance
- Be competitive with market prices

## 🚀 Getting Started

1. **Create your products folder**

   ```bash
   mkdir my-products
   cd my-products
   ```

2. **Create product subfolders**

   ```bash
   mkdir product-name
   cd product-name
   ```

3. **Add your images**

   - Copy all product images to the folder
   - Use descriptive filenames

4. **Create details.txt**

   - Use the format shown above
   - Be thorough with descriptions
   - Include relevant tags

5. **Test with MONICAH**
   ```bash
   make run
   ```

## ⚠️ Important Notes

- **File Names**: Use descriptive names for images
- **Details.txt**: Must be exactly named `details.txt`
- **Encoding**: Use UTF-8 encoding for special characters
- **Path**: Provide the full path to your products folder when prompted
- **Backup**: Keep backups of your product data

## 🔍 Troubleshooting

### Common Issues

1. **"Path does not exist"**

   - Double-check the folder path
   - Ensure the path is absolute or relative to the application

2. **"No items found"**

   - Verify each product folder has a `details.txt` file
   - Check that `details.txt` follows the exact format

3. **"Error reading file"**

   - Ensure `details.txt` is properly formatted
   - Check for special characters or encoding issues

4. **"Please upload at least one photo"**
   - Verify images are in supported formats
   - Check that image files are not corrupted
