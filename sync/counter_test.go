package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.inc()
		counter.inc()
		counter.inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCounter := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCounter)

		for i := 0; i < wantedCounter; i++ {
			go func() {
				counter.inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, 1000)
	})

}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d,want %d", got.Value(), want)
	}
}
