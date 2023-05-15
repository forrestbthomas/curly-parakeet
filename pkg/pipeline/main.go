package pipeline

import (
	"fmt"

	"github.com/forrestbthomas/curly-parakeet/pkg/task"
)

type Job struct {
	Fn    task.TaskWork
	Task  task.Tasker
	Needs []task.Tasker
}

type Pipeliner interface {
	Run([]task.TaskDefinition) chan int
	HasNext() bool
	Next() chan int
}

type Pipe struct {
	MutableTaskDefinitions   []task.TaskDefinition
	ImmutableTaskDefinitions []task.TaskDefinition // how to enforce?
}

func (p Pipe) HasNext() bool {
	return len(p.MutableTaskDefinitions) > 0
}

func (p *Pipe) Next() task.TaskDefinition {
	task := p.MutableTaskDefinitions[0]
	defs := p.MutableTaskDefinitions[1:]
	p.MutableTaskDefinitions = defs
	return task
}

func (p Pipe) Run(ch chan int) chan int {
	fmt.Println("run")
	if !p.HasNext() {
		return ch
	}
	task := p.Next()
	return p.Run(task(ch))
}

func New(jobs []Job) Pipe {
	defs := []task.TaskDefinition{}
	copyDefs := []task.TaskDefinition{}
	copy(defs, copyDefs)
	for _, job := range jobs {
		defs = append(defs, job.Task.Generator(job.Fn))
	}
	return Pipe{
		defs,
		copyDefs,
	}
}
