package sqlite

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"todo-api/pkg/db"
)

type Sqlite struct {
	db *sql.DB
}

func NewConn() *Sqlite {
	path := filepath.Join(os.Getenv("GOPATH"), "todo-api", "pkg", "db", "todo.db")
	_db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = _db.Ping()
	if err != nil {
		log.Println(err)
		_db.Close()
		return nil
	}

	dbConn := Sqlite{
		db: _db,
	}
	return &dbConn
}

func (sl *Sqlite) CloseConn() {
	sl.db.Close()
}

func (sl *Sqlite) AddTask(ctx context.Context, dto db.TasksDTO) (db.Task, error) {
	tx, err := sl.db.Begin()
	if err != nil {
		return db.Task{}, err
	}
	defer tx.Rollback()

	query := `insert into Tasks(Title, Description, DueDate) values (?, ?, ?)`
	res, err := tx.Exec(query, dto.Title, dto.Description, dto.DueDate)
	if err != nil {
		return db.Task{}, err
	}

	var task db.Task
	id, err := res.LastInsertId()
	if err != nil {
		return db.Task{}, err
	}

	query = `select * from Tasks where id=?`
	row := tx.QueryRow(query, id)
	if err = row.Err(); err != nil {
		return db.Task{}, nil
	}

	err = row.Scan(&task.Id, &task.Title, &task.Description, &task.DueDate, &task.Overdue)
	if err != nil {
		return db.Task{}, err
	}

	if err = tx.Commit(); err != nil {
		return db.Task{}, err
	}
	return task, nil
}
func (sl *Sqlite) Tasks(ctx context.Context) ([]db.TasksResp, error) {
	query := `select Title, Description, DueDate, Overdue from Tasks`
	rows, err := sl.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []db.TasksResp
	for rows.Next() {
		var t db.TasksResp
		err := rows.Scan(
			&t.Title,
			&t.Description,
			&t.DueDate,
			&t.Overdue)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
func (sl *Sqlite) UpdateTask(ctx context.Context, dto db.TasksDTO, id string) (db.TasksResp, error) {
	tx, err := sl.db.Begin()
	if err != nil {
		return db.TasksResp{}, err
	}
	defer tx.Rollback()

	query := `update Tasks set Title=?, Description=?, DueDate=? where Id=?`
	_, err = tx.Exec(query, dto.Title, dto.Description, dto.DueDate, id)
	if err != nil {
		return db.TasksResp{}, err
	}

	var task db.TasksResp

	query = `select Title, Description, DueDate, Overdue from Tasks where id=?`
	row := tx.QueryRow(query, id)

	err = row.Scan(&task.Title, &task.Description, &task.DueDate, &task.Overdue)
	if err != nil {
		return db.TasksResp{}, err
	}

	if err = tx.Commit(); err != nil {
		return db.TasksResp{}, err
	}
	return task, nil
}
func (sl *Sqlite) DeleteTask(ctx context.Context, id string) (bool, error) {
	query := `delete from Tasks where Id=?`

	res, err := sl.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAf, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAf == 0 {
		return false, nil
	}

	return true, nil
}

func (sl *Sqlite) CompleteTask(ctx context.Context, dto db.CompleteDTO, id string) (db.TasksResp, error) {

	tx, err := sl.db.Begin()
	if err != nil {
		return db.TasksResp{}, err
	}
	defer tx.Rollback()

	query := `update Tasks set Completed=? where Id=?`
	_, err = tx.Exec(query, dto.Completed, id)
	if err != nil {
		return db.TasksResp{}, err
	}

	var task db.TasksResp

	query = `select Title, Description, DueDate, Overdue, Completed from Tasks where id=?`
	row := tx.QueryRow(query, id)

	err = row.Scan(&task.Title, &task.Description, &task.DueDate, &task.Overdue, &task.Completed)
	if err != nil {
		return db.TasksResp{}, err
	}

	if err = tx.Commit(); err != nil {
		return db.TasksResp{}, err
	}

	return db.TasksResp{}, nil
}
