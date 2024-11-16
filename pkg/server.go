package pkg

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	api *API
}

func NewSrv() *Server {
	srv := Server{
		api: NewAPI(),
	}
	return &srv
}

func (srv *Server) RunServer() {
	srv.api.Endpoints()

	srv.api.wg.Add(1)
	go srv.api.CheckFunc(context.Background(), 1*time.Minute)

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", srv.api.router))
}
func (srv *Server) StopServer() {
	close(srv.api.stopChan)
	srv.api.wg.Wait()
	srv.api.Conn.CloseConn()
}
