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
	{
		"type": "function",
		"function": map[string]any{
			"name":        "get_articles_on_clivo",
			"description": "Get all published articles on Clivo",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{},
			},
		},
	},
	{
		"type": "function",
		"function": map[string]any{
			"name":        "get_recent_signups_on_clivo",
			"description": "Get users who recently signed up on Clivo",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"limit": map[string]any{
						"type":        "number",
						"description": "How many users to return, default 5",
					},
					"period": map[string]any{
						"type":        "string",
						"description": "Time period: today, week, or month",
					},
				},
			},
		},
	},
	{
		"type": "function",
		"function": map[string]any{
			"name":        "find_user_on_clivo",
			"description": "Find a Clivo user by their name, email or username",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query": map[string]any{
						"type":        "string",
						"description": "Name, email or username to search for",
					},
				},
				"required": []string{"query"},
			},
		},
	},
	{
		"type": "function",
		"function": map[string]any{
			"name":        "create_comment_on_clivo",
			"description": "Post a comment on a Clivo article on behalf of Habeeb",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"article_id": map[string]any{
						"type":        "string",
						"description": "ID of the article to comment on",
					},
					"content": map[string]any{
						"type":        "string",
						"description": "Comment content",
					},
					"user_id": map[string]any{
						"type": "string",
						"description": "The user that you will create the comment on behalf of. DEFAULT: always use the user_id associated with @habeebamoo08",
					},
				},
				"required": []string{"article_id", "content", "user_id"},
			},
		},
	},
	{
		"type": "function",
		"function": map[string]any{
			"name":        "get_subscribers_count_on_clivo",
			"description": "Get total number of subscribers on Clivo",
			"parameters": map[string]any{
				"type":       "object",
				"properties": map[string]any{},
			},
		},
	},
	{
		"type": "function",
		"function": map[string]any{
			"name":        "get_user_social_stats_on_clivo",
			"description": "Get followers and following count of a Clivo user",
			"parameters": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"username": map[string]any{
						"type":        "string",
						"description": "Username of the user",
					},
				},
				"required": []string{"username"},
			},
		},
	},
}
