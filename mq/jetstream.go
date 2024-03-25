package mq

import (
	"github.com/nats-io/nats.go/jetstream"
)

const JetSub_ChainTx = "chainData.tx"
const JetStream_ChainTx = "chainData_stream"

func (s *NatsServer) JetStream(msgSize int64) (jetstream.JetStream, error) {
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
		Subjects:    []string{"chainData", JetSub_ChainTx},
		Retention:   jetstream.LimitsPolicy,
		Compression: jetstream.S2Compression,
		MaxMsgs:     s.msgSize,
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
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
	})
	if err != nil {
		return nil, err
	}
	//
	return cons, err
}
