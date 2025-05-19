package main

import (
	"fmt"
)

type emailStatus int

const (
	emailBounced  emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)

func main() {
	statuses := []emailStatus{emailBounced, emailInvalid, emailDelivered, emailOpened}

	for _, status := range statuses {
		fmt.Printf("Status value: %d, name: %s\n", status, status.String())
	}
}

func (es emailStatus) String() string {
	switch es {
	case emailBounced:
		return "emailBounced"
	case emailInvalid:
		return "emailInvalid"
	case emailDelivered:
		return "emailDelivered"
	case emailOpened:
		return "emailOpened"
	default:
		return "unknown"
	}
}