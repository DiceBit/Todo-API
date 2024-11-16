package sqlite

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"
	"todo-api/pkg/db"
)

type Sqlite struct {
	Db *sql.DB
}

func InitDb(db *sql.DB) {
	query, err := os.ReadFile(filepath.Join(os.Getenv("APP_ROOT"), "pkg", "db", "schemas.sql"))
	if err != nil {
		log.Fatal("Can't find schemas.sql file ", err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		log.Fatal("Error initialization DB. ", err)
		return
	}
}
func NewConn() *Sqlite {
	path := filepath.Join(os.Getenv("APP_ROOT"), "pkg", "db", "todo.db")
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

	InitDb(_db)

	dbConn := Sqlite{
		Db: _db,
	}

	return &dbConn
}
func (sl *Sqlite) CloseConn() {
	sl.Db.Close()
}

func (sl *Sqlite) AddTask(ctx context.Context, dto db.TasksDTO) (db.Task, error) {
	tx, err := sl.Db.Begin()
	if err != nil {
		return db.Task{}, err
	}
	defer tx.Rollback()

	query := `insert into Tasks(Title, Description, DueDate) values (?, ?, ?)`
	res, err := tx.ExecContext(ctx, query, dto.Title, dto.Description, dto.DueDate)
	if err != nil {
		return db.Task{}, err
	}

	var task db.Task
	id, err := res.LastInsertId()
	if err != nil {
		return db.Task{}, err
	}

	query = `select * from Tasks where id=?`
	row := tx.QueryRowContext(ctx, query, id)
	if err = row.Err(); err != nil {
		return db.Task{}, err
	}

	err = row.Scan(&task.Id, &task.Title, &task.Description, &task.DueDate, &task.Overdue, &task.Completed)
	if err != nil {
		return db.Task{}, err
	}

	if err = tx.Commit(); err != nil {
		return db.Task{}, err
	}
	return task, nil
}
func (sl *Sqlite) Tasks(ctx context.Context) ([]db.TasksResp, error) {
	query := `select Title, Description, DueDate, Overdue, Completed from Tasks`
	rows, err := sl.Db.QueryContext(ctx, query)
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
			&t.Overdue,
			&t.Completed)
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
	tx, err := sl.Db.Begin()
	if err != nil {
		return db.TasksResp{}, err
	}
	defer tx.Rollback()

	query := `update Tasks set Title=?, Description=?, DueDate=? where Id=?`
	if _, err = tx.ExecContext(ctx, query, dto.Title, dto.Description, dto.DueDate, id); err != nil {
		return db.TasksResp{}, err
	}

	var task db.TasksResp

	query = `select Title, Description, DueDate, Overdue, Completed from Tasks where id=?`
	row := tx.QueryRowContext(ctx, query, id)

	err = row.Scan(&task.Title, &task.Description, &task.DueDate, &task.Overdue, &task.Completed)
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

	res, err := sl.Db.ExecContext(ctx, query, id)
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
	tx, err := sl.Db.Begin()
	if err != nil {
		return db.TasksResp{}, err
	}
	defer tx.Rollback()

	query := `update Tasks set Completed=? where Id=?`
	if _, err = tx.ExecContext(ctx, query, dto.Completed, id); err != nil {
		return db.TasksResp{}, err
	}

	var task db.TasksResp

	query = `select Title, Description, DueDate, Overdue, Completed from Tasks where id=?`
	row := tx.QueryRowContext(ctx, query, id)

	err = row.Scan(&task.Title, &task.Description, &task.DueDate, &task.Overdue, &task.Completed)
	if err != nil {
		return db.TasksResp{}, err
	}

	if err = tx.Commit(); err != nil {
		return db.TasksResp{}, err
	}
	return task, nil
}
func (sl *Sqlite) CheckTasks(ctx context.Context) error {
	now := time.Now().Format(time.DateOnly)
	query := `update Tasks set Overdue=true where DueDate < ? and Overdue=false`
	if _, err := sl.Db.ExecContext(ctx, query, now); err != nil {
		return err
	}
	return nil
}
