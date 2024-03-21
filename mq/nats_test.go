package mq

import (
	"context"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

func Test_jet_queue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		t.Fatal(err)
	}
	js, err := nc.JetStream()
	if err != nil {
		t.Fatal(err)
	}
	///
	con, err := js.AddConsumer("s-workqueue", &nats.ConsumerConfig{
		DeliverPolicy: nats.DeliverAllPolicy,
		Durable:       "chainTx",
		AckPolicy:     nats.AckExplicitPolicy,
		FilterSubject: "sub.tx",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(con)
	///
	sub, err := js.PullSubscribe("sub.tx", "chainTx")
	if err != nil {
		t.Fatal(err)
	}
	for {
		msgs, err := sub.Fetch(1, nats.Context(ctx))
		if err != nil {
			t.Fatal(err)

		}
		for _, msg := range msgs {
			t.Log(string(msg.Data))
			msg.Respond([]byte("hello"))
			msg.Ack()
		}
	}
}
func Test_jet_queue2(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		t.Fatal(err)
	}
	js, err := nc.JetStream()
	if err != nil {
		t.Fatal(err)
	}
	///
	con, err := js.AddConsumer("s-workqueue", &nats.ConsumerConfig{
		DeliverPolicy: nats.DeliverAllPolicy,
		Durable:       "chainTx",
		AckPolicy:     nats.AckExplicitPolicy,
		FilterSubject: "sub.tx.eth",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(con)
	///
	sub, err := js.PullSubscribe("sub.tx", "chainTx")
	if err != nil {
		t.Fatal(err)
	}
	for {
		msgs, err := sub.Fetch(1, nats.Context(ctx))
		if err != nil {
			t.Fatal(err)

		}
		for _, msg := range msgs {
			t.Log(string(msg.Data))
			msg.Ack()
		}
	}
}
