package observer

type subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notify()
}
