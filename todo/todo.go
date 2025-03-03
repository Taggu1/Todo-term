package todo

type Todo struct {
	DoneAt string `json:"done_at"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}

var InitialTodos = []Todo{
	{
		Title: "Press A to add Todo",
	},
	{
		Title: "Press ENTER to check todo",
	},
}
