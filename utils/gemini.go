package utils

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genai"
)

const (
	geminiDescriptionModel = "gemini-2.5-flash"
	emojiPhoneNumber       = "0️⃣7️⃣1️⃣8️⃣4️⃣4️⃣8️⃣4️⃣6️⃣1️⃣"
	geminiRewriteTimeout   = 30 * time.Second
)

type geminiDescriptionRewriter struct {
	client *genai.Client
}

func NewGeminiDescriptionRewriter(ctx context.Context) (*geminiDescriptionRewriter, error) {
	config, err := LoadConfig(".")
	if err != nil {
		return nil, fmt.Errorf("load .env config: %w", err)
	}

	if config.GEMINI_API_KEY == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is not set")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendGeminiAPI,
		APIKey:  config.GEMINI_API_KEY,
	})
	if err != nil {
		return nil, fmt.Errorf("create Gemini client: %w", err)
	}

	return &geminiDescriptionRewriter{client: client}, nil
}

func (r *geminiDescriptionRewriter) Rewrite(ctx context.Context, description string) (string, error) {
	description = strings.TrimSpace(description)
	if description == "" {
		return description, nil
	}

	prompt := fmt.Sprintf(`Rewrite this Facebook Marketplace product description to be concise, clear, natural, and persuasive.

Rules:
- Return only the rewritten description.
- Keep important product details.
- Do not invent details.
- Always make the description easy to scan using short lines and natural paragraph breaks instead of one wall of text.
- Use a few appropriate, persuasive emojis where they make the listing feel more engaging.
- Do not use Markdown formatting, asterisks, bullet symbols, or decorative separator symbols.
- Keep emojis natural and relevant; do not overuse them.
- Keep this phone number exactly as written: %s
- If the phone number is present, do not change its emoji format.
- If the phone number is missing, include it once at the end.

Description:
%s`, emojiPhoneNumber, description)

	config := &genai.GenerateContentConfig{
		ThinkingConfig: &genai.ThinkingConfig{
			ThinkingBudget: genai.Ptr[int32](0),
		},
	}

	ctx, cancel := context.WithTimeout(ctx, geminiRewriteTimeout)
	defer cancel()

	result, err := r.client.Models.GenerateContent(ctx, geminiDescriptionModel, genai.Text(prompt), config)
	if err != nil {
		return "", fmt.Errorf("rewrite description with Gemini: %w", err)
	}

	rewritten := strings.TrimSpace(result.Text())
	if rewritten == "" {
		return "", fmt.Errorf("Gemini returned an empty description")
	}

	if !strings.Contains(rewritten, emojiPhoneNumber) {
		rewritten = strings.TrimSpace(rewritten) + "\n\n" + emojiPhoneNumber
	}

	return rewritten, nil
}
