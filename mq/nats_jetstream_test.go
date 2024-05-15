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

func Test_JetBuild(t *testing.T) {
	url := "localhost:4222"
	nc, _ := nats.Connect(url)
	defer nc.Drain()

	jets, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	///
	///
	stream, err := jets.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:        "interest",
		Description: "interest",
		Subjects:    []string{"interesta", "interesta.>"},
		Retention:   jetstream.InterestPolicy,
		Compression: jetstream.S2Compression,
	})
	if err != nil {
		t.Fatal(err)
	}
	///
	for i := 0; i < 10; i++ {
		jets.PublishAsync("interest.none", []byte("none:"+strconv.Itoa(i)))
		jets.PublishAsync("interest.eth", []byte("eth:"+strconv.Itoa(i)))
	}
	///
	cons_eth, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverLastPolicy,
		Durable:       "interest-eth",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "interest.eth",
	})
	if err != nil {
		t.Fatal(err)
	}
	///
	//
	cons_eth.Consume(func(msg jetstream.Msg) {
		fmt.Println("eth:", string(msg.Data()))
		msg.Ack()
	})

	time.Sleep(3 * time.Second)
}
func Test_StreamBuild(t *testing.T) {
	url := "localhost:4222"
	nc, _ := nats.Connect(url)
	defer nc.Drain()

	jets, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	///
	///
	// _, err := jets.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
	// 	Name:        "interest",
	// 	Description: "interest",
	// 	Subjects:    []string{"interest", "interest.>"},
	// 	Retention:   jetstream.InterestPolicy,
	// 	Compression: jetstream.S2Compression,
	// })
	// if err != nil {
	// 	t.Fatal(err)
	// }
	///
	for i := 0; i < 10; i++ {
		jets.PublishAsync("interest.none", []byte("none:"+strconv.Itoa(i)))
		jets.PublishAsync("interest.eth", []byte("eth:"+strconv.Itoa(i)))
		// jets.PublishAsync("interest.rpg", []byte("rpg:"+strconv.Itoa(i)))
		// jets.PublishAsync("interest.bsc", []byte("bsc:"+strconv.Itoa(i)))
		// jets.PublishAsync("interest", []byte("blank:"+strconv.Itoa(i)))
	}
	///
	cons_eth, _ := jets.CreateOrUpdateConsumer(ctx, "interest", jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverLastPolicy,
		Durable:       "interest-eth2",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "interest.eth",
	})
	// cons_eth, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
	// 	DeliverPolicy: jetstream.DeliverLastPolicy,
	// 	Durable:       "interest-eth",
	// 	AckPolicy:     jetstream.AckExplicitPolicy,
	// 	FilterSubject: "interest.eth",
	// })
	// if err != nil {
	// 	t.Fatal(err)
	// }
	///
	//
	cons_eth.Consume(func(msg jetstream.Msg) {
		fmt.Println("eth:", string(msg.Data()))
		msg.Ack()
	})

	time.Sleep(3 * time.Second)
}
