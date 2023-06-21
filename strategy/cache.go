package strategy

import "fmt"

type cache struct {
	storage      map[string]string
	maxCapacity  int
	evictionAlgo cacheEvictAlgo
}

func New(e cacheEvictAlgo) *cache {
	s := map[string]string{}
	return &cache{
		storage:      s,
		maxCapacity:  3,
		evictionAlgo: e,
	}
}

func (c *cache) Add(k, v string) {
	if len(c.storage) == c.maxCapacity {
		c.evictionAlgo.evict(c)
	}
	c.storage[k] = v
}

func (c *cache) SetEvictionAlgo(e cacheEvictAlgo) {
	c.evictionAlgo = e
}

func (c *cache) Print() {
	for k, v := range c.storage {
		fmt.Println("Key: ", k, "Value: ", v)
	}
}
