package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"todo-api/pkg"
)

func main() {
	srv := pkg.NewSrv()
	go srv.RunServer()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	srv.StopServer()
	log.Println("Gracefully stopped")
}
