package agent

var ToolDefinitions = []map[string]any{
	{
		"type": "function",
		"function": map[string]any{
			"name":        "send_email",
			"description": "Send a beautifully styled HTML email on behalf of Habeeb. Always write a complete, well designed HTML+CSS email as the body.",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"to": map[string]any{
						"type":        "string",
						"description": "Recipient email address",
					},
					"subject": map[string]any{
						"type":        "string",
						"description": "Email subject line",
					},
					"body": map[string]any{
						"type":        "string",
						"description": "Complete HTML+CSS email body",
					},
				},
				"required": []string{"to", "subject", "body"},
			},
		},
	},
	{
		"type": "function",
		"function": map[string]any{
			"name":        "query_clivo_users",
			"description": "Get the total number of users registered on Clivo",
			"parameters": map[string]any{
				"type":       "object",
				"properties": map[string]any{},
			},
		},
	},
}