package main

import (
  "net/http"
	"github.com/beeanerd/tuwuduwu/routes"
)

func main() {
  r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
