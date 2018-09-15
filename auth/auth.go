package auth

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
	Member.add("test", "test@gmai.com", "123456")
	Member.add("guest", "test2@gmai.com", "123456")
}

// 登入頁
func ShowLogin(c *gin.Context) {
	reade(c, "login.html", gin.H{})
}

// 註冊頁
func ShowRegister(c *gin.Context) {
	reade(c, "register.html", gin.H{})
}

// 嘗試登入
func Attempt(c *gin.Context) {
	var form loginForm
	name := ""

	c.ShouldBind(&form)

	if (login(c, form.Account, form.Password)) {
		m, _ := Member.GetByAccount(form.Account)
		name = m.Name
	}

	reade(c, "login.html", gin.H{
		"name": name,
	})
}

// 登入
func login(c *gin.Context, account, password string) bool {
	if (! Member.exist(account, password)) {
		return false
	}

	c.SetCookie("login", strconv.FormatInt(rand.Int63(), 20), 3600, "", "", false, true)

	c.Set("isLogin", true)

	return true
}

// 註冊
func Register(c *gin.Context) {
	var form User
	message := ""

	if err := c.ShouldBind(&form); err == nil {
		Member.add(form.Name, form.Account, form.Password)

		message = "恭喜你註冊成功，請前往登入頁做登入"
	} else {
		message = "帳號密碼輸入有誤請重新填寫"
	}

	reade(c, "login.html", gin.H{
		"message": message,
	})
}

// 檢查操作權限
func CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		isLogin, _ := c.Get("isLogin")

		if (! isLogin.(bool)) {
			c.Redirect(http.StatusFound, "/login")
		}
	}
}

// 檢查是否已登入
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("login"); err == nil && token != "" {
			c.Set("isLogin", true)
		} else {
			c.Set("isLogin", false)
		}
	}
}

// 讀取view
func reade(c *gin.Context, view string, data gin.H) {
	isLogin, _ := c.Get("isLogin")

	data["isLogin"] = isLogin.(bool)

	c.HTML(http.StatusOK, view, data)
}
