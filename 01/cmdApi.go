package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type CommandRequest struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func main() {
	router := gin.Default()
	router.POST("/run", runCommandHandler)
	router.GET("/kill", runKill)
	err := router.Run(":1960")
	if err != nil {
		return
	}
}

func runCommandHandler(c *gin.Context) {
	var req CommandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	cmd := exec.Command(req.Command, req.Args...)
	output, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, string(output))
}

func runKill(c *gin.Context) {
	cmd := exec.Command("shutdown", "-h", "now")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.String(http.StatusOK, string(output))
}
