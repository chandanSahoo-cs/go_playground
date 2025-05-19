package main

import (
	"fmt"
)

type biller[C customer] interface {
	Charge(C)
	Name()
}

// don't edit below this line

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}


func main() {
	// Create some customers
	u := user{UserEmail: "user@example.com"}
	o := org{
		Admin: user{UserEmail: "admin@org.com"},
		Name:  "MyOrg",
	}

	// Create billers with different plans
	userBill := userBiller{Plan: "basic"}
	orgBill := orgBiller{Plan: "pro"}

	// Use userBiller to charge user
	userBillResult := userBill.Charge(u)
	fmt.Printf("User billed: %s, amount: %.2f\n", userBillResult.Customer.GetBillingEmail(), userBillResult.Amount)
	fmt.Println("Biller name:", userBill.Name())

	// Use orgBiller to charge org
	orgBillResult := orgBill.Charge(o)
	fmt.Printf("Org billed: %s, amount: %.2f\n", orgBillResult.Customer.GetBillingEmail(), orgBillResult.Amount)
	fmt.Println("Biller name:", orgBill.Name())
}
