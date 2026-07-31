package memory

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Habeebamoo/April/server/db"
)

func Extract(userMessage string, reply string) string {
	payload := map[string]any{
		"model":      os.Getenv("MODEL_NAME"),
		"max_tokens": 200,
		"messages": []map[string]any{
			{
				"role": "user",
				"content": `Extract the summary and important facts worth remembering from this conversation. 
					Be very concise. Only extract important details which will give you context to properly answer the next question.

					Habeeb said: "` + userMessage + `"
					April replied: "` + reply + `"`,
			},
		},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", os.Getenv("MODEL_BASE_URL")+"/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MODEL_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]any
	json.Unmarshal(respBody, &result)

	choices, ok := result["choices"].([]any)
	if !ok || len(choices) == 0 {
		return ""
	}

	message, _ := choices[0].(map[string]any)["message"].(map[string]any)
	content, _ := message["content"].(string)
	return content
}

func Save(userMessage string, reply string) {
	// Extract meaningful facts first
	extracted := Extract(userMessage, reply)

	// If nothing worth saving, skip
	if extracted == "" || extracted == "NOTHING" {
		log.Println("Memory: nothing worth saving")
		return
	}

	log.Println("Memory saving:", extracted)

	memory := Memory{Content: extracted}
	result := db.DB.Create(&memory)
	if result.Error != nil {
		log.Println("Failed to save memory:", result.Error)
	}
}