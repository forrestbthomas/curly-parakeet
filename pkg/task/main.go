package task

import "github.com/forrestbthomas/curly-parakeet/pkg/types"

type Tasker interface {
	GetName() string
	Generator(fn types.TaskWork) types.TaskDefinition
}
