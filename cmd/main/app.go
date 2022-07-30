package main

import (
	"RestAPI/internal/user"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("Create router")
	router := gin.Default()

	log.Println("Create handler")
	handler := user.NewHandler()

	log.Println("register router")
	handler.Register(router)

	log.Println("Server started in http://localhost:8080")
	start(router)
}

func start(router *gin.Engine) {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
	}

	server := http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}
