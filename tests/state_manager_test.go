package tests

import (
	"guess-it/pipeline"
	"testing"
)

// TestStateManagerStartEmpty verifies state manager starts with no values
func TestStateManagerStartEmpty(t *testing.T) {
	// Create new state manager
	state := pipeline.NewStateManager()

	// Get initial values
	values := state.Values()
	// Verify empty state
	if len(values) != 0 {
		t.Fatalf("expected empty state, got %d values", len(values))
	}
}

// TestStateManagerAppendSingleValue verifies appending one value
func TestStateManagerAppendSingleValue(t *testing.T) {
	state := pipeline.NewStateManager()

	// Append a single value
	state.Append(42)

	// Retrieve values
	values := state.Values()
	// Verify count
	if len(values) != 1 {
		t.Fatalf("expected 1 element, got %d", len(values))
	}
	// Verify value
	if values[0] != 42 {
		t.Fatalf("expected value 42, got %d", values[0])
	}
}

// TestStateManagerAppendMultipleValues verifies appending multiple values
func TestStateManagerAppendMultipleValues(t *testing.T) {
	state := pipeline.NewStateManager()

	// Append three values
	state.Append(10)
	state.Append(20)
	state.Append(30)

	// Retrieve values
	values := state.Values()
	// Verify count
	if len(values) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(values))
	}

	// Verify each value
	expected := []int{10, 20, 30}
	for i, v := range expected {
		if values[i] != v {
			t.Fatalf("expected %d at index %d, got %d", v, i, values[i])
		}
	}
}

// TestStateManagerPreservesOrder verifies insertion order is maintained
func TestStateManagerPreservesOrder(t *testing.T) {
	state := pipeline.NewStateManager()

	// Append values in specific order
	state.Append(5)
	state.Append(1)
	state.Append(8)

	// Retrieve values
	values := state.Values()

	// Verify order is preserved
	if values[0] != 5 || values[1] != 1 || values[2] != 8 {
		t.Fatalf("state did not preserve insertion order: %v", values)
	}
}

// TestStateManagerImmutability verifies returned values cannot modify internal state
func TestStateManagerImmutability(t *testing.T) {
	state := pipeline.NewStateManager()
	state.Append(7)

	// Get values and attempt to modify
	values := state.Values()
	values[0] = 99 // Attempt mutation

	// Get fresh copy
	fresh := state.Values()
	// Verify internal state unchanged
	if fresh[0] != 7 {
		t.Fatalf("state should be immutable, expected 7 got %d", fresh[0])
	}
}

// TestStateManagerMultipleSnapshots verifies snapshots are independent
func TestStateManagerMultipleSnapshots(t *testing.T) {
	state := pipeline.NewStateManager()
	state.Append(1)
	state.Append(2)

	// Take first snapshot
	snap1 := state.Values()
	// Add more data
	state.Append(3)
	// Take second snapshot
	snap2 := state.Values()

	// Verify snapshots have different lengths
	if len(snap1) != 2 {
		t.Fatalf("expected snapshot length 2, got %d", len(snap1))
	}
	if len(snap2) != 3 {
		t.Fatalf("expected snapshot length 3, got %d", len(snap2))
	}
}

// TestStateManagerDoesNotExposeInternalSlice verifies external modifications don't affect state
func TestStateManagerDoesNotExposeInternalSlice(t *testing.T) {
	state := pipeline.NewStateManager()
	state.Append(100)

	// Get values and append externally
	values := state.Values()
	values = append(values, 200)

	// Get fresh copy
	fresh := state.Values()
	// Verify external append didn't affect internal state
	if len(fresh) != 1 {
		t.Fatalf("external append should not affect internal state")
	}
}
