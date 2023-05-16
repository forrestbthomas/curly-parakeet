package examples

import (
	"fmt"

	"github.com/forrestbthomas/curly-parakeet/pkg/task"
)

// Fan In
func Sum(i int, _ chan int, t task.Tasker) {
	fmt.Println("adding", i, t.Get("state"))
	t.Set("state", t.Get("state")+i)
}

// Fan Out
func ListOdds(i int, ch chan int, _ task.Tasker) {
	fmt.Println("listing odds", i)
	for i > 0 {
		if i%2 != 0 {
			ch <- i
		}
		i--
	}
	close(ch) // required: impossible to know buffer len ahead of time, so sender must call close
}

// Parallel
func Filter(i int, ch chan int, _ task.Tasker) {
	fmt.Println("filtering", i)
	if i%2 != 0 {
		ch <- i
	}
}

func Doubler(i int, ch chan int, _ task.Tasker) {
	fmt.Println("doubling", i)
	ch <- i * 2
}

func Tripler(i int, ch chan int, _ task.Tasker) {
	fmt.Println("tripling", i)
	ch <- i * 3
}
