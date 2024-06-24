package mq

import (
	"fmt"
	"testing"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

func Test_StreamConsumer(t *testing.T) {
	url := "localhost:4222"
	natsServer := New(url)

	/////
	_, err := natsServer.CreateOrUpdateStream(JetStream_SyncChain, []string{JetSub_SyncChainTransfer, JetSub_SyncChainMint})
	if err != nil {
		t.Fatal(err)
	}
	////
	natsServer.JetStream().PublishAsync("testStream.a", []byte("testMsg"))
	/////
	consumer, err := natsServer.GetConsumer("testConsumer", "testStream", "testStream.a")
	if err != nil {
		t.Fatal(err)
	}
	consumer.Consume(func(msg jetstream.Msg) {
		t.Log(msg.Data())
		msg.Ack()
	})
	/////
	cons2, err := natsServer.GetConsumer("testConsumer2", "testStream", "testStream.a")
	if err != nil {
		t.Fatal(err)
	}
	cons2.Consume(func(msg jetstream.Msg) {
		t.Log(msg.Data())
		msg.Ack()
	})
	time.Sleep(3 * time.Second)
}
func Test_StreamConsumerMulSubject(t *testing.T) {
	url := "localhost:4222"
	natsServer := New(url)

	/////
	_, err := natsServer.CreateOrUpdateStream("testStream", []string{"testStream.>"})
	if err != nil {
		t.Fatal(err)
	}
	////
	natsServer.JetStream().PublishAsync("testStream.a", []byte("testMsg"))
	natsServer.JetStream().PublishAsync("testStream.a.offset", []byte("testMsgOffset"))
	/////
	consumer, err := natsServer.GetConsumer("testConsumer", "testStream", "testStream.a")
	if err != nil {
		t.Fatal(err)
	}
	consumer.Consume(func(msg jetstream.Msg) {
		t.Log(msg.Data())
		fmt.Println("cons1:", string(msg.Data()))
		msg.Ack()
	})
	/////
	cons2, err := natsServer.GetConsumer("testConsumer2", "testStream", "testStream.a.offset")
	if err != nil {
		t.Fatal(err)
	}
	cons2.Consume(func(msg jetstream.Msg) {
		t.Log(msg.Data())
		fmt.Println("cons2:", string(msg.Data()))
		msg.Ack()
	})
	time.Sleep(3 * time.Second)
}
