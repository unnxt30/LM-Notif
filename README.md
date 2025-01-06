# LM-Notif

## Installation
1. Clone the repository:
```bash
git clone https://github.com/unnxt30/LM-Notif.git
cd LM-Notif
```

2. Build the application:
```bash
go build
```

## Usage

1. Start the application:
```bash
go run main.go
```

2. Available Commands:
- Add a user:
```bash
addUser [userName] [role]     # role must be either USER or ADMIN
```

- Get all users:
```bash
getUsers
```

- Add a topic:
```bash
addTopic [topicName] [caller] # caller must be an ADMIN user
```

- Get all topics:
```bash
getTopics
```

- Subscribe a user to a topic:
```bash
subscribeTopic [topicName] [userName]
```

- Publish a message:
```bash
publishMessage '[JSON_STRING]'
```

- View Subscribed topics for a User:
```bash
viewSubscribedTopics [userName]
```

- Remove a user (admin role)
```bash
removeUser [userName] [caller]
```

- Remove a topic (admin role)
```bash
removeTopic [topicName] [userName]
```

### ⚠️ Important Note for publishMessage
The JSON string must be enclosed in single quotes. The format should be:
```bash
publishMessage '{"id":"1", "topicName":"foo", "text":"bar"}'
```

### ⚠️ Important Note for publishMessage
The Timestamp functionality doesn't work right now, please refrain from using it.

3. Exit the application:
```bash
quit
```
or
```bash
exit
```