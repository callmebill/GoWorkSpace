package app

type Todo struct {
	Id        int    `json:"id"` //结构体便签，用于json转换时声明Key
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todos []Todo
