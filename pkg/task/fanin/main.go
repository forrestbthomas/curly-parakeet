package task

import (
	"fmt"
	"log"
	"sync"

	"github.com/forrestbthomas/curly-parakeet/pkg/task"
)

type FanIn struct {
	state int
}

func (f *FanIn) Set(s string, v int) {
	switch s {
	case "state":
		f.state = v
	default:
		log.Fatal("invalid field to set: ", s)
	}
}

func (f *FanIn) Get(s string) int {
	switch s {
	case "state":
		return f.state
	default:
		log.Fatal("invalid field to get: ", s)
	}
	return 0
}

func (f *FanIn) Generator(fn task.TaskWork) task.TaskDefinition {
	output := make(chan int, 1)
	var wg sync.WaitGroup
	return func(input chan int) chan int {
		fmt.Println("fanning in")
		if len(input) <= 1 {
			log.Fatal("trying to fan in on less than two elements")
		}
		wg.Add(1)
		for el := range input {
			fn(el, output, f)
		}

		go func(w *sync.WaitGroup, c chan int) {
			c <- int(f.Get("state"))
			close(c)
			w.Done()
		}(&wg, output)
		f.Set("state", 0)
		return output
	}

}
