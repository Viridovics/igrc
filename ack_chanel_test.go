package igrc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfAckMsg(t *testing.T) {
	msgProcessed := false
	message := MakeAckMsg(100)

	messageChan := make(chan AckMsg[int])

	go func() {
		msg := <-messageChan
		time.Sleep(1 * time.Second)
		msgProcessed = true
		msg.Conf()
	}()
	messageChan <- message
	<-message.Ack()
	assert.True(t, msgProcessed)
}

func TestDoubleConfCall(t *testing.T) {
	message := MakeAckMsg(100)
	message.Conf()
	message.Conf()
}

func TestDoubleCloseCall(t *testing.T) {
	message := MakeAckMsg(100)
	message.Close()
	message.Close()
}
