/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/masputrawae/todo-cli/internal/model"
	"github.com/spf13/cobra"
)

var (
	addFlagPriority string
	addFlagStatus   string
	addFlagProject  string
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "A brief description of your command",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var todo model.TodoAddInput
		todo.Task = args[0]

		if addFlagPriority != "" {
			todo.Priority = &addFlagPriority
		}
		if addFlagStatus != "" {
			todo.Priority = &addFlagStatus
		}
		if addFlagProject != "" {
			todo.Priority = &addFlagProject
		}

		if err := Svc.Create(todo); err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addFlagPriority, "priority", "p", "", "priority")
	addCmd.Flags().StringVarP(&addFlagStatus, "status", "s", "", "status")
	addCmd.Flags().StringVarP(&addFlagProject, "project", "g", "", "project")
}
