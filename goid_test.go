package goid_test

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/t-yuki/goid"
)

func TestID(t *testing.T) {
	testID(t)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			testID(t)
			wg.Done()
		}()
	}
	wg.Wait()
}

func testID(t *testing.T) {
	n := goid.ID()
	lines := strings.Split(stackTrace(), "\n")
	for i, line := range lines {
		if !strings.HasPrefix(line, fmt.Sprintf("goroutine %d ", n)) {
			continue
		}
		if i+1 == len(lines) {
			break
		}
		if !strings.Contains(lines[i+1], ".stackTrace") {
			t.Errorf("there are goroutine id %d but it is not me: %s", n, lines[i+1])
		}
		return
	}
	t.Errorf("there are no goroutine %d", n)
}

func stackTrace() string {
	var n int
	for n = 4096; n < 16777216; n *= 2 {
		buf := make([]byte, n)
		ret := runtime.Stack(buf, true)
		if ret != n {
			return string(buf[:ret])
		}
	}
	panic(n)
}
