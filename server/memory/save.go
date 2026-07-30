package memory

import (
	"log"

	"github.com/Habeebamoo/April/server/db"
)

func Save(userMessage string, reply string) {
	memory := Memory{
		Content: "User said: '" + userMessage + "' | April replied: '" + reply + "'",
	}

	result := db.DB.Create(&memory)
	if result.Error != nil {
		log.Println("Failed to save memory:", result.Error)
	}
}

func SaveSession(summary string) {
	session := Session{
		Summary: summary,
	}

	result := db.DB.Create(&session)
	if result.Error != nil {
		log.Println("Failed to save session:", result.Error)
	}
}