package main

import (
	"example.org/todo-term/database"
	"example.org/todo-term/todo"
	"example.org/todo-term/utility"
)

func handleEnter(m *model) {

	if m.addingTodo {

		m.addingTodo = false
		newTodo := todo.Todo{
			Title:  m.todoText.Value(),
			Done:   false,
			DoneAt: "",
		}
		m.todos = append(m.todos, newTodo)
		m.todoText.Reset()
	} else if m.todoText.Value() != " " {
		for index, _ := range m.todos {
			if index == m.cursor {
				m.todos[index].Done = !m.todos[index].Done
				m.todos[index].DoneAt = utility.FormatedNowTime()
				m.doneTodos = append(m.doneTodos, m.todos[index])
				m.todos = utility.RemoveSliceElement(m.todos, index)
			}
		}
	}
	database.SaveTodos(m.todos, m.doneTodos)
}

func handleUp(m *model) {
	if m.cursor > 0 && !m.addingTodo {
		m.cursor--
	}
}

func handleDown(m *model) {
	if m.cursor < len(m.todos)-1 && !m.addingTodo {
		m.cursor++
	}
}
