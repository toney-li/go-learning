package singleton

import (
	"fmt"
	"sync"
)

// Singleton is a comment
type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func Instance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
func (s *Singleton) do() {
	fmt.Println()
}
