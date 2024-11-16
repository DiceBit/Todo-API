package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
	"todo-api/pkg/db"
)

func (api *API) CreateTasks(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if ServerInternalError(err, w) {
		return
	}

	var info db.TasksDTO
	err = json.Unmarshal(body, &info)
	if ServerInternalError(err, w) {
		return
	}

	if info.Title == "" {
		BadRequestError(errors.New("please, input title for your task"), w)
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	defer cancel()

	task, err := api.Conn.AddTask(ctx, info)
	if ServerInternalError(err, w) || RequestTimeoutError(err, w) {
		return
	}

	sendResp(task, w)
}
func (api *API) GetTasks(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	defer cancel()

	tasks, err := api.Conn.Tasks(ctx)
	if ServerInternalError(err, w) || RequestTimeoutError(err, w) {
		return
	}

	sendResp(tasks, w)
}
func (api *API) PutTasks(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	defer cancel()

	id := mux.Vars(req)["id"]

	body, err := io.ReadAll(req.Body)
	if ServerInternalError(err, w) {
		return
	}

	var info db.TasksDTO
	err = json.Unmarshal(body, &info)
	if ServerInternalError(err, w) {
		return
	}

	if info.Title == "" {
		BadRequestError(errors.New("please, input title for your task"), w)
		return
	}

	task, err := api.Conn.UpdateTask(ctx, info, id)
	if ServerInternalError(err, w) || RequestTimeoutError(err, w) {
		return
	}

	sendResp(task, w)
}
func (api *API) DeleteTasks(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	defer cancel()

	id := mux.Vars(req)["id"]
	isDelete, err := api.Conn.DeleteTask(ctx, id)
	if ServerInternalError(err, w) || RequestTimeoutError(err, w) {
		return
	}

	if !isDelete {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}
func (api *API) CompleteTask(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	defer cancel()

	id := mux.Vars(req)["id"]

	body, err := io.ReadAll(req.Body)
	if ServerInternalError(err, w) {
		return
	}
	var info db.CompleteDTO
	err = json.Unmarshal(body, &info)
	if ServerInternalError(err, w) {
		return
	}

	var task db.TasksResp

	task, err = api.Conn.CompleteTask(ctx, info, id)
	if ServerInternalError(err, w) || RequestTimeoutError(err, w) {
		return
	}

	sendResp(task, w)
}

func sendResp(obj any, w http.ResponseWriter) {
	resp, err := json.Marshal(obj)
	if ServerInternalError(err, w) {
		return
	}
	if _, err = w.Write(resp); err != nil {
		ServerInternalError(err, w)
		return
	}
}
