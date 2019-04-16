package main

import "sync"

// Storage is the struct which actually manages the data and handles all read/write operations
type Storage struct {
	hashmap map[string]string
	lock    sync.RWMutex
}

func (s *Storage) putValue(entry *Pair) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.hashmap[entry.Key] = entry.Value
}

func (s *Storage) getValue(key string) (string, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	i, ok := s.hashmap[key]

	return i, ok
}
