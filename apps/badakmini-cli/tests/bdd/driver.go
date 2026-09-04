package bdd

import (
	"context"
	"reflect"
)

// Result is the complete public command observation used by behaviour assertions.
type Result struct {
	ExitCode int
	Stdout   string
	Stderr   string
}

// Driver varies only the boundary used to prepare and invoke Badak Mini.
type Driver interface {
	Prepare(ctx context.Context, fixture string) error
	Invoke(ctx context.Context, args []string) error
	Result() Result
}

// DriverFactory gives every scenario a fresh adapter and isolated resources.
type DriverFactory func() Driver

// MissingDriverMethods reports contract operations absent from a candidate adapter type.
func MissingDriverMethods(candidate reflect.Type) []string {
	contract := reflect.TypeFor[Driver]()
	missing := make([]string, 0)
	for method := range contract.Methods() {
		if _, present := candidate.MethodByName(method.Name); !present {
			missing = append(missing, method.Name)
		}
	}

	return missing
}
