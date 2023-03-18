package cache

// Methods ...
type Methods interface {
	GetClient() string
}

type cache struct {
}

// New ...
func New() Methods {
	return &cache{}
}

// GetClient ...
func (c *cache) GetClient() (client string) {
	return "This is cache client"
}
