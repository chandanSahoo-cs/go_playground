/*
Implement the updateBalance function. It should take a customer pointer 
and a transaction, and return an error. Depending on the transactionType, 
it should either add or subtract the amount from the customers balance. 
If the customer does not have enough money, it should return the error 
insufficient funds. If the transactionType isn't a withdrawal or deposit,
it should return the error unknown transaction type. Otherwise, it should
process the transaction and return nil.
*/
package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(c *customer, t transaction) error {
	
	if(t.transactionType!=transactionDeposit && t.transactionType!=transactionWithdrawal){
		return errors.New("unknown transaction type")
	}

	if (t.transactionType==transactionWithdrawal){
		if(t.amount>c.balance){
			return errors.New("insufficient funds")
		}
		c.balance-=t.amount
	}else{
		c.balance+=t.amount
	}
	return nil
}


func main() {
	c := &customer{id: 1, balance: 100.0}

	// Deposit
	err := updateBalance(c, transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit})
	fmt.Println("Deposit:", err, "New Balance:", c.balance)

	// Withdrawal success
	err = updateBalance(c, transaction{customerID: 1, amount: 30.0, transactionType: transactionWithdrawal})
	fmt.Println("Withdrawal:", err, "New Balance:", c.balance)

	// Withdrawal fail
	err = updateBalance(c, transaction{customerID: 1, amount: 150.0, transactionType: transactionWithdrawal})
	fmt.Println("Overdraw Attempt:", err, "New Balance:", c.balance)

	// Invalid transaction
	err = updateBalance(c, transaction{customerID: 1, amount: 10.0, transactionType: "transfer"})
	fmt.Println("Invalid Type:", err, "New Balance:", c.balance)
}

/*
Sample Output:
Deposit: <nil> New Balance: 150
Withdrawal: <nil> New Balance: 120
Overdraw Attempt: insufficient funds New Balance: 120
Invalid Type: unknown transaction type New Balance: 120
*/