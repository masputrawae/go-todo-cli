package utils

import (
	"github.com/masputrawae/todo-cli/internal/model"
	"gopkg.in/yaml.v3"
	"log"
)

func LoadConfig(fp string) model.Config {
	// default config
	config := model.Config{
		Statuses: []model.Status{
			{ID: "active", Name: "Active", Emoji: "🔵"},
			{ID: "in-progress", Name: "In Progress", Emoji: "🟢"},
			{ID: "done", Name: "Done", Emoji: "⚫"},
			{ID: "cancelled", Name: "Active", Emoji: "🔴"},
			{ID: "archive", Name: "Archive", Emoji: "🟣"},
		},
		Priorities: []model.Priority{
			{ID: "highest", Short: "A", Name: "Highest", Color: ""},
			{ID: "high", Short: "B", Name: "High", Color: ""},
			{ID: "medium", Short: "C", Name: "Medium", Color: ""},
			{ID: "low", Short: "D", Name: "Low", Color: ""},
			{ID: "lowest", Short: "E", Name: "Lowest", Color: ""},
		},
	}

	file, err := ResolveFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		if err.Error() != "EOF" {
			log.Fatal(err)
		}
		if err := yaml.NewEncoder(file).Encode(&config); err != nil {
			log.Fatal(err)
		}
	}
	return config
}
