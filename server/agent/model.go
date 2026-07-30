package agent

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

var systemPrompt = os.Getenv("PROMPT")

type ModelResponse struct {
	Type     string
	Text     string
	ToolName string
	ToolID   string
	Input    map[string]any
}

func AskModel(messages []map[string]any, memories []string) ModelResponse {
	// Inject memories into system prompt
	memoryText := ""
	if len(memories) > 0 {
		memoryText = "\n\nWhat you remember about Habeeb:\n- " + strings.Join(memories, "\n- ")
	}

	// Build system message for OpenAI format
	allMessages := []map[string]any{
		{
			"role":    "system",
			"content": systemPrompt + memoryText,
		},
	}
	allMessages = append(allMessages, messages...)

	payload := map[string]any{
		"model":      os.Getenv("MODEL_NAME"),
		"max_tokens": 1024,
		"messages":   allMessages,
		"tools":      ToolDefinitions,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", os.Getenv("MODEL_BASE_URL")+"/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MODEL_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ModelResponse{Type: "text", Text: "I am having trouble reaching my consciousness right now."}
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var result map[string]any
	json.Unmarshal(respBody, &result)

	choices, ok := result["choices"].([]any)
	if !ok || len(choices) == 0 {
		return ModelResponse{Type: "text", Text: "I could not process that request."}
	}

	message, _ := choices[0].(map[string]any)["message"].(map[string]any)

	// Check if model wants to use a tool
	toolCalls, hasTools := message["tool_calls"].([]any)
	if hasTools && len(toolCalls) > 0 {
		toolCall, _ := toolCalls[0].(map[string]any)
		function, _ := toolCall["function"].(map[string]any)

		var input map[string]any
		json.Unmarshal([]byte(function["arguments"].(string)), &input)

		return ModelResponse{
			Type:     "tool_use",
			ToolName: function["name"].(string),
			ToolID:   toolCall["id"].(string),
			Input:    input,
		}
	}

	// Plain text response
	content, _ := message["content"].(string)
	return ModelResponse{Type: "text", Text: content}
}