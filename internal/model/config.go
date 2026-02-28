package model

type Config struct {
	Statuses   []Status   `yaml:"statuses"`
	Priorities []Priority `yaml:"priorities"`
}
