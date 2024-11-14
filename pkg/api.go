package pkg

import (
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"todo-api/pkg/db/sqlite"
)

type API struct {
	router *mux.Router
	conn   *sqlite.Sqlite
}

func NewAPI() *API {
	api := API{
		router: mux.NewRouter(),
		conn:   sqlite.NewConn(),
	}
	return &api
}

func (api *API) Endpoints() {
	api.router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("test"))
	}).Methods(http.MethodGet)

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
