package tests

import (
	"database/sql"
	"log"
	"testing"
	"todo-api/pkg"
	"todo-api/pkg/db/mockDb"
	"todo-api/pkg/db/sqlite"
)

var mock *sqlite.Sqlite
var api *pkg.API

func TestMain(m *testing.M) {
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
