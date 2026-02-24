package model

type TodoAPIRequest struct {
	Todos      []Todo
	Statuses   []Status
	Priorities []Priority
}

// ==== JSON MODELS (todo) =====
// todos
type Todo struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	Category string `json:"category,omitempty"`
	Task     string `json:"task"`
	Lastmod  string `json:"lastmod"`
}

// ===== INPUT TYPE MODELS (todo) =====
// todo create input type
type TodoCreateInputType struct {
	Status   *string
	Priority *string
	Category *string
	Task     *string
}

// todo update input type
type TodoUpdateInputType struct {
	ID       *int
	Status   *string
	Priority *string
	Category *string
	Task     *string
}

// todo delete input type
type TodoDeleteInputType struct {
	ID *int
}

// todo find input type
type TodoFindInputType struct {
	ID       *int
	Status   *string
	Priority *string
	Category *string
}

// ==== MODELS (status) =====
// status
type Status struct {
	ID    string
	Name  string
	Emoji string
	Color string
}

// ==== JSON MODELS (priority) =====
// priority
type Priority struct {
	ID    string
	Name  string
	Emoji string
	Color string
	Order int
}
