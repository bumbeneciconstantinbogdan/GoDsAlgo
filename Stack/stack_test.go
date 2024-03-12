package dataStructures

import (
	"fmt"
	"testing"
)

// Test push function for Stack implementation using slices
func TestPush(t *testing.T) {
	s := Stack[int]{}

	type Test struct {
		push int
		want int
	}

	tests := []Test{
		{push: 1, want: 1},
		{push: 2, want: 2},
		{push: 3, want: 3},
	}

	for index, tt := range tests {
		testname := fmt.Sprintf("Pushed %v", tt.push)
		t.Run(testname, func(t *testing.T) {
			s.Push(tt.push)
			if s.data[index] != tt.want {
				t.Errorf("pushed %v, want %v", tt.push, tt.want)
			}
		})
	}
}

// Test pop function for Stack implementation using slices
func TestPop(t *testing.T) {
	s := Stack[int]{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	type Test struct {
		want1 int
		want2 bool
	}

	tests := []Test{
		{want1: 3, want2: true},
		{want1: 2, want2: true},
		{want1: 1, want2: true},
		{want1: 0, want2: false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Poped %v, Has elements, %v", tt.want1, tt.want2)
		t.Run(testname, func(t *testing.T) {
			poped, hasElements := s.Pop()
			if poped != tt.want1 || hasElements != tt.want2 {
				t.Errorf("Poped %v, want %v, Has element %v, wanted %v", poped, tt.want1, hasElements, tt.want2)
			}
		})
	}
}

func TestSize(t *testing.T) {
	s := Stack[int]{}

	type Test struct {
		push int
		want int
	}

	tests := []Test{
		{push: 1, want: 1},
		{push: 1, want: 2},
		{push: 1, want: 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Pushed %v", tt.push)
		t.Run(testname, func(t *testing.T) {
			s.Push(tt.push)
			size := s.Size()
			if size != tt.want {
				t.Errorf("size %v, want %v", size, tt.want)
			}
		})
	}
}
