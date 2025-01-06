package model

import "time"

type Message struct {
	ID string `json:"id"`
	TopicName string `json:"topicName"`
	Text string `json:"text"`
	TimeStamp time.Time `json:"timeStamp"`
}



type Role string

const(
	AdminRole Role = "ADMIN"
	UserRole Role = "USER"
)

type User struct {
	Name string
	Role Role
	SubTopics map[string]bool
}

type Notification struct {
	TopicName string
	Message string
	SentTo User
}

type Topic struct {
	TopicName string
	UsersSubscribed []User
}

