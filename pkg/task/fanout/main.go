package task

import "github.com/forrestbthomas/curly-parakeet/pkg/types"

type Task struct {
	Name string
}

func (t Task) GetName() string {
	return t.Name
}

func (t Task) Generator(fn types.TaskWork) types.TaskDefinition {
	output := make(chan int)
	return func(input chan int) chan int {
		v := <-input
		fn(v, output)
		close(output)
		return output
	}

}
