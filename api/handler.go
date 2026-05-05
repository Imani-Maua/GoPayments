package api

import (
	"gopayments/logger"
	"gopayments/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func ProcessPayment(c *gin.Context){
	// step 1: read the JSON body into a struct

	var body struct {
		Name string
		Amount float64
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment := transaction.NewTransaction(body.Name, body.Amount)
    // step 2: validate the transaction
	validate := payment.Validate()
	if validate != nil {
		fmt.Println("Valdiation failed:", validate)
		return
	}else {
		fmt.Println("Validation passed!")
	}

    // step 3: process it
	payment.Process()

    // step 4: log it
	logger.Log(payment)

    // step 5: return a JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "payment processed successfully",
		"transaction": gin.H{
			"id": payment.ID,
			"name": payment.Name,
			"amount": payment.Amount,
			"status": payment.Status,
		},
	})

}