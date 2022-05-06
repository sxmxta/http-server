package common

import (
	"container/list"
	"errors"
	"sync"
)

type Queue struct {
	mutex sync.Mutex
	list  *list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (m *Queue) Push(data interface{}) {
	if data == nil {
		return
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.list.PushBack(data)
}

func (m *Queue) Pop() (interface{}, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if element := m.list.Front(); element != nil {
		m.list.Remove(element)
		return element.Value, nil
	}
	return nil, errors.New("pop failed")
}

func (m *Queue) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for element := m.list.Front(); element != nil; {
		elementNext := element.Next()
		m.list.Remove(element)
		element = elementNext
	}
}

func (m *Queue) Len() int {
	return m.list.Len()
}
