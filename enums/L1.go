package main

import (
	"fmt"
)

type email struct {
	status    string
	recipient *user
}

type user struct {
	email  string
	status string
}

type analytics struct {
	totalBounces int
}

func (u *user) updateStatus(status string) error {
	if status != "email_bounced" {
		return fmt.Errorf("invalid status: %s", status)
	}
	u.status = status
	return nil
}

func (a *analytics) track(event string) error {
	if event != "email_bounced" {
		return fmt.Errorf("invalid event: %s", event)
	}
	a.totalBounces++
	return nil
}

func (a *analytics) handleEmailBounce(em email) error {
	errUS := em.recipient.updateStatus(em.status)
	if errUS != nil {
		return fmt.Errorf("error updating user status: %w", errUS)
	}

	errT := a.track(em.status)
	if errT != nil {
		return fmt.Errorf("error tracking user bounce: %w", errT)
	}

	return nil
}

func main() {
	// Create a user and analytics instance
	u := &user{
		email:  "test@example.com",
		status: "active",
	}
	a := &analytics{}

	// Create an email with a valid bounce status
	em1 := email{
		status:    "email_bounced",
		recipient: u,
	}

	// Handle valid bounce - should succeed
	err := a.handleEmailBounce(em1)
	if err != nil {
		fmt.Println("Failed to handle bounce:", err)
	} else {
		fmt.Println("Handled bounce successfully")
		fmt.Println("User status:", u.status)
		fmt.Println("Total bounces tracked:", a.totalBounces)
	}

	// Create an email with invalid bounce status
	em2 := email{
		status:    "invalid_status",
		recipient: u,
	}

	// Handle invalid bounce - should fail on updateStatus
	err = a.handleEmailBounce(em2)
	if err!=nil {
		fmt.Println("Failed to handle bounce:", err)
	}
}

