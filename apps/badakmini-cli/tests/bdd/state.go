package bdd

import "sync"

// State owns one scenario's adapter and last observable result.
type State struct {
	mu     sync.Mutex
	driver Driver
}

// NewState creates isolated state with a fresh adapter.
func NewState(factory DriverFactory) *State {
	return &State{driver: factory()}
}

// Driver returns the adapter scoped to this scenario.
//
//nolint:ireturn // Scenario state deliberately exposes the shared adapter contract.
func (state *State) Driver() Driver {
	state.mu.Lock()
	defer state.mu.Unlock()

	return state.driver
}
