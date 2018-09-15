package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 聯絡人結構
type person struct {
	Name string
	User []string
}

type ContactBody struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

// 各用戶聯絡人關聯集合
type contact []person

var contactRelated contact

func init() {
	contactRelated.add("test", "guest")
}

// 新增聯絡人
func addContact(c *gin.Context) {
	var body ContactBody

	c.ShouldBindJSON(&body)

	m, ok := member.get(body.Contact)

	if (ok && contactRelated.add(body.Name, m.Name)) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "找不到用戶",
		})
	}
}

// 取聯絡人
func getContact(c *gin.Context) {
	n, _ := c.GetQuery("name")

	person, _, ok := contactRelated.get(n)

	if (ok) {
		c.JSON(http.StatusOK, person.User)
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

// 新增聯絡人
func (c contact) add(name, user string) bool {
	_, i, ok := c.get(name)

	if (! ok) {
		contactRelated = append(contactRelated, person{
			Name: name,
			User: []string{user},
		})
	} else {
		contactRelated[i].User = append(contactRelated[i].User, user)
	}

	return true
}

// 取出聯絡人
func (c contact) get(name string) (person, int, bool) {
	for i, p := range contactRelated {
		if p.Name == name {
			return p, i, true
		}
	}

	return person{}, 0, false
}
