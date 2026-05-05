package transaction

import (
	"errors"
	"fmt"
	"time"
)

type Transaction struct{
	ID, Name, Status string
	Amount float64
}

func NewTransaction(name string, amount float64) Transaction{
	id := fmt.Sprintf("TX-%d",time.Now().UnixNano())
	return Transaction{
		ID: id,
		Name: name,
		Amount: amount,
		Status: "pending",
	}
}

func (payment Transaction) Validate() error{
	if payment.Name == ""{
		return errors.New("invalid transaction: name is empty.")
	}

	if payment.Amount <= 0 {
		return errors.New("Amount cannot be a negative value.")
	}
	return nil
}


func (payment *Transaction) Process(){
	payment.Status = "success"
}

func (payment Transaction) Summary() string{
	return fmt.Sprintf("ID: %s \nName: %s \nAmount: %.2f \nStatus: %s", 
	payment.ID, payment.Name, payment.Amount, payment.Status)	
}

