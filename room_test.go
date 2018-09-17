package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMember() (int, int) {
	chat = rooms{}

	userId1 := member.add("name1", "name1@gmail.com", "123456")
	userId2 := member.add("name2", "name2@gmail.com", "123456")

	return userId1, userId2
}

func TestNewRoom(t *testing.T) {
	userId1, userId2 := createMember()

	roomId, err := newRoom(userId1, userId2)

	assert.Equal(t, 1, roomId)
	assert.Equal(t, nil, err)
}

func TestGetRoom(t *testing.T) {
	userId1, userId2 := createMember()

	roomId, _ := newRoom(userId1, userId2)

	room, ok := chat.get(roomId)

	assert.True(t, ok)
	assert.Equal(t, room.ID, 1)
	assert.Equal(t, room.Message, []message{})
	assert.Equal(t, room.User, []roomUser{
		roomUser{
			ID:   userId1,
			Name: "name1",
		},
		roomUser{
			ID:   userId2,
			Name: "name2",
		}})
}

func TestGetUser(t *testing.T) {
	userId1, userId2 := createMember()

	roomId, _ := newRoom(userId1, userId2)

	u, ok := chat.getUser(roomId, userId1)

	assert.True(t, ok)
	assert.Equal(t, u, roomUser{
		ID:   userId2,
		Name: "name2",
	})
}
