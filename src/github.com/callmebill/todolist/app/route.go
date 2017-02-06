package app

import (
	//"fmt"
	"net/http"
)

type Route struct {
	name    string
	method  string
	path    string
	handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GetTodos", "GET", "/todos", GetTodosHandler},
	Route{"CreatTodo", "POST", "/todos", CreatTodoHandler},
	Route{"UpdateTodo", "PUT", "/todos/{id}", UpdateTodoHandler},
	Route{"DeleteTodo", "DELETE", "/todos/{id}", DeleteTodoHandler},
}
