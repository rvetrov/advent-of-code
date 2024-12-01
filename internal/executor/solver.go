package executor

type taskSolver func(string) int

type Task struct {
	solvers []taskSolver
}
