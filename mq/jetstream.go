package mq

import (
	"github.com/nats-io/nats.go/jetstream"
)

// const JetSub_SyncChain = "syncChain"
const JetSub_SyncChainTransfer = "syncChain.transfer"
const JetSub_SyncChainMint = "syncChain.mint"

const JetStream_SyncChain = "syncChain_stream"

func (s *NatsServer) JetStream() jetstream.JetStream {
	return s.jets
}
func (s *NatsServer) CreateOrUpdateStream(name string, subs []string, args ...int64) (jetstream.Stream, error) {
	jets, err := jetstream.New(s.nc)
	if err != nil {
		return nil, err
	}
	///
	msgSize := int64(0)
	if len(args) > 0 {
		msgSize = args[0]
	}
	stream, err := jets.CreateOrUpdateStream(s.ctx, jetstream.StreamConfig{
		Name:        name,
		Description: name,
		Subjects:    subs,
		Retention:   jetstream.LimitsPolicy,
		Compression: jetstream.S2Compression,
		MaxMsgs:     msgSize,
	})
	////
	return stream, err
}

// func (s *NatsServer) CreateOrUpdateMintStream(args ...int64) (jetstream.Stream, error) {
// 	jets, err := jetstream.New(s.nc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	///
// 	msgSize := int64(0)
// 	if len(args) > 0 {
// 		msgSize = args[0]
// 	}
// 	stream, err := jets.CreateOrUpdateStream(s.ctx, jetstream.StreamConfig{
// 		Name:        JetStream_SyncChain,
// 		Description: JetStream_SyncChain,
// 		Subjects:    []string{JetSub_ChainMint},
// 		Retention:   jetstream.LimitsPolicy,
// 		Compression: jetstream.S2Compression,
// 		MaxMsgs:     msgSize,
// 	})
// 	////
// 	return stream, err
// }

//	func (s *NatsServer) GetChainTransferStream() (jetstream.Stream, error) {
//		///
//		jets, err := jetstream.New(s.nc)
//		if err != nil {
//			return nil, err
//		}
//		//
//		stream, err := jets.Stream(s.ctx, JetStream_SyncChain)
//		////
//		return stream, err
//	}
//
// //consumer
func (s *NatsServer) GetConsumer(name, stream string, sub string) (jetstream.Consumer, error) {

	cons, err := s.jets.CreateOrUpdateConsumer(s.ctx, stream, jetstream.ConsumerConfig{
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

// func (s *NatsServer) GetChainMintConsumer(name, sub string) (jetstream.Consumer, error) {
// 	cons, err := s.jets.CreateOrUpdateConsumer(s.ctx, JetStream_ChainMint, jetstream.ConsumerConfig{
// 		Durable:       name,
// 		FilterSubject: sub,
// 		DeliverPolicy: jetstream.DeliverAllPolicy,
// 		AckPolicy:     jetstream.AckExplicitPolicy,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	//
// 	return cons, err
// }
