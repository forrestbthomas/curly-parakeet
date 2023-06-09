package main

import (
	"fmt"

	"github.com/forrestbthomas/curly-parakeet/pkg/examples"
	"github.com/forrestbthomas/curly-parakeet/pkg/pipeline"
	fi "github.com/forrestbthomas/curly-parakeet/pkg/task/fanin"
	fo "github.com/forrestbthomas/curly-parakeet/pkg/task/fanout"
	p "github.com/forrestbthomas/curly-parakeet/pkg/task/parallel"
)

func InputTask(ints []int) chan int {
	output := make(chan int, len(ints))
	for _, val := range ints {
		output <- val
	}
	fmt.Println("making input", len(output))
	return output
}

func main() {
	// incoming data
	ints := []int{1, 2, 3, 4, 5}

	// Pipeline
	jobs := []pipeline.Job{
		{
			Fn:    examples.Sum,
			Task:  &fi.FanIn{},
			Needs: nil,
		},
		{
			Fn:    examples.ListOdds,
			Task:  &fo.FanOut{},
			Needs: nil,
		},
		{
			Fn:    examples.Tripler,
			Task:  &p.Parallel{},
			Needs: nil,
		},
		{
			Fn:    examples.Sum,
			Task:  &fi.FanIn{},
			Needs: nil,
		},
		{
			Fn:    examples.ListOdds,
			Task:  &fo.FanOut{},
			Needs: nil,
		},
	}

	pipe := pipeline.New(jobs)
	ch := InputTask(ints)
	close(ch)

	out := pipe.Run(ch)
	for el := range out {
		fmt.Println(el)
	}

}
