package strategy

type cacheEvictAlgo interface {
	evict(c *cache)
}
