package executor

import (
	"sort"
	"sync"
)

var (
	mu       sync.RWMutex
	registry = make(map[string]Executor)
)

func Register(e Executor) {
	if e == nil {
		return
	}

	name := e.Name()
	if name == "" {
		return
	}

	mu.Lock()
	defer mu.Unlock()
	registry[name] = e
}

func Get(name string) (Executor, bool) {
	mu.RLock()
	defer mu.RUnlock()
	e, ok := registry[name]
	return e, ok
}

func List() []Executor {
	mu.RLock()
	defer mu.RUnlock()

	executors := make([]Executor, 0, len(registry))
	for _, e := range registry {
		executors = append(executors, e)
	}

	sort.Slice(executors, func(i, j int) bool {
		return executors[i].Name() < executors[j].Name()
	})

	return executors
}
