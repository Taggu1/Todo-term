package todo

import "time"

type Todo struct {
	CreatedAt time.Time
	Title     string
	Done      bool
}

var InitialTodos = []Todo{
	{
		Title: "Press A to add Todo",
	},
	{
		Title: "Press ENTER to check todo",
	},
}
