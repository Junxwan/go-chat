package main

import (
	"github.com/appleboy/gofight"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func initial() {
	contactRelated = []person{}
	member = []user{}
}

func TestGetContact(t *testing.T) {
	initial()

	contactRelated = append(contactRelated, person{
		Name: "test",
		User: []string{"1"},
	})

	c, _, ok := contactRelated.get("test")

	assert.True(t, ok)
	assert.Equal(t, "test", c.Name)
}

func TestAddContact(t *testing.T) {
	initial()

	ok := contactRelated.add("name", "user")

	assert.True(t, ok)
	assert.Equal(t, contactRelated[0], person{
		Name: "name",
		User: []string{"user"},
	})
}

func TestGetContactApi(t *testing.T) {
	initial()

	contactRelated = append(contactRelated, person{
		Name: "add",
		User: []string{"sam"},
	})

	getContactApi(t, func(t *testing.T, r gofight.HTTPResponse) {
		assert.JSONEq(t, `["sam"]`, r.Body.String())
		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestGetContactApiNotFound(t *testing.T) {
	initial()

	getContactApi(t, func(t *testing.T, r gofight.HTTPResponse) {
		assert.JSONEq(t, `{}`, r.Body.String())
		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestAddContactApi(t *testing.T) {
	initial()

	contactRelated = append(contactRelated, person{
		Name: "junx",
		User: []string{"sam", "add"},
	})

	member = append(member, user{
		Name:     "add",
		Account:  "test@gmail.com",
		Password: "123456",
	})

	addContactApi(t, func(t *testing.T, r gofight.HTTPResponse) {
		assert.JSONEq(t, `{"msg":"成功"}`, r.Body.String())
		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestAddContactApiNotFound(t *testing.T) {
	initial()

	addContactApi(t, func(t *testing.T, r gofight.HTTPResponse) {
		assert.JSONEq(t, `{"msg": "找不到用戶"}`, r.Body.String())
		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func addContactApi(t *testing.T, f func(t *testing.T, r gofight.HTTPResponse)) {
	r := gofight.New()

	r.POST("/contact").
		SetDebug(false).
		SetJSON(gofight.D{
			"name":    "sam",
			"contact": "add",
		}).
		SetCookie(gofight.H{
			"login": "123445",
		}).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			f(t, r)
		})
}

func getContactApi(t *testing.T, f func(t *testing.T, r gofight.HTTPResponse)) {
	r := gofight.New()

	r.GET("/contact").
		SetDebug(false).
		SetQuery(gofight.H{
			"name": "add",
		}).
		SetCookie(gofight.H{
			"login": "123445",
		}).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			f(t, r)
		})
}
