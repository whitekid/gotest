package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceRefs(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}
	s := a[2:3]
	s[0] = 999

	require.Equal(t, []int{0, 1, 999, 3, 4}, a)
}

func TestSliceCopy1(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}
	s := make([]int, len(a))
	copy(s, a)
	s[0] = 999

	require.Equal(t, []int{0, 1, 2, 3, 4}, a)
}

func TestSliceCopy2(t *testing.T) {
	s1 := []int{0, 1, 2, 3}
	s2 := make([]int, len(s1))
	n1 := copy(s2, s1)

	require.EqualValues(t, s1, s2)
	require.Equal(t, len(s1), n1)
}

func TestSliceAppend(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	a := s[2:3]

	require.Equal(t, 1, len(a), printSlice(a))
}

func TestSliceRemoveElement(t *testing.T) {
	type args struct {
		s  []int
		at int
	}

	tests := [...]struct {
		name      string
		args      args
		wantPanic bool
		want      []int
	}{
		{"default", args{[]int{0, 1, 2, 3, 4}, 1}, false, []int{0, 2, 3, 4}},
		{"begin", args{[]int{0, 1, 2, 3, 4}, 0}, false, []int{1, 2, 3, 4}},
		{"end", args{[]int{0, 1, 2, 3, 4}, 4}, false, []int{0, 1, 2, 3}},
		{"out", args{[]int{0, 1, 2, 3, 4}, 7}, true, []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				require.True(t, (recover() != nil) == tt.wantPanic, "wantPanic = %t but panic occured.", tt.wantPanic)
			}()

			s := removeAt(tt.args.s, tt.args.at)

			require.EqualValues(t, tt.want, s)
		})
	}
}
