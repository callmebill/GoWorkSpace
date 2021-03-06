package app

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"strconv"
	s "strings"
)

var (
	CREATE_TABLE     = "CREATE TABLE IF NOT EXISTS todos(id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, completed BOOL)"
	INSERT_TODO      = "INSERT INTO todos (title, completed) VALUES (?, ?)"
	UPDATE_TODO      = "UPDATE todos SET title = ?, completed = ? where id = ?"
	DELETE_TODO_ITEM = "DELETE FROM todos where id = ?"
	DELETE_TODOs     = "DELETE FROM todos where id in ( ? )"
	SELECT_ALL_TODOS = "SELECT * FROM todos"
	SELECT_TODO      = "SELECT title,   FROM todos WHERE id = ?"
	MAX_ID           = "SELECT id FROM todos ORDER BY ID DESC LIMIT 1"
	db               *sql.DB
)

func init() {
	db = GetDB(db)
	_, err := db.Exec(CREATE_TABLE)
	checkErr(err)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "todos.sqlite3")
	checkErr(err)
	return db
}
func GetDB(db *sql.DB) *sql.DB {
	if db == nil {
		return NewDB()
	} else {
		return db
	}
}
func FindAllTodos() Todos {
	var result Todos
	rows, err := db.Query(SELECT_ALL_TODOS)
	checkErr(err)
	for rows.Next() {
		var (
			id        int
			title     string
			completed bool
			todo      Todo
		)
		err = rows.Scan(&id, &title, &completed)
		checkErr(err)
		todo = Todo{id, title, completed}
		result = append(result, todo)
	}
	return result
}
func CreateTodo(t Todo) Todo {
	// No transaction at all
	stmt, err := db.Prepare(INSERT_TODO)
	defer stmt.Close()
	checkErr(err)

	res, err := stmt.Exec(t.Title, t.Completed)
	checkErr(err)

	lastId, err := res.LastInsertId()
	checkErr(err)
	// Convert int64 to int, may lost precision
	t.Id = int(lastId)
	return t
}
func UpdateTodo(todo Todo) {
	stmt, err := db.Prepare(UPDATE_TODO)
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Completed, todo.Id)
	checkErr(err)
}

func DestroyTodoItem(id int) {
	_, err := db.Exec(DELETE_TODO_ITEM, id)
	checkErr(err)
}

func DestroyTodos(ids string) {
	// I don't know why this not work
	//_, err := db.Exec(DELETE_TODOs, ids)

	// Since above code with in clause not work
	// Here I issue multiple delete query for more than one todo items
	idsArr := s.Split(ids, ",")
	for _, idStr := range idsArr {
		id, err := strconv.Atoi(idStr)
		checkErr(err)
		DestroyTodoItem(id)
	}
}
