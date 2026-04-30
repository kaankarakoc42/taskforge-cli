package executor

var registry = map[string]Executor{}

func Register(name string, e Executor) {
	registry[name] = e
}

func Get(name string) Executor {
	return registry[name]
}
