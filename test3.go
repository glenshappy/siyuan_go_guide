package main

import (
	"fmt"
	"sync"
)

type Stats struct {
	mu sync.Mutex

	counters map[string]int
}

// Snapshot 返回当前状态。
func (s *Stats) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.counters
}



func  main()  {
	var stats *Stats = &Stats{counters: map[string]int{"haha":1}}
	// snapshot 不再受互斥锁保护
	// 因此对 snapshot 的任何访问都将受到数据竞争的影响
	// 影响 stats.counters
	snapshot := stats.Snapshot()
	snapshot["bbbb"]=2
	fmt.Printf("vvvstats:%#v",stats)
}
