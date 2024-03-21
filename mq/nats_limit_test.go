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

func Test_jet_limit(t *testing.T) {
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
		Name:        "limit",
		Description: "limit",
		Subjects:    []string{"limit", "limit.>"},
		Retention:   jetstream.LimitsPolicy,
		Compression: jetstream.S2Compression,
		MaxMsgs:     100,
	})
	if err != nil {
		t.Fatal(err)
	}
	///
	for i := 0; i < 10; i++ {
		j := i * 10
		jets.PublishAsync("limit.eth", []byte("eth:"+strconv.Itoa(j)))
		jets.PublishAsync("limit.rpg", []byte("rpg:"+strconv.Itoa(j)))
		jets.PublishAsync("limit.bsc", []byte("bsc:"+strconv.Itoa(j)))
		jets.PublishAsync("limit.none", []byte("none:"+strconv.Itoa(j)))
		jets.PublishAsync("limit", []byte("blank:"+strconv.Itoa(j)))
	}
	///
	cons_eth, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverLastPolicy,
		Durable:       "limit-eth",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "limit.eth",
	})
	if err != nil {
		t.Fatal(err)
	}
	///
	// _, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
	// 	DeliverPolicy: jetstream.DeliverLastPolicy,
	// 	Durable:       "limit-eth2",
	// 	AckPolicy:     jetstream.AckAllPolicy,
	// 	FilterSubject: "limit.eth",
	// })
	// if err == nil {
	// 	t.Fatal("should err: filtered consumer not unique on workqueue stream")
	// }
	////
	////
	cons_rpg, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverLastPolicy,
		Durable:       "limit-rpg",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "limit.rpg",
	})
	if err != nil {
		t.Fatal(err)
	}
	////
	cons_bsc, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverLastPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "limit.bsc",
	})
	if err != nil {
		t.Fatal(err)
	}
	//
	cons_blank, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverLastPolicy,
		Durable:       "limit",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "limit",
	})
	if err != nil {
		t.Fatal(err)
	}
	//
	cons_blank2, err := stream.Consumer(ctx, "limit")
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
