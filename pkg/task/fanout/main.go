package task

import (
	"log"

	"github.com/forrestbthomas/curly-parakeet/pkg/task"
)

type FanOut struct {
	state int
}

func (f *FanOut) Get(s string) int {
	switch s {
	case "state":
		return f.state
	}
	return 0
}

func (f *FanOut) Set(s string, v int) {
	switch s {
	case "state":
		f.state = v
	}
}

func (f *FanOut) Generator(fn task.TaskWork) task.TaskDefinition {
	done := make(chan int, 1)
	tmpSlice := []int{}
	tmpChan := make(chan int)
	return func(input chan int) chan int {
		if len(input) == 0 || len(input) > 1 {
			log.Fatal("Cannot FanOut with less than or more than 1 element: has ", len(input), " elements")
		}
		go func() {
			for el := range tmpChan {
				tmpSlice = append(tmpSlice, el)
			}
			done <- 1
			close(done)
		}()
		fn(<-input, tmpChan, f)
		<-done
		output := make(chan int, len(tmpSlice))
		for _, v := range tmpSlice {
			output <- v
		}
		return output
	}

}
