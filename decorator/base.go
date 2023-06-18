package decorator

type Laptop interface {
	GetPrice() int
}

type WindowsLaptop struct{}

func (w *WindowsLaptop) GetPrice() int {
	return 30000
}
