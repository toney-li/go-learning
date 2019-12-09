package rwmutex

import (
	"sync"
	"testing"
)

func TestRmutex(t *testing.T) {
	var mutex *sync.RWMutex
	mutex = new(sync.RWMutex)
	mutex.RLock()
	mutex.Lock()
	mutex.RUnlock()
	mutex.Unlock()
}
