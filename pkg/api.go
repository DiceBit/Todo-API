package pkg

import (
	"context"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"sync"
	"time"
	"todo-api/pkg/db"
	"todo-api/pkg/db/sqlite"
)

type API struct {
	router   *mux.Router
	Conn     db.DBInterface
	stopChan chan struct{}
	wg       sync.WaitGroup
}

func NewAPI() *API {
	api := API{
		router:   mux.NewRouter(),
		Conn:     sqlite.NewConn(),
		stopChan: make(chan struct{}),
	}
	return &api
}

func (api *API) Endpoints() {
	api.router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			next.ServeHTTP(w, r)
		})
	})

	api.router.HandleFunc("/tasks", api.CreateTasks).Methods(http.MethodPost)
	api.router.HandleFunc("/tasks", api.GetTasks).Methods(http.MethodGet)
	api.router.HandleFunc("/tasks/{id}", api.PutTasks).Methods(http.MethodPut)
	api.router.HandleFunc("/tasks/{id}", api.DeleteTasks).Methods(http.MethodDelete)
	api.router.HandleFunc("/tasks/{id}/complete", api.CompleteTask).Methods(http.MethodPatch)
}

func (api *API) CheckFunc(ctx context.Context, d time.Duration) {
	defer api.wg.Done()

	ticket := time.NewTicker(d)
	defer ticket.Stop()

	for {
		select {
		case <-api.stopChan:
			log.Println("Background task is stopped")
			return
		case <-ticket.C:
			log.Println("Starting goroutine")
			if err := api.Conn.CheckTasks(ctx); err != nil {
				log.Println(err)
			}
		default:
		}
	}
}
