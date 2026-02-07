package pipeline

type StateManager struct {
	values []int
}

func NewStateManager() *StateManager {
	return &StateManager{values: []int{}}
}

func (s *StateManager) Append(v int) {
	s.values = append(s.values, v)
}

func (s *StateManager) Values() []int {
	// επιστρέφουμε copy για immutability
	cp := make([]int, len(s.values))
	copy(cp, s.values)
	return cp
}
