package executor

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"adventofcode.com/internal/utils"
)

var errNotFound = fmt.Errorf("task not found")

type executor struct {
	path    string
	solvers map[string]Task
}

func New2023(path string) *executor {
	return &executor{
		path:    path,
		solvers: tasks2023,
	}
}

func (e *executor) Solve(taskNames []string) error {
	errs := []error{}

	for _, taskName := range taskNames {
		if task, found := e.solvers[taskName]; !found {
			errs = append(errs, errNotFound)
		} else if err := e.solveTask(taskName, task); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (e *executor) solveTask(taskName string, t Task) error {
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
			result := solver(input)
			ms := int(time.Since(started).Milliseconds())
			fmt.Printf("%s: %v -> %v, %d.%03ds\n", taskName, inputPath, resultPath, ms/1000, ms%1000)

			os.WriteFile(resultPath, []byte(strconv.Itoa(result)), 0644)
		}
	}
	return nil
}
