package igrc

type Ack <-chan struct{}

type AckMsg[T any] struct {
	body    T
	ackChan chan struct{}
}

func MakeAckMsg[T any](body T) AckMsg[T] {
	return AckMsg[T]{
		body:    body,
		ackChan: make(chan struct{}),
	}
}

func (m *AckMsg[T]) Conf() {
	close(m.ackChan)
}

func (m *AckMsg[T]) Close() {
	close(m.ackChan)
}

func (m *AckMsg[T]) Ack() Ack {
	return m.ackChan
}

func (m *AckMsg[T]) Body() T {
	return m.body
}
