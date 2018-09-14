package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func setupRouter() {
	router = gin.Default()

	router.StaticFile("/", "./view/index.html")
	router.StaticFile("/login", "./view/login.html")
	router.StaticFile("/register", "./view/register.html")

	router.Run()
}
