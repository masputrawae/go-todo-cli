package data

import "github.com/masputrawae/todo-cli/pkg/model"

var (
	Default = struct {
		Status   string
		Priority string
	}{
		Status:   "planning",
		Priority: "medium",
	}

	Statuses = []model.Status{
		{ID: "planning", Name: "Planning", Emoji: "📋", Color: "\033[90m"},       // Abu-abu
		{ID: "active", Name: "Active", Emoji: "✨", Color: "\033[92m"},           // Hijau terang
		{ID: "in-progress", Name: "In Progress", Emoji: "🔄", Color: "\033[94m"}, // Biru terang
		{ID: "done", Name: "Done", Emoji: "✅", Color: "\033[32m"},               // Hijau
		{ID: "cancelled", Name: "Cancelled", Emoji: "❌", Color: "\033[91m"},     // Merah terang
		{ID: "archive", Name: "Archive", Emoji: "📦", Color: "\033[90m"},         // Abu-abu
		{ID: "trash", Name: "Trash", Emoji: "🗑️", Color: "\033[90m"},            // Abu-abu
	}

	Priorities = []model.Priority{
		{ID: "highest", Name: "Highest", Emoji: "🔴", Color: "\033[91m", Order: 1}, // Merah terang
		{ID: "high", Name: "High", Emoji: "🟠", Color: "\033[91m", Order: 2},       // Merah terang (atau oranye, tapi ANSI tidak punya oranye)
		{ID: "medium", Name: "Medium", Emoji: "🟡", Color: "\033[93m", Order: 3},   // Kuning
		{ID: "low", Name: "Low", Emoji: "🟢", Color: "\033[92m", Order: 4},         // Hijau terang
		{ID: "lowest", Name: "Lowest", Emoji: "🔵", Color: "\033[94m", Order: 5},   // Biru terang
	}
)
