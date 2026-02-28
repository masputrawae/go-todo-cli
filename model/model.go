package model

type Todo struct {
	ID       int     `json:"id"`
	Task     string  `json:"task"`
	Group    *string `json:"group,omitempty"`
	Done     *bool   `json:"done,omitempty"`
	Priority *string `json:"priority,omitempty"`
}

type Priority struct {
	Name  string
	Emoji string
}
