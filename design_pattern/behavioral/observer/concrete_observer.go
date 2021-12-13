package observer

type ConcreteObserver struct {
	id   int
	high int
}

func (c *ConcreteObserver) update(value int) {
	c.high = value + 1
}

func (c *ConcreteObserver) getID() int {
	return c.getID()
}
