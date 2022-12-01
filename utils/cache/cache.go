package cache

import "log"

type cacheGen struct {
	set   bool
	New   func() any
	Value any
}

var data = map[string]*cacheGen{}

func Get[T any](key string) T {
	gen, ok := data[key]
	if !ok {
		log.Fatal("read from nil key:", key)
	}
	if !gen.set {
		gen.Value = gen.New()
		gen.set = true
	}
	val, ok := gen.Value.(T)
	if !ok {
		log.Fatal("cast failed for key:", key)
	}
	return val
}

func Set(key string, new func() any) {
	data[key] = &cacheGen{
		New: new,
	}
}

func Put(key string, value any) {
	data[key] = &cacheGen{
		set:   true,
		Value: value,
	}
}
