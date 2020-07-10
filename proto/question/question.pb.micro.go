// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/question/question.proto

package go_micro_service_question

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Question service

type QuestionService interface {
	CreateQuestion(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type questionService struct {
	c    client.Client
	name string
}

func NewQuestionService(name string, c client.Client) QuestionService {
	return &questionService{
		c:    c,
		name: name,
	}
}

func (c *questionService) CreateQuestion(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Question.CreateQuestion", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Question service

type QuestionHandler interface {
	CreateQuestion(context.Context, *Request, *Response) error
}

func RegisterQuestionHandler(s server.Server, hdlr QuestionHandler, opts ...server.HandlerOption) error {
	type question interface {
		CreateQuestion(ctx context.Context, in *Request, out *Response) error
	}
	type Question struct {
		question
	}
	h := &questionHandler{hdlr}
	return s.Handle(s.NewHandler(&Question{h}, opts...))
}

type questionHandler struct {
	QuestionHandler
}

func (h *questionHandler) CreateQuestion(ctx context.Context, in *Request, out *Response) error {
	return h.QuestionHandler.CreateQuestion(ctx, in, out)
}