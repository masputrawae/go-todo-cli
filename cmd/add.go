/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/masputrawae/go-todo/model"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var nTodo model.Todo
		nTodo.ID = Todos.GenID()
		nTodo.Task = args[0]

		if priority != "" {
			pr := strings.ToUpper(priority)
			if !Todos.IsValidPriority(pr) {
				log.Fatal("prioritas hanya boleh 1 karakter: A - E atau a - e")
			}
			nTodo.Priority = &pr
		}
		if group != "" {
			nTodo.Group = &group
		}
		Todos.Add(nTodo)
		Todos.Save()
	},
}

func init() {
	addCmd.Flags().StringVarP(&group, "group", "g", "", "group")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "", "priority")
	rootCmd.AddCommand(addCmd)
}
