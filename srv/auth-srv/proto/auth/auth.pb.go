// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package auth

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "go.micro.service.auth.Error")
	proto.RegisterType((*Request)(nil), "go.micro.service.auth.Request")
	proto.RegisterType((*Response)(nil), "go.micro.service.auth.Response")
}

func init() {
	proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5)
}

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x4d, 0xda, 0x38, 0x8a, 0xc2, 0xd0, 0x4a, 0x28, 0xa2, 0x25, 0xa7, 0x9e, 0xb6,
	0x90, 0x7e, 0x82, 0xfa, 0x07, 0xf1, 0xa0, 0x87, 0xb5, 0x22, 0x7a, 0x8b, 0x9b, 0xc1, 0x84, 0xb6,
	0xd9, 0xba, 0xbb, 0xf1, 0x73, 0xf8, 0x91, 0x25, 0x93, 0x68, 0x2f, 0xed, 0xad, 0x97, 0xf0, 0x7e,
	0xbc, 0x99, 0x37, 0x79, 0xb0, 0x30, 0x58, 0x1b, 0xed, 0xf4, 0x24, 0xad, 0x5c, 0xce, 0x1f, 0xc1,
	0x8c, 0x83, 0x4f, 0x2d, 0x56, 0x85, 0x32, 0x5a, 0x58, 0x32, 0xdf, 0x85, 0x22, 0x51, 0x9b, 0xf1,
	0x14, 0x82, 0x3b, 0x63, 0xb4, 0x41, 0x04, 0x5f, 0xe9, 0x8c, 0x22, 0x6f, 0xe4, 0x8d, 0x03, 0xc9,
	0x1a, 0xcf, 0xa1, 0x9b, 0x91, 0x4b, 0x8b, 0x65, 0xd4, 0x19, 0x79, 0xe3, 0x23, 0xd9, 0x52, 0xfc,
	0x0c, 0x3d, 0x49, 0x5f, 0x15, 0x59, 0x57, 0x8f, 0x54, 0x96, 0xcc, 0x43, 0xc6, 0x8b, 0xbe, 0x6c,
	0x09, 0x87, 0x10, 0xd6, 0xea, 0x29, 0x5d, 0x51, 0xbb, 0xfc, 0xcf, 0xd8, 0x87, 0xc0, 0xe9, 0x05,
	0x95, 0xd1, 0x21, 0x1b, 0x0d, 0xc4, 0x25, 0x84, 0x92, 0xec, 0x5a, 0x97, 0x96, 0x30, 0x82, 0x9e,
	0xad, 0x94, 0x22, 0x6b, 0x39, 0x36, 0x94, 0x7f, 0x88, 0x09, 0x04, 0x54, 0xff, 0x2f, 0x87, 0x1e,
	0x27, 0x17, 0x62, 0x6b, 0x2d, 0xc1, 0x9d, 0x64, 0x33, 0xba, 0xfd, 0x5e, 0xf2, 0xd3, 0x01, 0x7f,
	0x56, 0xb9, 0x1c, 0xe7, 0x70, 0xf6, 0x98, 0x2e, 0x68, 0xc6, 0x07, 0xe6, 0xb5, 0x87, 0x97, 0x3b,
	0x62, 0xdb, 0xd6, 0xc3, 0xab, 0x9d, 0x7e, 0x53, 0x20, 0x3e, 0xc0, 0x57, 0xc0, 0x5b, 0x5a, 0xbe,
	0x58, 0x32, 0x7b, 0x0e, 0x7e, 0x83, 0xfe, 0x3d, 0xb9, 0x9b, 0x54, 0xe5, 0x94, 0xed, 0x37, 0xfa,
	0xfa, 0xf4, 0xfd, 0x44, 0x4c, 0x36, 0xcf, 0xe7, 0xa3, 0xcb, 0x7a, 0xfa, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xd1, 0xa1, 0x5c, 0x2f, 0x53, 0x02, 0x00, 0x00,
}
