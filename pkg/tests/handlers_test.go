package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-api/pkg/db"
)

func TestAPI_CreateTasks(t *testing.T) {
	data := []byte(`{
    "title": "Task#1",
    "description": "test",
    "dueDate": "2024-11-20"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(data))
	rr := httptest.NewRecorder()
	api.CreateTasks(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Получен код: %d, ожидался: %d", rr.Code, http.StatusOK)
	}

	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	var info db.Task
	json.Unmarshal(body, &info)
	t.Logf("Request: %v\n", string(data))
	t.Logf("Response: %v\n", info)
}
func TestAPI_GetTasks(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rr := httptest.NewRecorder()
	api.GetTasks(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Получен код: %d, ожидался: %d", rr.Code, http.StatusOK)
	}

	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	var info []db.TasksResp
	json.Unmarshal(body, &info)
	t.Logf("Response: %v\n", info)
}
func TestAPI_PutTasks(t *testing.T) {
	data := []byte(`{
    "title": "Updates_Task#1",
    "description": "update_test",
    "dueDate": "2024-11-25"}`)
	req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(data))
	rr := httptest.NewRecorder()
	api.PutTasks(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Получен код: %d, ожидался: %d", rr.Code, http.StatusOK)
	}

	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	var info db.TasksResp
	json.Unmarshal(body, &info)
	t.Logf("Request: %v\n", string(data))
	t.Logf("Response: %v\n", info)
}
func TestAPI_CompleteTask(t *testing.T) {
	data := []byte(`{"completed": true}`)
	req := httptest.NewRequest(http.MethodPatch, "/tasks/1/complete", bytes.NewBuffer(data))
	rr := httptest.NewRecorder()
	api.CompleteTask(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Получен код: %d, ожидался: %d", rr.Code, http.StatusOK)
	}

	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	var info db.TasksResp
	json.Unmarshal(body, &info)
	t.Logf("Request: %v\n", string(data))
	t.Logf("Response: %v\n", info)
}
func TestAPI_DeleteTasks(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Test#1",
			want: 200,
		},
		{
			name: "Test#2",
			want: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
			rr := httptest.NewRecorder()

			api.DeleteTasks(rr, req)

			if rr.Code != tt.want {
				t.Errorf("Получен код: %v, ожидался: %v", rr.Code, tt.want)
			}
		})
	}
}
