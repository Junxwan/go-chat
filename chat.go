package main

import "github.com/gin-gonic/gin"

// 首頁
func showIndex(c *gin.Context) {
	reade(c, "index.html", gin.H{})
}
