package strategy

import "fmt"

type Fifo struct {
}

func (f *Fifo) evict(c *cache) {
	fmt.Println("Eviction by FIFO")

	for k := range c.storage {
		// delete first key
		delete(c.storage, k)
		break
	}
}
