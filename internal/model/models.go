package model

type Message struct {
	ID string 
	topicName string
	text string
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

