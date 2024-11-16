package mockDb

import (
	"context"
	"strconv"
	"todo-api/pkg/db"
)

type MockDb struct {
	db []db.TasksResp
}

func (m *MockDb) AddTask(ctx context.Context, dto db.TasksDTO) (db.Task, error) {
	m.db = append(m.db, db.TasksResp{
		Title:       dto.Title,
		Description: dto.Description,
		DueDate:     dto.DueDate,
		Overdue:     false,
		Completed:   false,
	})
	task := db.Task{
		Id:          1,
		Title:       dto.Title,
		Description: dto.Description,
		DueDate:     dto.DueDate,
		Overdue:     false,
		Completed:   false,
	}
	return task, nil
}
func (m *MockDb) Tasks(ctx context.Context) ([]db.TasksResp, error) {
	return m.db, nil
}
func (m *MockDb) UpdateTask(ctx context.Context, dto db.TasksDTO, id string) (db.TasksResp, error) {
	_id, _ := strconv.Atoi(id)
	m.db[_id].Title = dto.Title
	m.db[_id].Description = dto.Description
	m.db[_id].DueDate = dto.DueDate
	return m.db[_id], nil
}
func (m *MockDb) DeleteTask(ctx context.Context, id string) (bool, error) {
	_id, _ := strconv.Atoi(id)
	if len(m.db) <= _id {
		return false, nil
	}
	m.db = append(m.db[:_id], m.db[_id+1:]...)
	return true, nil
}
func (m *MockDb) CompleteTask(ctx context.Context, dto db.CompleteDTO, id string) (db.TasksResp, error) {
	_id, _ := strconv.Atoi(id)
	m.db[_id].Completed = dto.Completed
	return m.db[_id], nil
}

func (m *MockDb) CheckTasks(ctx context.Context) error { return nil }
func (m *MockDb) CloseConn() {
}
