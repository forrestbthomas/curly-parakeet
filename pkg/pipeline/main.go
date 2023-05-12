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
	Generator([]types.TaskDefinition, ...task.Tasker) []types.TaskDefinition
}

type Pipe struct{}

func New(jobs []Job) []types.TaskDefinition {
	defs := []types.TaskDefinition{}
	for _, job := range jobs {
		defs = append(defs, job.Task.Generator(job.Fn))
	}
	return defs
}
