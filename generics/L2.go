package main

import (
	"errors"
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	if newItem.GetCost()>balance {
		return nil,0.0,errors.New("insufficient funds")
	}
	balance-=newItem.GetCost();

	oldItems = append(oldItems, newItem)

	return oldItems,balance,nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}


func main() {
	var balance float64 = 50.0
	var subs []subscription

	// Create a monthly subscription
	sub := subscription{
		userEmail: "test@mail.com",
		startDate: time.Now(),
		interval:  "monthly",
	}

	// Charge subscription (works because T is subscription)
	subs, balance, err := chargeForLineItem(sub, subs, balance)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Charged:", sub.GetName())
		fmt.Println("Remaining balance:", balance)
	}

	// Now do a separate charge for a different type
	var usages []oneTimeUsagePlan

	usage := oneTimeUsagePlan{
		userEmail:        "test@mail.com",
		numEmailsAllowed: 2000, // Costs 60.0
	}

	usages, balance, err = chargeForLineItem(usage, usages, balance)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Charged:", usage.GetName())
		fmt.Println("Remaining balance:", balance)
	}

	// Print all charged items (must do it separately for each slice)
	fmt.Println("\nAll charged subscriptions:")
	for _, item := range subs {
		fmt.Printf("- %s ($%.2f)\n", item.GetName(), item.GetCost())
	}

	fmt.Println("\nAll charged usage plans:")
	for _, item := range usages {
		fmt.Printf("- %s ($%.2f)\n", item.GetName(), item.GetCost())
	}
}
