package logger
import (
	"time"
	"gopayments/transaction"
	"fmt"
)

func Log(transact transaction.Transaction) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s]\n%s\n", timestamp, transact.Summary())
}