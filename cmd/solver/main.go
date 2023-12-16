package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"adventofcode.com/internal/executor"
	"github.com/spf13/cobra"
)

const allTasksArg = "all"

func main() {
	var path string

	var rootCmd = cobra.Command{
		Use:     "solver",
		Long:    "Advent Of Code puzzles solver",
		Version: "1",
	}

	tasks := []string{allTasksArg}
	sort.Strings(executor.KnownTasks2023)
	tasks = append(tasks, executor.KnownTasks2023...)
	var year2023 = cobra.Command{
		Use:       fmt.Sprintf("2023 {%s}...", strings.Join(tasks, ", ")),
		Long:      "Puzzles of 2023th Advent of Code",
		Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.MinimumNArgs(1)),
		ValidArgs: tasks,
		RunE: func(cmd *cobra.Command, args []string) error {
			if slices.Contains(args, allTasksArg) {
				args = tasks[1:]
			}
			e := executor.New2023(path)
			return e.Solve(args)
		},
	}

	rootCmd.AddCommand(&year2023)
	year2023.Flags().StringVarP(&path, "path", "p", "./y2023", "path to directory of puzzles")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
