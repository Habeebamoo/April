package tools

import (
	"fmt"

	"github.com/Habeebamoo/April/server/db"
)

type User struct {
	ID uint `gorm:"primaryKey"`
}

func QueryClivoUsers() string {
	var count int64
	result := db.ClivoDB.Table("users").Count(&count)
	if result.Error != nil {
		return "Failed to query Clivo users"
	}
	return fmt.Sprintf("%d", count)
}