package main

import (
	"errors"
	"sync"
)

type Employee struct {
	ID     int    `json:"id"`
	NAME   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(is int, e Employee)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()

	e.ID = s.counter
	s.data[e.ID] = *e

	s.counter++

	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()

	res, ok := s.data[id]
	if !ok {
		return res, errors.New("employee not found")
	}
	defer s.Unlock()
	return res, nil
}

func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()

	s.data[id] = e

	s.Unlock()
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()

	delete(s.data, id)
	s.counter--

	s.Unlock()
}
