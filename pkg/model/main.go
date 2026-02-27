package model

type Todo struct {
	ID       int     `json:"id"`
	Task     string  `json:"task"`
	Group    *string `json:"group,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Done     *bool   `json:"done,omitempty"`
}
