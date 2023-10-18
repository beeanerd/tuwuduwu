package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/beeanerd/tuwuduwu/database"
	"github.com/beeanerd/tuwuduwu/models"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var todoItem models.TodoItem
	json.Unmarshal(reqBody, todoItem)

	TodoDB, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	database.AddTodoEntry(&todoItem, TodoDB)
}
