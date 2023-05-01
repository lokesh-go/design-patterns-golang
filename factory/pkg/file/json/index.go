package json

type methods interface {
	ReadFile() string
}

type json struct{}

// New ...
func New() methods {
	return &json{}
}

// ReadFile ...
func (j *json) ReadFile() (res string) {
	return "You are reading json file"
}
