package tests

import (
	"guess-it/pipeline"
	"testing"
)

func TestStateManagerStartEmpty(t *testing.T) {
	state := pipeline.NewStateManager()

	values := state.Values()
	if len(values) != 0 {
		t.Fatalf("expected empty state, got %d values", len(values))
	}
}

func TestStateManagerAppendSingleValue(t *testing.T) {
	state := pipeline.NewStateManager()

	state.Append(42)

	values := state.Values()
	if len(values) != 1 {
		t.Fatalf("expected 1 element, got %d", len(values))
	}
	if values[0] != 42 {
		t.Fatalf("expected value 42, got %d", values[0])
	}
}

func TestStateManagerAppendMultipleValues(t *testing.T) {
	state := pipeline.NewStateManager()

	state.Append(10)
	state.Append(20)
	state.Append(30)

	values := state.Values()
	if len(values) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(values))
	}

	expected := []int{10, 20, 30}
	for i, v := range expected {
		if values[i] != v {
			t.Fatalf("expected %d at index %d, got %d", v, i, values[i])
		}
	}
}

func TestStateManagerPreservesOrder(t *testing.T) {
	state := pipeline.NewStateManager()

	state.Append(5)
	state.Append(1)
	state.Append(8)

	values := state.Values()

	if values[0] != 5 || values[1] != 1 || values[2] != 8 {
		t.Fatalf("state did not preserve insertion order: %v", values)
	}
}

func TestStateManagerImmutability(t *testing.T) {
	state := pipeline.NewStateManager()
	state.Append(7)

	values := state.Values()
	values[0] = 99 // attempt mutation

	fresh := state.Values()
	if fresh[0] != 7 {
		t.Fatalf("state should be immutable, expected 7 got %d", fresh[0])
	}
}

func TestStateManagerMultipleSnapshots(t *testing.T) {
	state := pipeline.NewStateManager()
	state.Append(1)
	state.Append(2)

	snap1 := state.Values()
	state.Append(3)
	snap2 := state.Values()

	if len(snap1) != 2 {
		t.Fatalf("expected snapshot length 2, got %d", len(snap1))
	}
	if len(snap2) != 3 {
		t.Fatalf("expected snapshot length 3, got %d", len(snap2))
	}
}

func TestStateManagerDoesNotExposeInternalSlice(t *testing.T) {
	state := pipeline.NewStateManager()
	state.Append(100)

	values := state.Values()
	values = append(values, 200)

	fresh := state.Values()
	if len(fresh) != 1 {
		t.Fatalf("external append should not affect internal state")
	}
}
