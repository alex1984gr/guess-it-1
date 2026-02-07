package pipeline

// StateManager stores historical values for prediction
type StateManager struct {
	values []int
}

// NewStateManager creates a new StateManager instance
func NewStateManager() *StateManager {
	return &StateManager{values: []int{}}
}

// Append adds a new value to the historical data
func (s *StateManager) Append(v int) {
	s.values = append(s.values, v)
}

// Values returns a copy of all stored values for immutability
func (s *StateManager) Values() []int {
	// Return a copy to prevent external modification
	cp := make([]int, len(s.values))
	copy(cp, s.values)
	return cp
}
