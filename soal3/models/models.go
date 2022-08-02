package models

type User struct {
	ID    uint   `json:"id"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type Message struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Sharing string `json:"sharing"`
}

type Group struct {
	GroupName string `json:"group_name"`
	Role      string `json:"role"`
}

type MessageOnGroup struct {
	Role    string `json:"role"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Sharing string `json:"sharing"`
}
