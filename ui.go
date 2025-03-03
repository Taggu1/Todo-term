package main

import (
	"fmt"

	"example.org/todo-term/database"
	"example.org/todo-term/todo"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	todos      []todo.Todo
	doneTodos  []todo.Todo
	cursor     int
	todoText   textinput.Model
	addingTodo bool
}

func initialModel() model {
	db := database.LoadTodos()
	return model{
		todos:      db.Todos,
		doneTodos:  db.DoneTodos,
		addingTodo: false,
		todoText:   AddTodoWidget(),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			handleUp(&m)
		case "down", "s":
			handleDown(&m)
		case "a":
			m.addingTodo = true
		case "enter":
			handleEnter(&m)
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	if m.addingTodo {
		m.todoText, cmd = m.todoText.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	s := "Welcome to term-todo\n\n"

	// var todoTextStyle = lipgloss.NewStyle().
	// 	Foreground(lipgloss.Color("#cdd6f4")).
	// 	Inline(true)

	for index, todo := range m.todos {

		current := " "
		if index == m.cursor {
			current = "x"
		}

		s += fmt.Sprintf("[%s] %s", current, todo.Title)
		s += "\n"
	}

	var doneTextStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#bac2de")).
		Inline(true)

	s += "\n"
	s += doneTextStyle.Render("-----Done----")
	s += "\n\n"

	var doneTodoTextStyle = lipgloss.NewStyle().
		Strikethrough(true).
		AlignHorizontal(lipgloss.Left).
		Foreground(lipgloss.Color("#45475a")).
		Inline(true)

	for _, todo := range m.doneTodos {

		s += doneTodoTextStyle.Render(fmt.Sprintf("[✅] %s At %v", todo.Title, todo.DoneAt))
		s += "\n"
	}

	s += "\n"

	if m.addingTodo {
		s += m.todoText.View()
	}

	return s
}
