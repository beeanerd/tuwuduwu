package models

type ToDoItem struct {
  Status Status     
  Name string
  Date string
  Type string // Might make this an enum as well
  Class string
}

type Status int64
const (
  Not_Started Status = iota
  Started
  Finished
)
