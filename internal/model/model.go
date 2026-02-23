package model

import (
	"database/sql"
)

type Todo struct {
	ID       int
	Status   string
	Priority string
	Category string
	Task     string
	Lastmod  sql.NullTime
}

type Meta struct {
	Name  string
	Emoji string
}
