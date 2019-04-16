package main

import "sync"

// Storage is the struct which actually manages the data and handles all read/write operations
type Storage struct {
	hashmap map[string]string
	lock    sync.RWMutex
}

func (s *Storage) putValue(entry *Pair) {
	s.lock.Lock()
	s.hashmap[entry.Key] = entry.Value
	s.lock.Unlock()
}
