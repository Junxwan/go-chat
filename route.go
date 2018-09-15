package main

import (
	"github.com/gin-gonic/gin"
	"go-chat/auth"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("view/*")

	router.Use(auth.CheckLogin())

	router.Static("/images", "./images")

	router.GET("/", auth.CheckPermission(), showIndex)
	router.GET("/login", auth.ShowLogin)
	router.POST("/login", auth.Attempt)
	router.GET("/register", auth.ShowRegister)
	router.POST("/register", auth.Register)
	router.GET("/chat", auth.CheckPermission(), getChat)
	router.POST("/contact", auth.CheckPermission(), addContact)
	router.GET("/contact", getContact)

	return router
}
