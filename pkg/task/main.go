package task

type Tasker interface {
	Generator(fn TaskWork) TaskDefinition
	Get(string) int
	Set(string, int)
}

type TaskState int
type TaskWork func(int, chan int, Tasker)
type TaskDefinition func(chan int) chan int
