package mq

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func Test_jet_stream_workqueue(t *testing.T) {
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		t.Fatal(err)
	}
	jets, err := jetstream.New(nc)
	if err != nil {
		t.Fatal(err)
	}
	///
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	///
	stream, err := jets.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:        "chainTx",
		Description: "chain txs",
		Subjects:    []string{"chainTx", "chainTx.>"},
		Retention:   jetstream.WorkQueuePolicy,
		Compression: jetstream.S2Compression,
	})
	if err != nil {
		t.Fatal(err)
	}
	///
	for i := 0; i < 10; i++ {
		jets.PublishAsync("chainTx.eth", []byte("eth:"+strconv.Itoa(i)))
		jets.PublishAsync("chainTx.rpg", []byte("rpg:"+strconv.Itoa(i)))
		jets.PublishAsync("chainTx.bsc", []byte("bsc:"+strconv.Itoa(i)))
		jets.PublishAsync("chainTx.none", []byte("none:"+strconv.Itoa(i)))
		jets.PublishAsync("chainTx", []byte("blank:"+strconv.Itoa(i)))
	}
	///
	cons_eth, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverAllPolicy,
		Durable:       "chainTx-eth",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "chainTx.eth",
	})
	if err != nil {
		t.Fatal(err)
	}
	///
	// _, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
	// 	DeliverPolicy: jetstream.DeliverAllPolicy,
	// 	Durable:       "chainTx-eth2",
	// 	AckPolicy:     jetstream.AckExplicitPolicy,
	// 	FilterSubject: "chainTx.eth",
	// })
	// if err == nil {
	// 	t.Fatal("should err: filtered consumer not unique on workqueue stream")
	// }
	////
	////
	cons_rpg, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverAllPolicy,
		Durable:       "chainTx-rpg",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "chainTx.rpg",
	})
	if err != nil {
		t.Fatal(err)
	}
	////
	cons_bsc, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "chainTx.bsc",
	})
	if err != nil {
		t.Fatal(err)
	}
	//
	cons_blank, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverAllPolicy,
		Durable:       "chainTx",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "chainTx",
	})
	if err != nil {
		t.Fatal(err)
	}
	//
	cons_blank2, err := stream.Consumer(ctx, "chainTx")
	if err != nil {
		t.Fatal(err)
	}
	//
	cons_eth.Consume(func(msg jetstream.Msg) {
		fmt.Println("eth:", string(msg.Data()))
		msg.Ack()
	})
	// cons_eth2.Consume(func(msg jetstream.Msg) {
	// 	fmt.Println("eth:", string(msg.Data()))
	// 	msg.Ack()
	// })
	cons_bsc.Consume(func(msg jetstream.Msg) {
		fmt.Println("bsc:", string(msg.Data()))
		msg.Ack()
	})
	cons_rpg.Consume(func(msg jetstream.Msg) {
		fmt.Println("rpg:", string(msg.Data()))
		msg.Ack()
	})
	cons_blank.Consume(func(msg jetstream.Msg) {
		fmt.Println("blank1:", string(msg.Data()))
		msg.Ack()
	})
	cons_blank2.Consume(func(msg jetstream.Msg) {
		fmt.Println("blank2:", string(msg.Data()))
		msg.Ack()
	})

	time.Sleep(3 * time.Second)
}
