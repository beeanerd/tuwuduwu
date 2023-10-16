package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/beeanerd/tuwuduwu/models"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {

  reqBody, err := io.ReadAll(r.Body)

  if err != nil {
    log.Fatal(err)
  }

  var todoItem models.ToDoItem
  json.Unmarshal(reqBody, &todoItem)


  /* this populated todoItem with the Request's body, now I will need to add this to some
  more permanent storage */


}
