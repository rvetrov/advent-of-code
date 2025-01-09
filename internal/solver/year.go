package solver

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"time"

	"golang.org/x/exp/maps"

	"adventofcode.com/internal/utils"
)

var errNotFound = fmt.Errorf("task not found")

type YearSolver struct {
	Name         string
	tasksDirName string
	solvers      map[string]DaySolver
}

func (slv *YearSolver) KnownTasks() []string {
	tasks := maps.Keys(slv.solvers)
	sort.Strings(tasks)
	return tasks
}

func (slv *YearSolver) Solve(taskNames []string) error {
	var errs []error

	for _, taskName := range taskNames {
		if task, found := slv.solvers[taskName]; !found {
			errs = append(errs, errNotFound)
		} else if err := slv.solveTask(taskName, task); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (slv *YearSolver) solveTask(taskName string, t DaySolver) error {
	inputPath := path.Join(slv.tasksDirName, taskName, "input.big")
	if _, err := os.Stat(inputPath); err != nil {
		return err
	}

	if input, err := utils.ReadInput(inputPath); err != nil {
		return err
	} else {
		for i, solver := range t.solvers {
			if solver == nil {
				continue
			}
			resultPath := path.Join(slv.tasksDirName, taskName, fmt.Sprintf("output.v%d", i+1))

			startedAt := time.Now()

			var resultStr string
			switch solver.(type) {
			case func(string) int:
				result := solver.(func(string) int)(input)
				resultStr = fmt.Sprint(result)
			case func(string) int64:
				result := solver.(func(string) int64)(input)
				resultStr = fmt.Sprint(result)
			case func(string) string:
				resultStr = solver.(func(string) string)(input)
			default:
				panic(fmt.Sprintf("Unknown solver type: %T", solver))
			}

			ms := int(time.Since(startedAt).Milliseconds())
			fmt.Printf("%s: %v -> %v, %d.%03ds\n", taskName, inputPath, resultPath, ms/1000, ms%1000)

			if err = os.WriteFile(resultPath, []byte(resultStr), 0644); err != nil {
				return err
			}
		}
	}
	return nil
}
