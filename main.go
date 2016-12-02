package main

import (
	"net/http"
	"os"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
)

var Socketio_Server * socketio.Server

func main() {
	// Connect to defaul port on server
	port := os.Getenv("PORT")
	// If not specified, use port 5000
	if port == "" {port = "5000"}

	// Initialize Gin server
	router := gin.New()
	// Initialize Gin logger
	router.Use(gin.Logger())
	// Load all html in templates directory
	router.LoadHTMLGlob("templates/*.html")

	// REST API
	router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
	})

	// Socket.io Setup
	Socketio_Server, _ = socketio.NewServer(nil)
	router.GET  ("/socket.io/", socketHandler)
	router.POST ("/socket.io/", socketHandler)
	router.Handle ("WS",  "/socket.io/", socketHandler)
	router.Handle ("WSS", "/socket.io/", socketHandler)
	router.Static("/fake/", "./assets/")

	//Run Server
	router.Run(":" + port)
}

func socketHandler (c  *gin.Context) {
    Socketio_Server.On("connection", func(socket socketio.Socket) {
        fmt.Println("on connection")
        socket.Join("galaxy")
				socket.On("ping", func(msg string){
						fmt.Println(msg, " pinging")
						msg = time.Now().Format(time.RFC850)
						socket.BroadcastTo("galaxy", "pong", msg)
				})
        socket.On("chat message", func(msg string) {
            fmt.Println("emit:", socket.Emit("chat message", msg))
            socket.BroadcastTo("galaxy", "chat message", msg)
        })
        socket.On("disconnection", func() {
            fmt.Println("on disconnect")
        })
    })
    Socketio_Server.On ( "error", func(socket socketio.Socket, err error) {
        fmt.Printf ( "[ WebSocket ] Error : %v", err.Error () )
    })
    Socketio_Server.ServeHTTP ( c.Writer, c.Request )
}
