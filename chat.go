package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 管理用戶websocket
type connection struct {
	ws map[int]*websocket.Conn
}

var cron connection

var socket = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	cron.ws = make(map[int]*websocket.Conn)
}

// 首頁
func showIndex(c *gin.Context) {
	reade(c, "index.html", gin.H{})
}

// 註冊websocket
func registerWs(c *gin.Context) {
	ws, err := socket.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	cron.add(1, ws)
}

// 新增用戶websocket
func (c connection) add(id int, ws *websocket.Conn) {
	c.ws[id] = ws
}
