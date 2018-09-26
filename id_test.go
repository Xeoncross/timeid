package timeid

import (
	"testing"
)

// Not really a valid concurrent test, just here for code coverage
func TestNext(t *testing.T) {

	prev := Next()

	for i := 0; i < 100000; i++ {
		id := Next()
		if id <= prev {
			t.Error("failed to increase")
		}
		prev = id
	}

}

func BenchmarkNext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Next()
	}
}
