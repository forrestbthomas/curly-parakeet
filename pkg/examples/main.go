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
	close(ch)
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

func DoubleLen(i int, ch chan int, _ task.Tasker) {
	fmt.Println("doubling length", i)
	ch <- i
	ch <- i
}
