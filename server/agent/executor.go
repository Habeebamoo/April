package agent

import (
	"fmt"

	"github.com/Habeebamoo/April/server/tools"
)

func ExecuteTool(toolName string, input map[string]any) string {
	switch toolName {
	case "send_email":
		return tools.SendEmail(input)
	case "query_clivo_users":
		return tools.QueryClivoUsers()
	case "get_recent_articles_on_clivo":
			return tools.GetRecentArticles(input)
	case "get_recent_signups_on_clivo":
			return tools.GetRecentSignups(input)
	case "find_user_on_clivo":
			return tools.FindUser(input)
	case "create_comment_on_clivo":
			return tools.CreateComment(input)
	case "get_subscribers_count_on_clivo":
			return tools.GetSubscribersCount()
	case "get_user_social_stats_on_clivo":
			return tools.GetUserSocialStats(input)
	default:
		return fmt.Sprintf("Unknown tool: %s", toolName)
	}
}