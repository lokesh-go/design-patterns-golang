package factory

type dbMethods interface {
	GetClient() string
}

type fileMethods interface {
	ReadFile() string
}

type db struct{}
type file struct{}

func newdb() dbMethods {
	return &db{}
}

func newfile() fileMethods {
	return &file{}
}

// GetClient: default response
func (d *db) GetClient() (res string) {
	return "Invalid db name Input"
}

// ReadFile: default response
func (f *file) ReadFile() (res string) {
	return "Invalid file extension Input"
}
