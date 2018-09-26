package timeid

import (
	"sync"
	"time"
)

var mu sync.Mutex

// Next, monotonically increasing timestamp ID
func Next() int64 {
	mu.Lock()
	defer mu.Unlock()
	return time.Now().UnixNano()
}
