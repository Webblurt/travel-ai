package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	models "travel-ai/internal/models"

	"github.com/sashabaranov/go-openai"
)

func (c *Client) GenerateItinerary(req models.OpenAIReq) (models.OpenAIResp, error) {
	ctx := context.Background()
	tmpl, err := template.New("travelPrompt").Parse(c.cfg.Prompts.TravelAI.UserPrompt)
	if err != nil {
		return models.OpenAIResp{}, fmt.Errorf("template parse error: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, req); err != nil {
		return models.OpenAIResp{}, fmt.Errorf("template exec error: %w", err)
	}
	prompt := buf.String()

	resp, err := c.openai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: c.cfg.Prompts.TravelAI.SystemRole},
			{Role: "user", Content: prompt},
		},
	})
	if err != nil {
		return models.OpenAIResp{}, fmt.Errorf("openai error: %w", err)
	}

	var itinerary models.OpenAIResp
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &itinerary); err != nil {
		return models.OpenAIResp{}, fmt.Errorf("invalid JSON: %w", err)
	}

	return itinerary, nil
}
