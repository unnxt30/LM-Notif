package store

import (
	"errors"

	models "github.com/unnxt30/LM-Notif/internal/model"
)

type MemoryStore struct {
	users map[string]*models.User
	topics map[string]*models.Topic
}

var GlobalMemoryStore *MemoryStore

func init() {
	GlobalMemoryStore = &MemoryStore{
		users: make(map[string]*models.User),
		topics: make(map[string]*models.Topic),
	}
}

func (m *MemoryStore) AddUser(user *models.User) error {
	if _, ok := m.users[user.Name]; ok {
		return errors.New("user already exists")
	}

	m.users[user.Name] = user
	return nil
}

func (m *MemoryStore) GetUser(name string) (*models.User, error) {
	user, ok := m.users[name]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MemoryStore) GetAllUsers() map[string]*models.User {
	return m.users
}


func (m *MemoryStore) AddTopic(topic *models.Topic) error {
	if _, ok := m.topics[topic.TopicName]; ok {
		return errors.New("topic already exists")
	}
	m.topics[topic.TopicName] = topic
	m.topics[topic.TopicName].UsersSubscribed = make([]models.User, 0)
	return nil
}