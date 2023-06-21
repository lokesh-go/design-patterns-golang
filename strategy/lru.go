package strategy

import "fmt"

type LRU struct {
}

func (l *LRU) evict(c *cache) {
	fmt.Println("Eviction by LRU")
}
