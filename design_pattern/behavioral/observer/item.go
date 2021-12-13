package observer

type Item struct {
	observers []Observer
	value     int
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) deregister(observer Observer) {

}

func (i *Item) notify() {
	for _, observer := range i.observers {
		observer.update(i.value)
	}
}
