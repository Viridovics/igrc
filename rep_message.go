package igrc

import "sync"

type RepMsg[T any, R any] struct {
	body    T
	repChan chan R
	closed  bool
	lock    *sync.Mutex
}

func MakeRepMsg[T any, R any](body T) RepMsg[T, R] {
	return RepMsg[T, R]{
		body:    body,
		repChan: make(chan R, 1),
		lock:    new(sync.Mutex),
	}
}

func (m *RepMsg[T, R]) Reply(resp R) {
	m.execWithClosing(func() {
		m.repChan <- resp
	})
}

func (m *RepMsg[T, R]) Close() {
	m.execWithClosing(func() {})
}

func (m *RepMsg[T, R]) execWithClosing(f func()) {
	if m.closed {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.closed {
		return
	}
	f()
	close(m.repChan)
	m.closed = true
}

func (m *RepMsg[T, R]) Response() <-chan R {
	return m.repChan
}

func (m *RepMsg[T, R]) Body() T {
	return m.body
}
