package main

import (
	"net/http"
	"os"
	"fmt"
	"strconv"
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
	rooms = make(map[socketio.Socket]string)
	names = make(map[socketio.Socket]string)
	sockets = make([]socketio.Socket,0)
	router.GET  ("/socket.io/", socketHandler)
	router.POST ("/socket.io/", socketHandler)
	router.Handle ("WS",  "/socket.io/", socketHandler)
	router.Handle ("WSS", "/socket.io/", socketHandler)
	router.Static("/fake/", "./assets/")

	//Run Server
	router.Run(":" + port)
}

var sockets []socketio.Socket
var rooms map[socketio.Socket]string
var names map[socketio.Socket]string

func socketHandler (c  *gin.Context) {
    Socketio_Server.On("connection", func(socket socketio.Socket) {
				// Initialize Socket
				sockets = append(sockets, socket)
				rooms[socket] = "galaxy"
				names[socket] = "wuffle"
				socket.Join(rooms[socket])
				// Socket Receiver
				socket.On("ping", func(msg string){
						Socketio_Server.BroadcastTo(rooms[socket], "pong", msg)
				})
				socket.On("status", func(msg string){
						// res := "\nCurrent room: " + rooms[socket]
						// res = res + "\n # of rooms: " + strconv.Itoa(len(rooms))
						// res = res + "\n # of names: " + strconv.Itoa(len(names))
						res := ""
						res = res + "\n # of sockets: " + strconv.Itoa(len(sockets))
						for _, x := range sockets{
								if x == socket {
										res = res + "\nYour info: <" + names[x] + "> is in <" + rooms[x] + ">"
								}else {
										res = res + "\n<" + names[x] + "> is in <" + rooms[x] + ">"
								}
						}
						socket.Emit("pipe", res)
				})
				socket.On("connect to room", func(roomName string){
						socket.Leave(rooms[socket])
						rooms[socket] = roomName
						socket.Join(rooms[socket])
						fmt.Println("connec to room: ", rooms[socket])
				})
				socket.On("update username", func(username string){
						names[socket] = username
						fmt.Println("Update user name: ", names[socket])
				})
        socket.On("disconnection", func() {
						fmt.Println(names[socket], " left ", rooms[socket])
						socket.Leave(rooms[socket])
						delete(rooms, socket)
						delete(names, socket)
						for index, x := range sockets{
								if x == socket {
										sockets = append(sockets[:index], sockets[index+1:]...)
										break;
								}
						}
        })
    })
    Socketio_Server.On ( "error", func(socket socketio.Socket, err error) {
        fmt.Printf ( "[ WebSocket ] Error : %v", err.Error () )
    })
    Socketio_Server.ServeHTTP ( c.Writer, c.Request )
}
