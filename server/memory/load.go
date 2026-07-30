package memory

import (
	"log"

	"github.com/Habeebamoo/April/server/db"
)

type Memory struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

type Session struct {
	ID        uint   `gorm:"primaryKey"`
	Summary   string `gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

func Load() []string {
	var memories []Memory
	var sessions []Session

	// Get last 20 memories
	result := db.DB.Order("created_at desc").Limit(20).Find(&memories)
	if result.Error != nil {
		log.Println("Failed to load memories:", result.Error)
	}

	// Get last 10 session summaries
	result = db.DB.Order("created_at desc").Limit(10).Find(&sessions)
	if result.Error != nil {
		log.Println("Failed to load sessions:", result.Error)
	}

	// Combine into one slice
	var contents []string
	for _, m := range memories {
		contents = append(contents, m.Content)
	}
	for _, s := range sessions {
		contents = append(contents, "Previous session: "+s.Summary)
	}

	return contents
}