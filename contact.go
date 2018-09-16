package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 聯絡人結構
type person struct {
	id   int
	room int
}

type contact []person

var contacts contact

func init() {
	contacts.add(1, 2)
}

func getContact(c *gin.Context) {
	id, _ := c.GetQuery("id")
	userId, _ := strconv.Atoi(id)
	ct, _ := contacts.getContact(userId)

	c.JSON(http.StatusOK, ct)
}

// 新增聯絡人
func (p contact) add(id, contactId int) error {
	rId, err := newRoom(id, contactId)

	if (err != nil) {
		return err
	}

	contacts = append(contacts, person{
		id:   id,
		room: rId,
	})

	contacts = append(contacts, person{
		id:   contactId,
		room: rId,
	})

	return nil
}

// 取出聯絡人
func (p contact) getContact(id int) (map[int]string, bool) {
	for _, c := range p {
		if c.id == id {
			return roomAll.get(c.room).User, true
		}
	}

	return map[int]string{}, false
}
