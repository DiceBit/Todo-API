package tests

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
	"testing"
	"todo-api/pkg/db"
)

func TestSqlite_AddTask(t *testing.T) {
	mock.Db.Exec(`delete from Tasks`)
	tests := []struct {
		name string
		dto  db.TasksDTO
		want db.Task
	}{
		{
			name: "Test1",
			dto: db.TasksDTO{
				Title:       "TestTask#1",
				Description: "TestDescription#1",
				DueDate:     "2024-01-12",
			},
			want: db.Task{
				Id:          1,
				Title:       "TestTask#1",
				Description: "TestDescription#1",
				DueDate:     "2024-01-12",
				Overdue:     false,
				Completed:   false,
			},
		},
		{
			name: "Test2",
			dto: db.TasksDTO{
				Title:       "",
				Description: "",
				DueDate:     "",
			},
			want: db.Task{
				Id:          2,
				Title:       "",
				Description: "",
				DueDate:     "",
				Overdue:     false,
				Completed:   false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := mock.AddTask(context.Background(), tt.dto)
			if !reflect.DeepEqual(&got, &tt.want) {
				t.Errorf("mock.AddTask() = %v, want %v", got, tt.want)
			}
		})
	}

}
func TestSqlite_Tasks(t *testing.T) {
	mock.Db.Exec(`delete from Tasks`)

	tests := []struct {
		name string
		want []db.TasksResp
	}{
		{
			name: "Test1",
			want: nil,
		},
		{
			name: "Test2",
			want: []db.TasksResp{
				{
					Title:       "Title1",
					Description: "desc test",
					DueDate:     "2024-01-12",
					Overdue:     true,
					Completed:   true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := mock.Tasks(context.Background())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mock.Tasks() = %v, want %v", got, tt.want)
			}
		})
		_, _ = mock.Db.Exec(`insert into Tasks(Title, Description, DueDate, Overdue, Completed)
							VALUES ('Title1', 'desc test', '2024-01-12', true, true)`)
	}
}
func TestSqlite_UpdateTask(t *testing.T) {
	mock.Db.Exec(`delete from Tasks`)
	mock.Db.Exec(`insert into Tasks(id, Title, Description, DueDate, Overdue, Completed)
							VALUES (1, 'Title1', 'desc test', '2024-01-12', false, false)`)

	tests := []struct {
		name string
		dto  db.TasksDTO
		id   string
		want db.TasksResp
	}{
		{
			name: "Test1",
			dto: db.TasksDTO{
				Title:       "ChangedTitle1",
				Description: "changeDesc",
				DueDate:     "2011-12-12",
			},
			id: "1",
			want: db.TasksResp{
				Title:       "ChangedTitle1",
				Description: "changeDesc",
				DueDate:     "2011-12-12",
				Overdue:     false,
				Completed:   false,
			},
		},
		{
			name: "Test2",
			dto: db.TasksDTO{
				Title:       "1",
				Description: "2",
				DueDate:     "3",
			},
			id:   "123",
			want: db.TasksResp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := mock.UpdateTask(context.Background(), tt.dto, tt.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mock.UpdateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSqlite_DeleteTask(t *testing.T) {
	mock.Db.Exec(`delete from Tasks`)
	mock.Db.Exec(`insert into Tasks(id, Title, Description, DueDate, Overdue, Completed)
							VALUES (1, 'Title1', 'desc test', '2024-01-12', false, false)`)

	tests := []struct {
		name string
		id   string
		want bool
	}{
		{
			name: "Test1",
			id:   "1",
			want: true,
		},
		{
			name: "Test2",
			id:   "2",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mock.DeleteTask(context.Background(), tt.id)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mock.DeleteTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSqlite_CompleteTask(t *testing.T) {
	mock.Db.Exec(`delete from Tasks`)
	mock.Db.Exec(`insert into Tasks(id, Title, Description, DueDate, Overdue, Completed)
							VALUES (1, 'Title1', 'desc test', '2024-01-12', false, false)`)

	tests := []struct {
		name string
		dto  db.CompleteDTO
		id   string
		want db.TasksResp
	}{
		{
			name: "Test1",
			dto: db.CompleteDTO{
				Completed: true,
			},
			id: "1",
			want: db.TasksResp{
				Title:       "Title1",
				Description: "desc test",
				DueDate:     "2024-01-12",
				Overdue:     false,
				Completed:   true,
			},
		},
		{
			name: "Test2",
			dto: db.CompleteDTO{
				Completed: true,
			},
			id:   "123",
			want: db.TasksResp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := mock.CompleteTask(context.Background(), tt.dto, tt.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mock.CompleteTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
