package mongodb

// Methods ...
type Methods interface {
	GetClient() string
}

type mongodb struct {
}

// New ...
func New() Methods {
	return &mongodb{}
}

// GetClient ...
func (m *mongodb) GetClient() (client string) {
	return "This is mongodb client"
}
