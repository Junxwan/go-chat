package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func initial() {
	contacts = []person{}
	member = []user{}
	chat = rooms{}
}

func TestAdd(t *testing.T) {
	initial()

	userId1, userId2 := createMember()

	err := contacts.add(userId1, userId2)

	assert.Equal(t, nil, err)
	assert.Equal(t, contacts[0], person{
		id:   1,
		room: []int{1},
	})
	assert.Equal(t, contacts[1], person{
		id:   2,
		room: []int{1},
	})
}

func TestAddMany(t *testing.T) {
	initial()

	userId1, userId2 := createMember()
	userId3 := member.add("name2", "name2@gmail.com", "123456")

	contacts.add(userId1, userId2)
	contacts.add(userId1, userId3)

	c, _, _ := contacts.get(userId1)

	assert.Equal(t, c, person{
		id:   userId1,
		room: []int{1, 2},
	})

}

func TestGet(t *testing.T) {
	initial()

	p := person{
		id:   1,
		room: []int{1},
	}

	contacts = append(contacts, p)

	c, i, err := contacts.get(1)

	assert.Equal(t, nil, err)
	assert.Equal(t, 0, i)
	assert.Equal(t, person{
		id:   1,
		room: []int{1},
	}, c)
}
