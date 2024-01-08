package igrc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepMsg(t *testing.T) {
	msgProcessed := false
	message := MakeRepMsg[int, float64](100)

	messageChan := make(chan RepMsg[int, float64])

	go func() {
		msg := <-messageChan
		assert.Equal(t, 100, msg.Body())
		time.Sleep(1 * time.Second)
		msgProcessed = true
		msg.Reply(1.0)
	}()
	messageChan <- message
	resp := <-message.WaitResp()
	assert.True(t, msgProcessed)
	assert.Equal(t, 1.0, resp)
}

func TestRepMsgDoubleReplyCall(t *testing.T) {
	message := MakeRepMsg[int, bool](100)
	message.Reply(true)
	resp, ok := <-message.WaitResp()
	assert.True(t, resp)
	assert.True(t, ok)
	message.Reply(true)
	_, ok = <-message.WaitResp()
	assert.False(t, ok)
}

func TestRepMsgDoubleCloseCall(t *testing.T) {
	message := MakeRepMsg[int, bool](100)
	message.Close()
	message.Close()
}

func TestRepMsgReplyAfterClose(t *testing.T) {
	message := MakeRepMsg[int, bool](100)
	message.Close()
	message.Reply(true)
	_, ok := <-message.WaitResp()
	assert.False(t, ok)
}
