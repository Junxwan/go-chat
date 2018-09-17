package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type loginForm struct {
	Account  string `form:"account" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func init() {
	member.add("test", "test@gmai.com", "123456")
	member.add("guest", "test2@gmai.com", "123456")
	member.add("test2", "test2@gmai.com", "123456")
	member.add("test3", "test3@gmai.com", "123456")
}

// 登入頁
func showLogin(c *gin.Context) {
	reade(c, "login.html", gin.H{})
}

// 註冊頁
func showRegister(c *gin.Context) {
	reade(c, "register.html", gin.H{})
}

// 嘗試登入
func attempt(c *gin.Context) {
	var form loginForm
	id := 0

	c.ShouldBind(&form)

	if (login(c, form.Account, form.Password)) {
		m, _ := member.getByAccount(form.Account)
		id = m.ID
	}

	reade(c, "login.html", gin.H{
		"id": id,
	})
}

// 登入
func login(c *gin.Context, account, password string) bool {
	if (! member.exist(account, password)) {
		return false
	}

	c.SetCookie("login", strconv.FormatInt(rand.Int63(), 20), 3600, "", "", false, true)

	c.Set("isLogin", true)

	return true
}

// 註冊
func register(c *gin.Context) {
	var form user
	message := ""

	if err := c.ShouldBind(&form); err == nil {
		member.add(form.Name, form.Account, form.Password)

		message = "恭喜你註冊成功，請前往登入頁做登入"
	} else {
		message = "帳號密碼輸入有誤請重新填寫"
	}

	reade(c, "login.html", gin.H{
		"message": message,
	})
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
