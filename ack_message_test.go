package igrc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAckMsg(t *testing.T) {
	msgProcessed := false
	message := MakeAckMsg(100)

	messageChan := make(chan AckMsg[int])

	go func() {
		msg := <-messageChan
		assert.Equal(t, 100, msg.Body())
		time.Sleep(1 * time.Second)
		msgProcessed = true
		msg.Ack()
	}()
	messageChan <- message
	<-message.C()
	assert.True(t, msgProcessed)
}

func TestAckMsgDoubleAckCall(t *testing.T) {
	message := MakeAckMsg(100)
	message.Ack()
	message.Ack()
}

func TestAckMsgDoubleCloseCall(t *testing.T) {
	message := MakeAckMsg(100)
	message.Close()
	message.Close()
}
