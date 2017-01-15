package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	s "strings"

	"github.com/gorilla/mux"
)

func GetTodosHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	todos := FindAllTodos()
	err := json.NewEncoder(resp).Encode(todos)
	checkErr(err)
}

func CreatTodoHandler(resp http.ResponseWriter, req *http.Request) {
	todo := decodeTodoFromReq(resp, req)
	t := CreateTodo(todo)
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(resp).Encode(t); err != nil {
		panic(err)
	}

}

func decodeTodoFromReq(resp http.ResponseWriter, req *http.Request) Todo {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(resp).Encode(err); err != nil {
			panic(err)
		}
	}

	return todo
}

func UpdateTodoHandler(resp http.ResponseWriter, req *http.Request) {
	todo := decodeTodoFromReq(resp, req)
	UpdateTodo(todo)
}

func DeleteTodoHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todoId := vars["id"]
	if s.Contains(todoId, ",") {
		// We pass the data to backend directly, which is a securit hole!!!
		// Do not do this in production
		DestroyTodos(todoId)

	} else {
		id, err := strconv.Atoi(todoId)
		if err != nil {
			panic(err)
		}

		DestroyTodoItem(id)
	}

}
