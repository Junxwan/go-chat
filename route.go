package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("view/*")

	router.Use(checkLogin())

	router.Static("/images", "./images")

	router.GET("/", checkPermission(), showIndex)
	router.GET("/login", showLogin)
	router.POST("/login", attempt)
	router.GET("/register", showRegister)
	router.POST("/register", register)
	router.GET("/chat", checkPermission(), getChat)
	router.POST("/contact", checkPermission(), addContact)
	router.GET("/contact", getContact)

	return router
}
