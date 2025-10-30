package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/sashabaranov/go-openai"
)

func AskOpenAI[T any](c *Client, promptType string, req interface{}) (*T, error) {
	ctx := context.Background()

	var promptConfig struct {
		SystemRole string
		UserPrompt string
	}
	found := false
	for _, p := range c.cfg.Prompts.TravelAI {
		if p.PromptType == promptType {
			promptConfig.SystemRole = p.SystemRole
			promptConfig.UserPrompt = p.UserPrompt
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("prompt type %q not found", promptType)
	}

	tmpl, err := template.New("promptType").Parse(promptConfig.UserPrompt)
	if err != nil {
		return nil, fmt.Errorf("template parse error: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, req); err != nil {
		return nil, fmt.Errorf("template exec error: %w", err)
	}
	prompt := buf.String()

	resp, err := c.openai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: promptConfig.SystemRole},
			{Role: "user", Content: prompt},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("openai error: %w", err)
	}

	content := resp.Choices[0].Message.Content
	var data T
	if err := json.Unmarshal([]byte(content), &data); err != nil {
		var s string
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", content)), &s); err == nil {
			return nil, fmt.Errorf("non-JSON response: %s", s)
		}
		return nil, fmt.Errorf("failed to unmarshal response into type %T: %w", data, err)
	}

	return &data, nil
}

func (c *Client) AskOpenAI(promptType string, req interface{}) (interface{}, error) {
	return AskOpenAI[interface{}](c, promptType, req)
}
