package observer

// Item ...
type item struct {
	list    []UserObserver
	name    string
	inStock bool
}

func New(name string) *item {
	return &item{name: name}
}

func (i *item) RegisterUserObserver(u UserObserver) {
	i.list = append(i.list, u)
}

func (i *item) UpdateAvailability() {
	i.inStock = true
	i.notifyAll()
}

func (i *item) notifyAll() {
	for _, userObserver := range i.list {
		userObserver.send(i.name)
	}
}
