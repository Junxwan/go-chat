package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type message struct {
	To  int
	Msg string
}

var Chat []message

func init()  {
	Chat = append(Chat, message{To: 1, Msg: "Hello,"})
	Chat = append(Chat, message{To: 1, Msg: "it's me."})
	Chat = append(Chat, message{To: 2, Msg: "... about who we used to be."})
	Chat = append(Chat, message{To: 1, Msg: "I was wondering..."})
}

// 首頁
func showIndex(c *gin.Context) {
	reade(c, "index.html", gin.H{})
}

// 聊天室訊息
func getChat(c *gin.Context) {
	c.JSON(http.StatusOK, Chat)
}
