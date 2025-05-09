/*
Implement the importance methods for each message type. They should return the importance score for each message type.
For a directMessage the importance score is based on if the message isUrgent or not. If it is urgent, the importance score is 50 otherwise the importance score is equal to the DM's priorityLevel.
For a groupMessage the importance score is based on the message's priorityLevel
All systemAlert messages should return a 100 importance score.
Complete the processNotification function. It should identify the type and return different values for each type
For a directMessage, return the sender's username and importance score.
For a groupMessage, return the group's name and the importance score.
For a systemAlert, return the alert code and the importance score.
If the notification does not match any known type, return an empty string and a score of 0.
*/
package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int{
	if dm.isUrgent{
		return 50
	}
	return dm.priorityLevel
}

func (groupM groupMessage) importance() int{
	return groupM.priorityLevel
}

func (sa systemAlert) importance() int{
	return 100
}

func processNotification(n notification) (string, int) {
	switch not := n.(type){
	case directMessage :
		return not.senderUsername,not.importance()
	case groupMessage:
		return not.groupName,not.importance()
	case systemAlert:
		return not.alertCode,not.importance()
	}
	return "",0
}

func main() {
	notifications := []notification{
		directMessage{
			senderUsername: "alice",
			messageContent: "Please review my PR",
			priorityLevel:  20,
			isUrgent:       true,
		},
		directMessage{
			senderUsername: "bob",
			messageContent: "Lunch?",
			priorityLevel:  10,
			isUrgent:       false,
		},
		groupMessage{
			groupName:      "devs",
			messageContent: "Standup at 10am",
			priorityLevel:  30,
		},
		systemAlert{
			alertCode:      "ERR_500",
			messageContent: "Server down!",
		},
	}

	for _, n := range notifications {
		id, score := processNotification(n)
		fmt.Printf("ID: %s | Importance: %d\n", id, score)
	}
}

/*
Sample Output:
ID: alice | Importance: 50
ID: bob | Importance: 10
ID: devs | Importance: 30
ID: ERR_500 | Importance: 100
*/