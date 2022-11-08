package types

import (
	"sort"
	"sync"

	"golang.org/x/exp/maps"
)

func MakeSafeCommitHash(hash map[int]CommitStore) *SafeCommitHash {
	return &SafeCommitHash{hash: hash}
}

type SafeCommitHash struct {
	Mutex sync.Mutex
	hash  map[int]CommitStore
}

func (s *SafeCommitHash) Write(key int, value CommitStore) {
	s.Mutex.Lock()
	s.hash[key] = value
	s.Mutex.Unlock()
}

func (s *SafeCommitHash) Read(key int) (CommitStore, bool) {
	s.Mutex.Lock()
	if entry, ok := s.hash[key]; ok {
		s.Mutex.Unlock()
		return entry, true
	}
	s.Mutex.Unlock()
	return CommitStore{}, false
}

// potential optimization: write a specific delete function (the for loop in commitCalculator.Delete) that only locks/unlocks once
func (s *SafeCommitHash) Delete(key int) {
	s.Mutex.Lock()
	delete(s.hash, key)
	s.Mutex.Unlock()
}

func (s *SafeCommitHash) GetOffsets() []int {
	offsets := maps.Keys(s.hash)
	sort.Ints(offsets)
	return offsets
}
