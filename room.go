package main

import "fmt"

type room struct {
	ID      int
	User    []roomUser
	Message []message
}

type roomUser struct {
	ID   int
	Name string
}

type message struct {
	ID  int
	Msg string
}

type rooms []room

var chat rooms

func newRoom(userId1, userId2 int) (int, error) {
	m1, ok1 := member.get(userId1)

	if (! ok1) {
		return 0, fmt.Errorf("找不到ID用戶:%d", userId1)
	}

	m2, ok2 := member.get(userId2)

	if (! ok2) {
		return 0, fmt.Errorf("找不到ID用戶:%d", userId2)
	}

	roomId := len(chat) + 1

	chat = append(chat, room{
		ID: roomId,
		User: []roomUser{
			roomUser{
				ID:   userId1,
				Name: m1.Name,
			},
			roomUser{
				ID:   userId2,
				Name: m2.Name,
			}},
		Message: []message{},
	})

	return roomId, nil
}

func (r rooms) get(roomId int) (room, bool) {
	for _, v := range r {
		if v.ID == roomId {
			return v, true
		}
	}

	return room{}, false
}

func (r rooms) getUser(roomId, masterId int) (roomUser, bool) {
	user := roomUser{}

	room, _ := r.get(roomId)

	for _, v := range room.User {
		if v.ID != masterId {
			return v, true
		}
	}

	return user, false
}
