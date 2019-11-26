package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	ErrNotFound  = errors.New("Not Found")
	ErrForbidden = errors.New("Forbidden")
)

func NotFoundError() error {
	return ErrNotFound
}

type SampleError struct {
	err error
}

func NewSampleError(e error) error {
	return &SampleError{
		err: e,
	}
}

func (e *SampleError) Error() string {
	return fmt.Sprintf("sample error: %w", e.err)
}

func (e *SampleError) Unwrap() error { return e.err }

func TestErrorWrap(t *testing.T) {
	e := NewSampleError(NotFoundError())

	require.True(t, errors.Is(e, ErrNotFound))
}
