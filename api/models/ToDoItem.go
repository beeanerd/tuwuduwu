package models

import (
	"database/sql"
	"sync"
)

type TodoItem struct {
	Name string `json:"name"`
	Due  string `json:"due"`
	Type string `json:"type"`
	Desc string `json:"desc"`
}

type Todo struct {
	Mu sync.Mutex `json:"mu"`
	Db *sql.DB    `json:"db"`
}
