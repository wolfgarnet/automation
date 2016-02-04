package registry

var factories map[string]func(map[string]interface{}) interface{}

type Factory func(map[string]interface{}) interface{}

func AddFactory(name string, factory Factory) {
	factories[name] = factory
}

func GetFactory(name string) Factory {
	f, ok := factories[name]
	if !ok {
		return nil
	}
	return f
}
