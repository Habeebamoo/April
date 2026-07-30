package tools

import (
	"context"
	"fmt"
	"os"

	resend "github.com/resend/resend-go/v2"
)

func SendEmail(input map[string]any) string {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	to, _ := input["to"].(string)
	subject, _ := input["subject"].(string)
	body, _ := input["body"].(string)

	_, err := client.Emails.SendWithContext(context.TODO(), &resend.SendEmailRequest{
		From:    "Habeeb via April <hello@myclivo.com>",
		To:      []string{to},
		Subject: subject,
		Html:    body,
	})

	if err != nil {
		return fmt.Sprintf("Failed to send email: %s", err.Error())
	}

	return fmt.Sprintf("Email sent to %s", to)
}