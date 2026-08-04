package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type TavilyResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

type TavilyResponse struct {
	Results []TavilyResult `json:"results"`
	Answer  string         `json:"answer"`
}

func WebSearch(input map[string]any) string {
	query, ok := input["query"].(string)
	if !ok || query == "" {
		return "Please provide a search query"
	}

	payload := map[string]any{
		"api_key":        os.Getenv("TAVILY_API_KEY"),
		"query":          query,
		"search_depth":   "basic",
		"include_answer": true,
		"max_results":    5,
	}

	body, _ := json.Marshal(payload)
	resp, err := http.Post("https://api.tavily.com/search", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "Failed to search the web"
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result TavilyResponse
	json.Unmarshal(respBody, &result)

	// If Tavily gives a direct answer, use it
	if result.Answer != "" {
		return fmt.Sprintf("Answer: %s", result.Answer)
	}

	// Otherwise return top results
	if len(result.Results) == 0 {
		return "No results found"
	}

	var lines []string
	for _, r := range result.Results {
		lines = append(lines, fmt.Sprintf("- %s\n  %s\n  %s", r.Title, r.URL, r.Content))
	}
	return strings.Join(lines, "\n\n")
}