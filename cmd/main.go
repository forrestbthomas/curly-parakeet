package main

import (
	"fmt"

	"github.com/forrestbthomas/curly-parakeet/pkg/examples"
	"github.com/forrestbthomas/curly-parakeet/pkg/pipeline"
	fi "github.com/forrestbthomas/curly-parakeet/pkg/task/fanin"
)

func InputTask(ints []int) chan int {
	output := make(chan int, len(ints))
	defer close(output)
	for _, val := range ints {
		output <- val
	}
	return output
}

func main() {
	// incoming data
	ints := []int{1, 2, 3, 4, 5}

	// Pipeline
	jobs := []pipeline.Job{
		{
			Fn:    examples.Sum,
			Task:  fi.Task{},
			Needs: nil,
		},
	}
	pipe := pipeline.New(jobs)

	// generate input task
	input := InputTask(ints)

	// run tasks
	for _, task := range pipe {
		input = task(input)
	}

	fmt.Println(<-input)

}
