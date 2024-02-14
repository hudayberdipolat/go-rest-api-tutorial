package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/hudayberdipolat/rest-api-tutorial/internal/user"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("register user handler")
	handler := user.NewUserHandler()
	handler.Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	log.Println("start application")
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("server is listening port 127.0.0.1:3000")
	log.Fatal(server.Serve(listener))
}
