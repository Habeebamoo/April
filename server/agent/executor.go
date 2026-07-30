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
	default:
		return fmt.Sprintf("Unknown tool: %s", toolName)
	}
}