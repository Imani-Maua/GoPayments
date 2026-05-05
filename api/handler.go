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

	var payment transaction.Transaction
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    // step 2: validate the transaction
	validate := payment.Validate()
	if validate != nil {
		fmt.Println("Valdiation failed:", validate)
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
		"transaction": payment.Summary(),
	})

}