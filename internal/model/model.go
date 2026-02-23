package model

import "time"

type Todo struct {
	ID       int
	Status   string
	Priority string
	Category string
	Task     string
	Lastmod  time.Time
}

type Meta struct {
	Name  string
	Emoji string
}
