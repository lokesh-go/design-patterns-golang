package decorator

type LaptopWithGraphicCard struct {
	Laptop Laptop
}

func (l *LaptopWithGraphicCard) GetPrice() int {
	nvidia940MX := 25000
	return l.Laptop.GetPrice() + nvidia940MX
}

type LaptopWith16GBMemory struct {
	laptop Laptop
}

func (l *LaptopWith16GBMemory) GetPrice() int {
	ram16GB := 5000
	return l.laptop.GetPrice() + ram16GB
}
