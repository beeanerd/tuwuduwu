package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/beeanerd/tuwuduwu/models"
)

const (
	database string = "./tasks.db"
	create   string = `
  CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER NOT NULL PRIMARY KEY,
  due DATETIME NOT NULL,
  title TEXT,
  type TEXT,
  desc TEXT
  );`
)

/* Accesses the database and bootstraps the table if one doesn't already exist
 */
func ConnectDatabase() (*models.Todo, error) {
	db, err := sql.Open("go-sqlite3", database)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return &models.Todo{
		Db: db,
	}, nil
}

func AddTodoEntry(todoItem *models.TodoItem, Todo *models.Todo) (int, error) {
	res, err := Todo.Db.Exec(
		"INSERT INTO tasks VALUES(NULL,?,?,?,?);",
		todoItem.Due,
		todoItem.Name,
		todoItem.Type,
		todoItem.Desc,
	)
	if err != nil {
		return 0, nil
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}
