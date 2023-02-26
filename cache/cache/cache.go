package cache

import "github.com/luizmoitinho/generics/blog/entities"

type Cacheable interface {
	entities.Category | entities.Post
}

type storage[T Cacheable] struct {
	data map[string]T
}

func New[T Cacheable]() *storage[T] {
	c := storage[T]{}
	c.data = make(map[string]T)
	return &c
}

func (s *storage[T]) Set(key string, value T) {
	s.data[key] = value
}

func (s *storage[T]) Get(key string) (value T) {
	return s.data[key]
}
