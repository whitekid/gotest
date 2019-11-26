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

func TestSliceCopy(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}
	s := make([]int, len(a))
	copy(s, a)
	s[0] = 999

	require.Equal(t, []int{0, 1, 2, 3, 4}, a)
}
