package main

import (
	"time"

	"example.org/todo-term/todo"
)

func handleEnter(m *model) {
	if m.addingTodo {
		m.addingTodo = false
		newTodo := todo.Todo{
			Title:     m.todoText.Value(),
			Done:      false,
			CreatedAt: time.Now(),
		}
		m.todos = append(m.todos, newTodo)
		m.todoText.Reset()
	} else {
		for index, _ := range m.todos {
			if index == m.cursor {
				m.todos[index].Done = !m.todos[index].Done
			}
		}
	}
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
