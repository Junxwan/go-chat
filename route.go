package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() {
	router.Use(checkLogin())

	router.GET("/", checkPermission(), showIndex)
	router.GET("/login", showLogin)
	router.POST("/login", attempt)
	router.GET("/register", showRegister)
	router.POST("/register", register)
	router.GET("/chat", checkPermission(), getChat)

	router.Run()
}

// 讀取view
func reade(c *gin.Context, view string, data gin.H) {
	isLogin, _ := c.Get("isLogin")

	data["isLogin"] = isLogin.(bool)

	c.HTML(http.StatusOK, view, data)
}
