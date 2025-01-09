package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"adventofcode.com/internal/solver"
)

const allTasksArg = "all"

func main() {
	var rootCmd = cobra.Command{
		Use:     "solver",
		Long:    "Advent Of Code puzzles solver",
		Version: "1",
	}

	for _, yearSolver := range []*solver.YearSolver{solver.Y2023, solver.Y2024} {
		yearExecutor := yearSolver
		tasks := append([]string{allTasksArg}, yearExecutor.KnownTasks()...)
		var yearCmd = cobra.Command{
			Use:       fmt.Sprintf("%s {%s}...", yearExecutor.Name, strings.Join(tasks, ", ")),
			Long:      fmt.Sprintf("Puzzles of Advent of Code %s ", yearExecutor.Name),
			Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.MinimumNArgs(1)),
			ValidArgs: tasks,
			RunE: func(cmd *cobra.Command, args []string) error {
				if slices.Contains(args, allTasksArg) {
					args = tasks[1:]
				}
				return yearExecutor.Solve(args)
			},
		}

		rootCmd.AddCommand(&yearCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
