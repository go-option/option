package option

func New[T any, O ~func(*T)](options []O, defaults ...O) *T {
	opts := new(T)
	for _, option := range defaults {
		option(opts)
	}
	for _, option := range options {
		option(opts)
	}
	return opts
}
