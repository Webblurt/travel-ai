package ai

import (
	"context"
	"fmt"
)

// Stub for future Claude integration
type AIClient struct{}

func NewAIClient() *AIClient {
	return &AIClient{}
}

// GenerateItinerary simulates LLM output.
func (c *AIClient) GenerateItinerary(ctx context.Context, destination string, days int) (map[string]interface{}, error) {
	fmt.Println("[AI] Simulating itinerary generation for:", destination)

	// Stubbed example output
	return map[string]interface{}{
		"days": []map[string]interface{}{
			{
				"day": 1,
				"plan": []string{
					fmt.Sprintf("Arrival in %s, city walk, dinner at local restaurant", destination),
				},
			},
			{
				"day": 2,
				"plan": []string{
					fmt.Sprintf("Morning museum visit in %s", destination),
					"Afternoon coffee and art market",
				},
			},
		},
	}, nil
}
