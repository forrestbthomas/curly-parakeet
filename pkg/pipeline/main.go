package pipeline

import (
	"github.com/forrestbthomas/curly-parakeet/pkg/task"
	"github.com/forrestbthomas/curly-parakeet/pkg/types"
)

type Pipeliner interface {
	Generator([]types.TaskDefinition, ...task.Tasker) []types.TaskDefinition
}

type Pipe struct{}

type TaskMap map[task.Tasker][]types.TaskWork

func (p *Pipe) Generator(taskMap TaskMap) []types.TaskDefinition {
	tasks := []types.TaskDefinition{}
	for task, fns := range taskMap {
		for _, fn := range fns {
			tasks = append(tasks, task.Generator(fn))
		}
	}
	return tasks
}

func New() *Pipe {
	return &Pipe{}
}
