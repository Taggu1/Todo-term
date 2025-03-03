package database

import (
	"encoding/json"
	"log"
	"os"

	"example.org/todo-term/todo"
)

type Database struct {
	Todos     []todo.Todo `json:"todos"`
	DoneTodos []todo.Todo `json:"done_todos"`
}

func SaveTodos(todos []todo.Todo, doneTodos []todo.Todo) {

	j, _ := json.Marshal(Database{Todos: todos, DoneTodos: doneTodos})

	err := os.WriteFile("database.json", j, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadTodos() Database {
	var db Database
	b, err := os.ReadFile("database.json")
	if err != nil {
		return Database{
			Todos:     todo.InitialTodos,
			DoneTodos: []todo.Todo{},
		}
	}

	err = json.Unmarshal(b, &db)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
