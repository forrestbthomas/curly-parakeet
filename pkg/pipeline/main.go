package pipeline

import (
	"github.com/forrestbthomas/curly-parakeet/pkg/task"
	"github.com/forrestbthomas/curly-parakeet/pkg/types"
)

type Job struct {
	Fn    types.TaskWork
	Task  task.Tasker
	Needs []task.Tasker
}

type Pipeliner interface {
	Run([]types.TaskDefinition) chan int
}

type Pipe struct {
	TaskDefinitions []types.TaskDefinition
}

func (p Pipe) Run(ch chan int) chan int {
	for _, task := range p.TaskDefinitions {
		ch = task(ch)
	}
	return ch
}

func New(jobs []Job) Pipe {
	defs := []types.TaskDefinition{}
	for _, job := range jobs {
		defs = append(defs, job.Task.Generator(job.Fn))
	}
	return Pipe{defs}
}
