package mq

import (
	"github.com/nats-io/nats.go/jetstream"
)

const SubJet_ChainTx = "chainTx"
const JetStream_ChainTx = "chainTx"

func (s *NatsServer) JetStream() (jetstream.JetStream, error) {
	jets, err := jetstream.New(s.nc)
	if err != nil {
		return nil, err
	}
	return jets, nil
}

func (s *NatsServer) GetChainTxStream() (jetstream.Stream, error) {
	///
	jets, err := jetstream.New(s.nc)
	if err != nil {
		return nil, err
	}
	//
	stream, err := jets.CreateOrUpdateStream(s.ctx, jetstream.StreamConfig{
		Name:        JetStream_ChainTx,
		Description: JetStream_ChainTx,
		Subjects:    []string{"chainTx", "chainTx.>"},
		Retention:   jetstream.LimitsPolicy,
		Compression: jetstream.S2Compression,
		MaxMsgs:     10000,
	})
	////
	return stream, err
}

func (s *NatsServer) GetConsumer(name, sub string) (jetstream.Consumer, error) {
	stream, err := s.GetChainTxStream()
	if err != nil {
		return nil, err
	}
	//
	cons, err := stream.CreateOrUpdateConsumer(s.ctx, jetstream.ConsumerConfig{
		Durable:       name,
		FilterSubject: sub,
		DeliverPolicy: jetstream.DeliverLastPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
	})
	if err != nil {
		return nil, err
	}
	//
	return cons, err
}
