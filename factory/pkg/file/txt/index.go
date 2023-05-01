package txt

type methods interface {
	ReadFile() string
}

type txt struct{}

// New ...
func New() methods {
	return &txt{}
}

// ReadFile ...
func (t *txt) ReadFile() (res string) {
	return "You are reading txt file"
}
