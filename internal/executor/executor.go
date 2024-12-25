package executor

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

type Executor struct {
	Name    string
	path    string
	solvers map[string]Task
}

func (e *Executor) KnownTasks() []string {
	tasks := maps.Keys(e.solvers)
	sort.Strings(tasks)
	return tasks
}

func (e *Executor) Solve(taskNames []string) error {
	var errs []error

	for _, taskName := range taskNames {
		if task, found := e.solvers[taskName]; !found {
			errs = append(errs, errNotFound)
		} else if err := e.solveTask(taskName, task); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (e *Executor) solveTask(taskName string, t Task) error {
	inputPath := path.Join(e.path, taskName, "input.big")
	if _, err := os.Stat(inputPath); err != nil {
		return err
	}

	if input, err := utils.ReadInput(inputPath); err != nil {
		return err
	} else {
		for i, solver := range t.solvers {
			resultPath := path.Join(e.path, taskName, fmt.Sprintf("output.v%d", i+1))

			started := time.Now()

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
			ms := int(time.Since(started).Milliseconds())
			fmt.Printf("%s: %v -> %v, %d.%03ds\n", taskName, inputPath, resultPath, ms/1000, ms%1000)

			if err = os.WriteFile(resultPath, []byte(resultStr), 0644); err != nil {
				return err
			}
		}
	}
	return nil
}
