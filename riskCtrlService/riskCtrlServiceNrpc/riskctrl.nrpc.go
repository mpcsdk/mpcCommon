// This code was autogenerated from riskctrl.proto, do not edit.
package riskCtrlServiceNrpc

import (
	"context"
	"log"
	"time"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
	"github.com/nats-io/nats.go"
	github_com_golang_protobuf_ptypes_empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/franklihub/nrpc"
)

// RiskCtrlServiceServer is the interface that providers of the service
// RiskCtrlService should implement.
type RiskCtrlServiceServer interface {
	Alive(ctx context.Context, req *github_com_golang_protobuf_ptypes_empty.Empty) (*github_com_golang_protobuf_ptypes_empty.Empty, error)
	TxsRequest(ctx context.Context, req *TxRequestReq) (*TxRequestRes, error)
	TfaRequest(ctx context.Context, req *TfaRequestReq) (*TfaRequestRes, error)
	TfaInfo(ctx context.Context, req *TfaInfoReq) (*TfaInfoRes, error)
	SendPhoneCode(ctx context.Context, req *SendPhoneCodeReq) (*SendPhoneCodeRes, error)
	SendMailCode(ctx context.Context, req *SendMailCodeReq) (*SendMailCodeRes, error)
	VerifyCode(ctx context.Context, req *VerifyCodeReq) (*VerifyCodeRes, error)
}

// RiskCtrlServiceHandler provides a NATS subscription handler that can serve a
// subscription using a given RiskCtrlServiceServer implementation.
type RiskCtrlServiceHandler struct {
	ctx     context.Context
	workers *nrpc.WorkerPool
	nc      nrpc.NatsConn
	server  RiskCtrlServiceServer

	encodings []string
}

func NewRiskCtrlServiceHandler(ctx context.Context, nc nrpc.NatsConn, s RiskCtrlServiceServer) *RiskCtrlServiceHandler {
	return &RiskCtrlServiceHandler{
		ctx:    ctx,
		nc:     nc,
		server: s,

		encodings: []string{"protobuf"},
	}
}

func NewRiskCtrlServiceConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s RiskCtrlServiceServer) *RiskCtrlServiceHandler {
	return &RiskCtrlServiceHandler{
		workers: workers,
		nc:      nc,
		server:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *RiskCtrlServiceHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}

func (h *RiskCtrlServiceHandler) Subject() string {
	return "RiskCtrlService.>"
}

func (h *RiskCtrlServiceHandler) Handler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	_, _, name, tail, err := nrpc.ParseSubject(
		"", 0, "RiskCtrlService", 0, msg.Subject)
	if err != nil {
		log.Printf("RiskCtrlServiceHanlder: RiskCtrlService subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
	case "Alive":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("AliveHanlder: Alive subject parsing failed: %v", err)
			break
		}
		var req github_com_golang_protobuf_ptypes_empty.Empty
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("AliveHandler: Alive request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {

			ctx := context.Background()
			ctx = context.WithValue(ctx, "tracingMiddlewareHandled", 1)
			var (
				span trace.Span
				tr   = otel.GetTracerProvider().Tracer(
					"nrpc-trace",
					trace.WithInstrumentationVersion("v0.0.1"),
				)
			)
			ctx, span = tr.Start(
				otel.GetTextMapPropagator().Extract(
					ctx,
					propagation.HeaderCarrier(msg.Header),
				),
				msg.Subject,
				trace.WithSpanKind(trace.SpanKindServer),
			)
			span.SetAttributes(nrpc.CommonLabels()...)
			// Inject tracing context.
			request.Context = ctx
			/////
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.Alive(ctx, &req)
				defer span.End()
				if err != nil {
					span.SetStatus(codes.Error, fmt.Sprintf("%+v", err))
					return nil, err
				}
				return innerResp, err
			}
		}
	case "TxsRequest":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("TxsRequestHanlder: TxsRequest subject parsing failed: %v", err)
			break
		}
		var req TxRequestReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("TxsRequestHandler: TxsRequest request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {

			ctx := context.Background()
			ctx = context.WithValue(ctx, "tracingMiddlewareHandled", 1)
			var (
				span trace.Span
				tr   = otel.GetTracerProvider().Tracer(
					"nrpc-trace",
					trace.WithInstrumentationVersion("v0.0.1"),
				)
			)
			ctx, span = tr.Start(
				otel.GetTextMapPropagator().Extract(
					ctx,
					propagation.HeaderCarrier(msg.Header),
				),
				msg.Subject,
				trace.WithSpanKind(trace.SpanKindServer),
			)
			span.SetAttributes(nrpc.CommonLabels()...)
			// Inject tracing context.
			request.Context = ctx
			/////
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.TxsRequest(ctx, &req)
				defer span.End()
				if err != nil {
					span.SetStatus(codes.Error, fmt.Sprintf("%+v", err))
					return nil, err
				}
				return innerResp, err
			}
		}
	case "TfaRequest":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("TfaRequestHanlder: TfaRequest subject parsing failed: %v", err)
			break
		}
		var req TfaRequestReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("TfaRequestHandler: TfaRequest request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {

			ctx := context.Background()
			ctx = context.WithValue(ctx, "tracingMiddlewareHandled", 1)
			var (
				span trace.Span
				tr   = otel.GetTracerProvider().Tracer(
					"nrpc-trace",
					trace.WithInstrumentationVersion("v0.0.1"),
				)
			)
			ctx, span = tr.Start(
				otel.GetTextMapPropagator().Extract(
					ctx,
					propagation.HeaderCarrier(msg.Header),
				),
				msg.Subject,
				trace.WithSpanKind(trace.SpanKindServer),
			)
			span.SetAttributes(nrpc.CommonLabels()...)
			// Inject tracing context.
			request.Context = ctx
			/////
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.TfaRequest(ctx, &req)
				defer span.End()
				if err != nil {
					span.SetStatus(codes.Error, fmt.Sprintf("%+v", err))
					return nil, err
				}
				return innerResp, err
			}
		}
	case "TfaInfo":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("TfaInfoHanlder: TfaInfo subject parsing failed: %v", err)
			break
		}
		var req TfaInfoReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("TfaInfoHandler: TfaInfo request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {

			ctx := context.Background()
			ctx = context.WithValue(ctx, "tracingMiddlewareHandled", 1)
			var (
				span trace.Span
				tr   = otel.GetTracerProvider().Tracer(
					"nrpc-trace",
					trace.WithInstrumentationVersion("v0.0.1"),
				)
			)
			ctx, span = tr.Start(
				otel.GetTextMapPropagator().Extract(
					ctx,
					propagation.HeaderCarrier(msg.Header),
				),
				msg.Subject,
				trace.WithSpanKind(trace.SpanKindServer),
			)
			span.SetAttributes(nrpc.CommonLabels()...)
			// Inject tracing context.
			request.Context = ctx
			/////
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.TfaInfo(ctx, &req)
				defer span.End()
				if err != nil {
					span.SetStatus(codes.Error, fmt.Sprintf("%+v", err))
					return nil, err
				}
				return innerResp, err
			}
		}
	case "SendPhoneCode":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("SendPhoneCodeHanlder: SendPhoneCode subject parsing failed: %v", err)
			break
		}
		var req SendPhoneCodeReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("SendPhoneCodeHandler: SendPhoneCode request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {

			ctx := context.Background()
			ctx = context.WithValue(ctx, "tracingMiddlewareHandled", 1)
			var (
				span trace.Span
				tr   = otel.GetTracerProvider().Tracer(
					"nrpc-trace",
					trace.WithInstrumentationVersion("v0.0.1"),
				)
			)
			ctx, span = tr.Start(
				otel.GetTextMapPropagator().Extract(
					ctx,
					propagation.HeaderCarrier(msg.Header),
				),
				msg.Subject,
				trace.WithSpanKind(trace.SpanKindServer),
			)
			span.SetAttributes(nrpc.CommonLabels()...)
			// Inject tracing context.
			request.Context = ctx
			/////
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.SendPhoneCode(ctx, &req)
				defer span.End()
				if err != nil {
					span.SetStatus(codes.Error, fmt.Sprintf("%+v", err))
					return nil, err
				}
				return innerResp, err
			}
		}
	case "SendMailCode":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("SendMailCodeHanlder: SendMailCode subject parsing failed: %v", err)
			break
		}
		var req SendMailCodeReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("SendMailCodeHandler: SendMailCode request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {

			ctx := context.Background()
			ctx = context.WithValue(ctx, "tracingMiddlewareHandled", 1)
			var (
				span trace.Span
				tr   = otel.GetTracerProvider().Tracer(
					"nrpc-trace",
					trace.WithInstrumentationVersion("v0.0.1"),
				)
			)
			ctx, span = tr.Start(
				otel.GetTextMapPropagator().Extract(
					ctx,
					propagation.HeaderCarrier(msg.Header),
				),
				msg.Subject,
				trace.WithSpanKind(trace.SpanKindServer),
			)
			span.SetAttributes(nrpc.CommonLabels()...)
			// Inject tracing context.
			request.Context = ctx
			/////
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.SendMailCode(ctx, &req)
				defer span.End()
				if err != nil {
					span.SetStatus(codes.Error, fmt.Sprintf("%+v", err))
					return nil, err
				}
				return innerResp, err
			}
		}
	case "VerifyCode":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("VerifyCodeHanlder: VerifyCode subject parsing failed: %v", err)
			break
		}
		var req VerifyCodeReq
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("VerifyCodeHandler: VerifyCode request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {

			ctx := context.Background()
			ctx = context.WithValue(ctx, "tracingMiddlewareHandled", 1)
			var (
				span trace.Span
				tr   = otel.GetTracerProvider().Tracer(
					"nrpc-trace",
					trace.WithInstrumentationVersion("v0.0.1"),
				)
			)
			ctx, span = tr.Start(
				otel.GetTextMapPropagator().Extract(
					ctx,
					propagation.HeaderCarrier(msg.Header),
				),
				msg.Subject,
				trace.WithSpanKind(trace.SpanKindServer),
			)
			span.SetAttributes(nrpc.CommonLabels()...)
			// Inject tracing context.
			request.Context = ctx
			/////
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.VerifyCode(ctx, &req)
				defer span.End()
				if err != nil {
					span.SetStatus(codes.Error, fmt.Sprintf("%+v", err))
					return nil, err
				}
				return innerResp, err
			}
		}
	default:
		log.Printf("RiskCtrlServiceHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type: nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}
	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("RiskCtrlServiceHandler: RiskCtrlService handler failed to publish the response: %s", err)
		}
	} else {
	}
}

