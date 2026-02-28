package model

import "time"

type Todo struct {
	ID       int        `json:"id"`
	Task     string     `json:"task"`
	Status   *string    `json:"status,omitempty"`
	Priority *string    `json:"priority,omitempty"`
	Project  *string    `json:"project,omitempty"`
	Lastmod  *time.Time `json:"lastmod,omitempty"`
}

type TodoAddInput struct {
	Task     string
	Status   *string
	Priority *string
	Project  *string
}

type TodoEditInput struct {
	ID       int
	Task     *string
	Status   *string
	Priority *string
	Project  *string
}
