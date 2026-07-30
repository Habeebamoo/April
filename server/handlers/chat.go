package handlers

import (
	"net/http"

	"github.com/Habeebamoo/April/server/agent"
	"github.com/Habeebamoo/April/server/memory"
	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}

func ChatHandler(c *gin.Context) {
	var req ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message is required"})
		return
	}

	// Load memories
	memories := memory.Load()

	// Run agent
	reply := agent.Run(req.Message, memories)

	// Save memory in background
	go memory.Save(req.Message, reply)

	c.JSON(http.StatusOK, ChatResponse{Reply: reply})
}