type RiskCtrlServiceClient struct {
	nc      nrpc.NatsConn
	Subject string
	Encoding string
	Timeout time.Duration
}

func NewRiskCtrlServiceClient(nc nrpc.NatsConn) *RiskCtrlServiceClient {
	return &RiskCtrlServiceClient{
		nc:      nc,
		Subject: "RiskCtrlService",
		Encoding: "protobuf",
		Timeout: 5 * time.Second,
	}
}

func (c *RiskCtrlServiceClient) Alive(
	ctx context.Context,req *github_com_golang_protobuf_ptypes_empty.Empty) (*github_com_golang_protobuf_ptypes_empty.Empty, error) {

	subject := c.Subject + "." + "Alive"
	////
	// otel trace
	tr := otel.GetTracerProvider().Tracer(
		"nrpc-trace",
		trace.WithInstrumentationVersion("v0.0.1"),
	)
	ctx, span := tr.Start(ctx, subject, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	span.SetAttributes(nrpc.CommonLabels()...)
	// call msg
	rawRequest, _ := nrpc.Marshal(c.Encoding, req)
	reqMsg := nats.NewMsg(subject)
	reqMsg.Data = rawRequest
	// Inject tracing content into  header.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(reqMsg.Header))
	////
	// call
	var resp = github_com_golang_protobuf_ptypes_empty.Empty{}
	// if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
	if err := nrpc.CallMsg(ctx, reqMsg, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskCtrlServiceClient) TxsRequest(
	ctx context.Context,req *TxRequestReq) (*TxRequestRes, error) {

	subject := c.Subject + "." + "TxsRequest"
	////
	// otel trace
	tr := otel.GetTracerProvider().Tracer(
		"nrpc-trace",
		trace.WithInstrumentationVersion("v0.0.1"),
	)
	ctx, span := tr.Start(ctx, subject, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	span.SetAttributes(nrpc.CommonLabels()...)
	// call msg
	rawRequest, _ := nrpc.Marshal(c.Encoding, req)
	reqMsg := nats.NewMsg(subject)
	reqMsg.Data = rawRequest
	// Inject tracing content into  header.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(reqMsg.Header))
	////
	// call
	var resp = TxRequestRes{}
	// if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
	if err := nrpc.CallMsg(ctx, reqMsg, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskCtrlServiceClient) TfaRequest(
	ctx context.Context,req *TfaRequestReq) (*TfaRequestRes, error) {

	subject := c.Subject + "." + "TfaRequest"
	////
	// otel trace
	tr := otel.GetTracerProvider().Tracer(
		"nrpc-trace",
		trace.WithInstrumentationVersion("v0.0.1"),
	)
	ctx, span := tr.Start(ctx, subject, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	span.SetAttributes(nrpc.CommonLabels()...)
	// call msg
	rawRequest, _ := nrpc.Marshal(c.Encoding, req)
	reqMsg := nats.NewMsg(subject)
	reqMsg.Data = rawRequest
	// Inject tracing content into  header.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(reqMsg.Header))
	////
	// call
	var resp = TfaRequestRes{}
	// if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
	if err := nrpc.CallMsg(ctx, reqMsg, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskCtrlServiceClient) TfaInfo(
	ctx context.Context,req *TfaInfoReq) (*TfaInfoRes, error) {

	subject := c.Subject + "." + "TfaInfo"
	////
	// otel trace
	tr := otel.GetTracerProvider().Tracer(
		"nrpc-trace",
		trace.WithInstrumentationVersion("v0.0.1"),
	)
	ctx, span := tr.Start(ctx, subject, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	span.SetAttributes(nrpc.CommonLabels()...)
	// call msg
	rawRequest, _ := nrpc.Marshal(c.Encoding, req)
	reqMsg := nats.NewMsg(subject)
	reqMsg.Data = rawRequest
	// Inject tracing content into  header.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(reqMsg.Header))
	////
	// call
	var resp = TfaInfoRes{}
	// if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
	if err := nrpc.CallMsg(ctx, reqMsg, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskCtrlServiceClient) SendPhoneCode(
	ctx context.Context,req *SendPhoneCodeReq) (*SendPhoneCodeRes, error) {

	subject := c.Subject + "." + "SendPhoneCode"
	////
	// otel trace
	tr := otel.GetTracerProvider().Tracer(
		"nrpc-trace",
		trace.WithInstrumentationVersion("v0.0.1"),
	)
	ctx, span := tr.Start(ctx, subject, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	span.SetAttributes(nrpc.CommonLabels()...)
	// call msg
	rawRequest, _ := nrpc.Marshal(c.Encoding, req)
	reqMsg := nats.NewMsg(subject)
	reqMsg.Data = rawRequest
	// Inject tracing content into  header.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(reqMsg.Header))
	////
	// call
	var resp = SendPhoneCodeRes{}
	// if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
	if err := nrpc.CallMsg(ctx, reqMsg, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskCtrlServiceClient) SendMailCode(
	ctx context.Context,req *SendMailCodeReq) (*SendMailCodeRes, error) {

	subject := c.Subject + "." + "SendMailCode"
	////
	// otel trace
	tr := otel.GetTracerProvider().Tracer(
		"nrpc-trace",
		trace.WithInstrumentationVersion("v0.0.1"),
	)
	ctx, span := tr.Start(ctx, subject, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	span.SetAttributes(nrpc.CommonLabels()...)
	// call msg
	rawRequest, _ := nrpc.Marshal(c.Encoding, req)
	reqMsg := nats.NewMsg(subject)
	reqMsg.Data = rawRequest
	// Inject tracing content into  header.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(reqMsg.Header))
	////
	// call
	var resp = SendMailCodeRes{}
	// if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
	if err := nrpc.CallMsg(ctx, reqMsg, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *RiskCtrlServiceClient) VerifyCode(
	ctx context.Context,req *VerifyCodeReq) (*VerifyCodeRes, error) {

	subject := c.Subject + "." + "VerifyCode"
	////
	// otel trace
	tr := otel.GetTracerProvider().Tracer(
		"nrpc-trace",
		trace.WithInstrumentationVersion("v0.0.1"),
	)
	ctx, span := tr.Start(ctx, subject, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	span.SetAttributes(nrpc.CommonLabels()...)
	// call msg
	rawRequest, _ := nrpc.Marshal(c.Encoding, req)
	reqMsg := nats.NewMsg(subject)
	reqMsg.Data = rawRequest
	// Inject tracing content into  header.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(reqMsg.Header))
	////
	// call
	var resp = VerifyCodeRes{}
	// if err := nrpc.Call(req, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
	if err := nrpc.CallMsg(ctx, reqMsg, &resp, c.nc, subject, c.Encoding, c.Timeout); err != nil {
		return nil, err
	}

	return &resp, nil
}

type Client struct {
	nc      nrpc.NatsConn
	defaultEncoding string
	defaultTimeout time.Duration
	RiskCtrlService *RiskCtrlServiceClient
}

func NewClient(nc nrpc.NatsConn) *Client {
	c := Client{
		nc: nc,
		defaultEncoding: "protobuf",
		defaultTimeout: 5*time.Second,
	}
	c.RiskCtrlService = NewRiskCtrlServiceClient(nc)
	return &c
}

func (c *Client) SetEncoding(encoding string) {
	c.defaultEncoding = encoding
	if c.RiskCtrlService != nil {
		c.RiskCtrlService.Encoding = encoding
	}
}

func (c *Client) SetTimeout(t time.Duration) {
	c.defaultTimeout = t
	if c.RiskCtrlService != nil {
		c.RiskCtrlService.Timeout = t
	}
}