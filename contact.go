package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type addContactJson struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// 聯絡人結構
type person struct {
	id   int
	room []int
}

type contact []person

var contacts contact

func init() {
	contacts.add(1, 2)
}

func getContact(c *gin.Context) {
	room := rooms{}

	id, _ := c.GetQuery("id")
	userId, _ := strconv.Atoi(id)

	ct, _, _ := contacts.get(userId)

	for _, v := range ct.room {
		r, _ := chat.get(v)

		room = append(room, r)
	}

	c.JSON(http.StatusOK, room)
}

func addContact(c *gin.Context) {
	var form addContactJson

	c.ShouldBindJSON(&form)

	m, _ := member.getByName(form.Name)

	id, _ := strconv.Atoi(form.ID)

	contacts.add(id, m.ID)

	c.JSON(http.StatusOK, gin.H{
		"msg": "成功",
	})
}

// 新增聯絡人
func (p contact) add(id, contactId int) error {
	roomId, err := newRoom(id, contactId)

	if (err != nil) {
		return err
	}

	_, i, err := p.get(id)

	if (err != nil) {
		contacts = append(contacts, person{
			id:   id,
			room: []int{roomId},
		})
	} else {
		contacts[i].room = append(contacts[i].room, roomId)
	}

	_, i2, err := p.get(contactId)

	if (err != nil) {
		contacts = append(contacts, person{
			id:   contactId,
			room: []int{roomId},
		})
	} else {
		contacts[i2].room = append(contacts[i2].room, roomId)
	}

	return nil
}

func (p contact) get(userId int) (person, int, error) {
	for i, v := range p {
		if (v.id == userId) {
			return v, i, nil
		}
	}

	return person{}, 0, fmt.Errorf("找不到聯絡人")
}
