package regal

type paramOption struct {
	combine  int
	schedule int
}

type ParamOption interface {
	apply(option *paramOption)
}

type funcOption struct {
	f func(option *paramOption)
}

func (fdo *funcOption) apply(do *paramOption) {
	fdo.f(do)
}

func newFuncOption(f func(*paramOption)) *funcOption {
	return &funcOption{f: f}
}

func withCombine(s int) ParamOption {
	return newFuncOption(func(o *paramOption) {
		o.combine = s
	})
}

func withSchedule(s int) ParamOption {
	return newFuncOption(func(o *paramOption) {
		o.schedule = s
	})
}
