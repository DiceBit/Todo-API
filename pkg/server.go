package pkg

import (
	"log"
	"net/http"
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
	log.Println("Server started")
	log.Fatal(http.ListenAndServe("localhost:8080", srv.api.router))
}

func (srv *Server) StopServer() {
	srv.api.conn.CloseConn()
}

/*
TODO: Фоновая задача: добавьте горутину, которая будет периодически (например, раз в минуту)
проверять все задачи на наличие истекшего срока (due_date).
Если срок истек, обновите статус задачи, установив поле overdue в true.
Используйте канал для контроля завершения фоновой задачи при остановке приложения.
*/
