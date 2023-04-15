package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/num", getNum)
	router.POST("/add", addNum)
	router.POST("/mul", mulNum)

	router.Run("localhost:8080")
}

type Data struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type Result struct {
	Result int `json:"result"`
}

var num1 int = 4

func getNum(c *gin.Context) {
	c.JSON(http.StatusOK, num1)
}

func addNum(c *gin.Context) {
	var newNum Data
	c.BindJSON(&newNum)
	add := newNum.Num1 + newNum.Num2
	// resp := Result{Result: add}
	fmt.Printf("%x", add)
	c.JSON(http.StatusOK, add)
}

func mulNum(c *gin.Context) {
	var newNum Data
	c.BindJSON(&newNum)
	add := newNum.Num1 * newNum.Num2
	// resp := Result{Result: add}
	fmt.Printf("%x", add)
	c.JSON(http.StatusOK, add)
}
