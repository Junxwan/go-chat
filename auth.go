package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用者結構
type user struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type member struct {
	Account []user
}

// 現已註冊的使用者
var account member

// 登入頁
func showLogin(c *gin.Context) {
	reade(c, "login.html", gin.H{})
}

// 註冊頁
func showRegister(c *gin.Context) {
	reade(c, "register.html", gin.H{})
}

// 新增會員
func (a member) addUser(username, password string) {
	account.Account = append(a.Account, user{
		Username: username,
		Password: password,
	})
}

// 檢查帳號是否正確
func (a member) isUser(username, password string) bool {
	for _, u := range a.Account {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// 檢查操作權限
func checkPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		isLogin, _ := c.Get("isLogin")

		if (! isLogin.(bool)) {
			c.Redirect(http.StatusFound, "/login")
		}
	}
}

// 檢查是否已登入
func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("login"); err == nil && token != "" {
			c.Set("isLogin", true)
		} else {
			c.Set("isLogin", false)
		}
	}
}
