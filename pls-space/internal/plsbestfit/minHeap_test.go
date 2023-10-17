package plsbestfit

import (
	"container/heap"
	"testing"
)

func TestMinHeap(t *testing.T) {
	testCases := []struct {
		name     string
		input    MinHeap
		expected MinHeap
	}{
		{
			name: "Test case 1",
			input: MinHeap{
				&ParkingSpot{SpotNumber: 3},
				&ParkingSpot{SpotNumber: 1},
				&ParkingSpot{SpotNumber: 2},
			},
			expected: MinHeap{
				&ParkingSpot{SpotNumber: 1},
				&ParkingSpot{SpotNumber: 2},
				&ParkingSpot{SpotNumber: 3},
			},
		},
		{
			name:     "Test case 2",
			input:    MinHeap{},
			expected: MinHeap{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			inputHeap := make(MinHeap, len(testCase.input))
			copy(inputHeap, testCase.input)

			heap.Init(&inputHeap)

			for _, expectedSpot := range testCase.expected {
				if len(inputHeap) == 0 {
					t.Error("Expected more elements in the heap")
					break
				}

				popped := heap.Pop(&inputHeap).(*ParkingSpot)
				if popped.SpotNumber != expectedSpot.SpotNumber {
					t.Errorf("Expected SpotNumber %d, got %d", expectedSpot.SpotNumber, popped.SpotNumber)
				}
			}

			if len(inputHeap) != 0 {
				t.Error("Expected no more elements in the heap")
			}
		})
	}
}
