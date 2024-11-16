package tests

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"
	"todo-api/pkg"
	"todo-api/pkg/db/mockDb"
	"todo-api/pkg/db/sqlite"
)

var mock *sqlite.Sqlite
var api *pkg.API

func TestMain(m *testing.M) {
	//todo delete
	if err := os.Setenv("APP_ROOT", filepath.Join("C:\\", "Users", "danii", "PROGRAMMING", "GolandProjects", "todo-api")); err != nil {
		log.Fatal(err)
	}
	log.Println(os.Getenv("APP_ROOT"))

	_db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer _db.Close()

	sqlite.InitDb(_db)

	mock = &sqlite.Sqlite{Db: _db}

	api = &pkg.API{
		Conn: &mockDb.MockDb{},
	}

	m.Run()
}
