package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	err := router.Run(":9080")
	if err != nil {
		return 
	}
}
