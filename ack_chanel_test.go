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
		time.Sleep(1 * time.Second)
		msgProcessed = true
		msg.Conf()
	}()
	messageChan <- message
	<-message.Ack()
	assert.True(t, msgProcessed)
}
