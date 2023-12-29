package igrc

import "sync"

type Ack <-chan struct{}

type AckMsg[T any] struct {
	body    T
	ackChan chan struct{}
	closed  bool
	lock    *sync.Mutex
}

func MakeAckMsg[T any](body T) AckMsg[T] {
	return AckMsg[T]{
		body:    body,
		ackChan: make(chan struct{}),
		lock:    new(sync.Mutex),
	}
}

func (m *AckMsg[T]) Conf() {
	m.Close()
}

func (m *AckMsg[T]) Close() {
	if m.closed {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.closed {
		return
	}
	close(m.ackChan)
	m.closed = true
}

func (m *AckMsg[T]) Ack() Ack {
	return m.ackChan
}

func (m *AckMsg[T]) Body() T {
	return m.body
}
