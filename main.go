package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := setupRouter()

	router.Run()
}

// 讀取view
func reade(c *gin.Context, view string, data gin.H) {
	isLogin, _ := c.Get("isLogin")

	data["isLogin"] = isLogin.(bool)

	c.HTML(http.StatusOK, view, data)
}
