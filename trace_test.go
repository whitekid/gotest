package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

// run with go test -trace=trace.out
// "net/http/pprof" for trace standard HTTP interface
// then run "go tool trace trace.out" to view trace
//
// region is for logging time interval during goroutine's execution
// tasks is tracing logical operations such as RPC request, may require multiple goroutine

func TestTraceHTTP(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func fetchURL(s string) error {
	req, err := http.NewRequest(http.MethodGet, s, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed with status: %d", resp.StatusCode)
	}

	return nil
}

func regionA() {
	fetchURL("https://google.com")
	fetchURL("https://daum.net")
}
func taskA() {
	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(2)
	go func() {
		fetchURL("https://google.com")
		wg.Done()
	}()
	go func() {
		fetchURL("https://daum.net")
		wg.Done()
	}()
}

func TestTrace(t *testing.T) {
	ctx := context.Background()
	trace.Log(ctx, "TestTrace", "begin")
	trace.WithRegion(ctx, "regionA", regionA)

	ctx, task := trace.NewTask(ctx, "fetch")
	trace.Log(ctx, "fetch", "URL")
	taskA()
	task.End()
}
