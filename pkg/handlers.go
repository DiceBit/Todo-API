package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
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

	task, err := api.conn.AddTask(context.Background(), info)
	if ServerInternalError(err, w) {
		return
	}

	sendResp(task, w)
}
func (api *API) GetTasks(w http.ResponseWriter, req *http.Request) {
	tasks, err := api.conn.Tasks(context.Background())
	if ServerInternalError(err, w) {
		return
	}

	sendResp(tasks, w)
}
func (api *API) PutTasks(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	body, err := io.ReadAll(req.Body)
	ServerInternalError(err, w)

	var info db.TasksDTO
	err = json.Unmarshal(body, &info)
	ServerInternalError(err, w)

	if info.Title == "" {
		BadRequestError(errors.New("please, input title for your task"), w)
		return
	}

	task, err := api.conn.UpdateTask(context.Background(), info, id)
	if ServerInternalError(err, w) {
		return
	}

	sendResp(task, w)
}
func (api *API) DeleteTasks(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	isDelete, err := api.conn.DeleteTask(context.TODO(), id)
	if ServerInternalError(err, w) {
		return
	}

	if isDelete {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (api *API) CompleteTask(w http.ResponseWriter, req *http.Request) {
	/*
		TODO do
		Отметка завершения задачи

		Параметры запроса:
			JSON объект с полем completed (boolean).

		Ответ:
			обновленный объект задачи.
	*/
	id := mux.Vars(req)["id"]

	body, err := io.ReadAll(req.Body)
	ServerInternalError(err, w)
	var info db.CompleteDTO
	err = json.Unmarshal(body, &info)
	ServerInternalError(err, w)

	task, err := api.conn.CompleteTask(context.TODO(), info, id)
	if ServerInternalError(err, w) {
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

//TODO: !!! добавить поле completed в БД и в запросы, чекнуть предыдущие методы
