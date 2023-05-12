package main

import (
	"fmt"

	"github.com/forrestbthomas/curly-parakeet/pkg/pipeline"
	fi "github.com/forrestbthomas/curly-parakeet/pkg/task/fanin"
	fo "github.com/forrestbthomas/curly-parakeet/pkg/task/fanout"
	p "github.com/forrestbthomas/curly-parakeet/pkg/task/parallel"
	"github.com/forrestbthomas/curly-parakeet/pkg/types"
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

	// Pipeline Definition
	FanIn := fi.Task{}
	FanOut := fo.Task{}
	Parallel := p.Task{}
	pipeMap := pipeline.TaskMap{
		FanIn: []types.TaskWork{
			fi.Sum,
		},
		FanOut: []types.TaskWork{
			fo.ListMultiples,
		},
		Parallel: []types.TaskWork{
			p.Doubler,
			p.Filter,
			p.Tripler,
			p.DoubleLen,
		},
	}

	// Pipeline
	pipe := pipeline.New()
	tasks := pipe.Generator(pipeMap)

	// generate input task
	input := InputTask(ints)

	// run tasks
	for _, task := range tasks {
		input = task(input)
	}

	fmt.Println(len(input))

}
