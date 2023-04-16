package exmqq

import (
	"testing"
	"time"
)

func TestPub(t *testing.T) {
	client := CreateMqttClient()
	go Subscribe(client)
	time.Sleep(time.Second * 4)
	Pubilsh(client)

}
