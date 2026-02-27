/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"sort"

	"github.com/masputrawae/go-todo/pkg/model"
	"github.com/masputrawae/go-todo/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "A brief description of your command",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sort.Slice(Todos, func(i, j int) bool { return Todos[i].ID < Todos[j].ID })
		id := 1
		for i := range Todos {
			if Todos[i].ID == id {
				id++
			}
		}
		todo := model.Todo{ID: id, Task: args[0]}
		if priority != "" && utils.IsLetter(priority) {
			todo.Priority = &priority
		}
		if group != "" {
			todo.Group = &group
		}

		Todos = append(Todos, todo)
		utils.SaveTodo("todos.json", Todos)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&group, "group", "g", "", "group")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "", "priority")
}
