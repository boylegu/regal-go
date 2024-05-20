package regal

type paramOption struct {
	combine  int
	schedule int
	priorKey string
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

func WithCombine(s int) ParamOption {
	return newFuncOption(func(o *paramOption) {
		o.combine = s
	})
}

func WithSchedule(s int) ParamOption {
	return newFuncOption(func(o *paramOption) {
		o.schedule = s
	})
}

func WithPriorKey(s string) ParamOption {
	return newFuncOption(func(o *paramOption) {
		o.priorKey = s
	})
}

func defaultParams() paramOption {
	return paramOption{
		combine:  1,
		schedule: 1,
		priorKey: "",
	}
}
