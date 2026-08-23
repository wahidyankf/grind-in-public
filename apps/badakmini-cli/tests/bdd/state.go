package bdd

import (
	"context"
	"errors"
	"sync"
)

type scenarioStateKey struct{}

// State owns one scenario's adapter and last observable result.
type State struct {
	mu     sync.Mutex
	driver Driver
}

// NewState creates isolated state with a fresh adapter.
func NewState(factory DriverFactory) *State {
	return &State{driver: factory()}
}

func contextWithState(ctx context.Context, state *State) context.Context {
	return context.WithValue(ctx, scenarioStateKey{}, state)
}

func stateFromContext(ctx context.Context) (*State, error) {
	state, ok := ctx.Value(scenarioStateKey{}).(*State)
	if !ok {
		return nil, errors.New("godog scenario context has no Badak Mini state")
	}
	return state, nil
}

// Driver returns the adapter scoped to this scenario.
//
//nolint:ireturn // Scenario state deliberately exposes the shared adapter contract.
func (state *State) Driver() Driver {
	state.mu.Lock()
	defer state.mu.Unlock()

	return state.driver
}
