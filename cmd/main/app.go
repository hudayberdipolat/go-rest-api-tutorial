package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/hudayberdipolat/rest-api-tutorial/internal/user"
	"github.com/hudayberdipolat/rest-api-tutorial/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("-------------------------------------------------")
	logger.Info("create router")
	router := httprouter.New()
	logger.Info("register user handler")
	handler := user.NewUserHandler(logger)
	handler.Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info("server is listening port 127.0.0.1:3000")
	log.Fatal(server.Serve(listener))
}
