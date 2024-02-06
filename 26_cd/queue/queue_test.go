package todo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue_Push(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "empty queue",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "single item",
			input: []string{"item1"},
			want:  []string{"item1"},
		},
		{
			name:  "multiple items",
			input: []string{"item1", "item2", "item3"},
			want:  []string{"item1", "item2", "item3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			q := NewQueue[string]()

			// act
			for _, i := range tc.input {
				q.Push(i)
			}

			// assert
			require.Equal(t, len(tc.want), len(q.items))
			require.Equal(t, tc.want, q.items)
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "empty queue",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "single item",
			input: []string{"item1"},
			want:  []string{},
		},
		{
			name:  "multiple items",
			input: []string{"item1", "item2", "item3"},
			want:  []string{"item2", "item3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			q := NewQueue[string]()
			q.items = tc.input

			// act
			q.Pop()

			// assert
			require.Equal(t, len(tc.want), len(q.items))
			require.Equal(t, tc.want, q.items)
		})
	}
}
