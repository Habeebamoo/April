package agent

import (
	"encoding/json"
	"log"
)

func Run(userMessage string, memories []string) string {

	messages := []map[string]any{
		{"role": "user", "content": userMessage},
	}

	response := AskModel(messages, memories)

	if response.Type == "text" {
		return response.Text
	}

	// Execute the tool
	toolResult := ExecuteTool(response.ToolName, response.Input)

	// Marshal input back to JSON string correctly
	argumentsBytes, err := json.Marshal(response.Input)
	if err != nil {
		log.Println("Failed to marshal arguments:", err)
		return "Something went wrong"
	}

	// Add model tool call to messages
	messages = append(messages, map[string]any{
		"role":    "assistant",
		"content": nil,
		"tool_calls": []map[string]any{
			{
				"id":   response.ToolID,
				"type": "function",
				"function": map[string]any{
					"name":      response.ToolName,
					"arguments": string(argumentsBytes),
				},
			},
		},
	})

	// Add tool result to messages
	messages = append(messages, map[string]any{
		"role":         "tool",
		"tool_call_id": response.ToolID,
		"content":      toolResult,
	})

	// Final call to model
	finalResponse := AskModel(messages, memories)
	return finalResponse.Text
}