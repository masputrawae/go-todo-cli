package data

import (
	"github.com/masputrawae/todo-cli/internal/model"
)

var Statuses = []model.Meta{
	{Name: "Planning", Emoji: "📝"},
	{Name: "Active", Emoji: "🔥"},
	{Name: "In Progress", Emoji: "🚀"},
	{Name: "Done", Emoji: "✅"},
	{Name: "Archive", Emoji: "📦️"},
	{Name: "Cancelled", Emoji: "❌"},
}

var Priorities = []model.Meta{
	{Name: "Highest", Emoji: "🔴"},
	{Name: "High", Emoji: "🟠"},
	{Name: "Medium", Emoji: "🟢"},
	{Name: "Low", Emoji: "🔵"},
	{Name: "Lowest", Emoji: "🟣"},
}
