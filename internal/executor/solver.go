package executor

type taskSolver any

type Task struct {
	solvers []taskSolver
}
