package main

import (
	"gopayments/api"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.POST("/payment", api.ProcessPayment)

	r.Run(":8080")

}

