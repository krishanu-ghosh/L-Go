package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
)

type Transaction struct {
	Amount float64 `json:"amount"`
}

var Balance float64 = 0

func main() {
	routes := gin.Default()
	RouteHandler(routes)
	err := routes.Run(":9080")
	if err != nil {
		return
	}
}

func RouteHandler(r *gin.Engine) {
	r.POST("/deposit", DepositHandler)
	r.POST("/withdraw", WithdrawHandler)
	r.GET("/balance", BalanceHandler)

}

func BalanceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Current Balance :": Balance})
}

func DepositHandler(c *gin.Context) {
	var transaction Transaction
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		return
	}

	if transaction.Amount < 0 {
		c.JSON(http.StatusOK, "Invalid transaction amount")
	} else {
		Balance = Balance + transaction.Amount
		c.JSON(http.StatusOK, Balance)
		SendTransactionEmail([]string{"krishanughosh1512@gmail.com"}, "credited to", transaction.Amount, Balance)
	}
}

func WithdrawHandler(c *gin.Context) {
	var transaction Transaction
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		return
	}

	if Balance < transaction.Amount {
		c.JSON(http.StatusOK, "Low balance. Unable to withdraw")
	} else {
		Balance = Balance - transaction.Amount
		c.JSON(http.StatusOK, Balance)
		SendTransactionEmail([]string{"krishanughosh1512@gmail.com"}, "debited from", transaction.Amount, Balance)
	}
}

func SendTransactionEmail(receiverEmail []string, typeoftransaction string, amount float64, balance float64) {
	var to = receiverEmail
	// Set up authentication information
	auth := smtp.PlainAuth("", "kghosh@datopic.com", "Kr15h4nu", "smtp.gmail.com")

	// Set up the message
	from := "kghosh@datopic.com"
	subject := "Transaction notification"
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nFrom: %s\r\nMIME-version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n", to, subject, from)

	// Set up the HTML content.
	html := fmt.Sprint("Amount of Rs ", amount, " has been ", typeoftransaction, " your account. Remaining A/C balance is Rs ", balance)
	message += html

	// Send the email via SMTP.
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Transaction email sent to", to[0])
}
