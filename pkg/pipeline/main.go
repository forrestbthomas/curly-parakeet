package pipeline

import (
	"github.com/forrestbthomas/curly-parakeet/pkg/task"
	"github.com/forrestbthomas/curly-parakeet/pkg/types"
)

type Job struct {
	fn    types.TaskWork
	task  task.Tasker
	needs []task.Tasker
}

type Pipeliner interface {
	Generator([]types.TaskDefinition, ...task.Tasker) []types.TaskDefinition
}

type Pipe struct{}

func New(jobs []Job) []types.TaskDefinition {
	defs := []types.TaskDefinition{}
	for _, job := range jobs {
		defs = append(defs, job.task.Generator(job.fn))
	}
	return defs
}
