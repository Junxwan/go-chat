package main

import "fmt"

type room struct {
	ID      int
	Message []message
	User    map[int]string
}

type rooms map[int]room

var roomAll = make(rooms, 100)

// 建立一個聊天房間
func newRoom(userId1, userId2 int) (int, error) {
	m1, ok1 := member.get(userId1)

	if (! ok1) {
		return 0, fmt.Errorf("找不到ID用戶:%d", userId1)
	}

	m2, ok2 := member.get(userId2)

	if (! ok2) {
		return 0, fmt.Errorf("找不到ID用戶:%d", userId2)
	}

	userId := make(map[int]string)
	userId[userId1] = m1.Name
	userId[userId2] = m2.Name

	id := userId1 + userId1

	roomAll[id] = room{
		ID:      id,
		Message: []message{},
		User:    userId,
	}

	return id, nil
}

// 取房間
func (r rooms) get(id int) room {
	return roomAll[id]
}
