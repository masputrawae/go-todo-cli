package store

type Data struct {
	Name  string
	Emoji string
	Color string
}

var Priorities = map[string]Data{
	"A": {Name: "Highest", Emoji: "🔴", Color: "\033[31m"},
	"B": {Name: "High", Emoji: "🟠", Color: "\033[33m"},
	"C": {Name: "Medium", Emoji: "🟢", Color: "\033[32m"},
	"D": {Name: "Low", Emoji: "🔵", Color: "\033[34m"},
	"E": {Name: "Lowest", Emoji: "🟣", Color: "\033[35m"},
}
