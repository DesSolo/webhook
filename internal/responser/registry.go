package responser

var registry = map[string]Responser{}

// Register register responser
func Register(r Responser) {
	registry[r.Kind()] = r
}

// Get get responser from registry
func Get(kind string) (Responser, bool) {
	r, ok := registry[kind]
	return r, ok
}
