/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/masputrawae/go-todo/model"
	"github.com/masputrawae/go-todo/todo"
	"github.com/spf13/cobra"
)

var (
	lsByID       int
	lsByPriority string
	lsByGroup    string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var results []model.Todo
		all := Todos.FindAll()
		filtered := false
		if lsByID != -1 {
			index, err := Todos.FindIndexByID(lsByID)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			results = append(results, Todos.FindByIndex(index))
			filtered = true
		}
		if lsByGroup != "" {
			for i := range all {
				if all[i].Group == &lsByGroup {
					results = append(results, all[i])
				}
			}
			filtered = true
		}
		if lsByPriority != "" {
			for i := range all {
				if all[i].Priority == &lsByPriority {
					results = append(results, all[i])
				}
			}
			filtered = true
		}

		if !filtered {
			results = all
		}

		for i := range results {
			fmt.Printf("[%d] ", results[i].ID)
			if results[i].Priority != nil {
				p := *results[i].Priority
				r := rune(p[0])
				fmt.Printf("(%s) ", todo.Priorities[r].Emoji)
			}
			fmt.Printf("%s ", results[i].Task)
			if results[i].Group != nil {
				fmt.Printf("(%s)", *results[i].Group)
			}
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&lsByID, "id", "i", -1, "id")
	listCmd.Flags().StringVarP(&lsByGroup, "group", "g", "", "group")
	listCmd.Flags().StringVarP(&lsByPriority, "priority", "p", "", "priority")
}
