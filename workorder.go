package main

import (
	"time"
)

type WorkOrder struct {
	Created time.Time
	Updated time.Time
	Room    int
	Guest   string
	Issue   string
	Urgency int
}

func NewWorkOrder(room, urgency int, guest, issue string) WorkOrder {
	return WorkOrder{
		Created: time.Now(),
		Updated: time.Now(),
		Room:    room,
		Guest:   guest,
		Issue:   issue,
		Urgency: urgency,
	}
}
