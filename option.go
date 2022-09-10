package option

func New[T any, O ~func(*T)](options []O, defaults ...O) *T {
	return Apply(nil, options, defaults...)
}

func Apply[T any, O ~func(*T)](opts *T, options []O, defaults ...O) *T {
	if opts == nil {
		opts = new(T)
	}
	for _, option := range defaults {
		option(opts)
	}
	for _, option := range options {
		option(opts)
	}
	return opts
}
