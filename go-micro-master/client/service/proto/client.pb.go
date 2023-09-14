// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/go-micro/v2/client/proto/client.proto

package go_micro_client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Request struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d418333f021a3308, []int{0}
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

func (m *Request) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Request) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d418333f021a3308, []int{1}
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

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Message struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	ContentType          string   `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d418333f021a3308, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.client.Request")
	proto.RegisterType((*Response)(nil), "go.micro.client.Response")
	proto.RegisterType((*Message)(nil), "go.micro.client.Message")
}

func init() {
	proto.RegisterFile("github.com/micro/go-micro/v2/client/proto/client.proto", fileDescriptor_d418333f021a3308)
}

var fileDescriptor_d418333f021a3308 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xeb, 0xb6, 0xff, 0xa4, 0xff, 0xa3, 0x12, 0x92, 0xc5, 0x60, 0x32, 0xa0, 0x92, 0x29,
	0x0b, 0x0e, 0xa2, 0x33, 0x62, 0xc8, 0x8c, 0x84, 0x02, 0x62, 0x45, 0x89, 0x7b, 0x4a, 0x2d, 0x25,
	0xb6, 0x89, 0x9d, 0x4a, 0xf9, 0x90, 0x7c, 0x27, 0x24, 0x27, 0xa5, 0x08, 0xda, 0x85, 0xed, 0xde,
	0xfd, 0xac, 0x77, 0xe7, 0x77, 0xb0, 0xae, 0xa4, 0xdb, 0x76, 0x25, 0x17, 0xba, 0x49, 0x1b, 0x29,
	0x5a, 0x9d, 0x56, 0xfa, 0x66, 0x28, 0x44, 0x2d, 0x51, 0xb9, 0xd4, 0xb4, 0xda, 0xed, 0x05, 0xf7,
	0x82, 0x9e, 0x57, 0x9a, 0xfb, 0x37, 0x7c, 0x68, 0xc7, 0x3b, 0x08, 0x73, 0x7c, 0xef, 0xd0, 0x3a,
	0xca, 0x20, 0xb4, 0xd8, 0xee, 0xa4, 0x40, 0x46, 0x56, 0x24, 0xf9, 0x9f, 0xef, 0x25, 0x8d, 0x60,
	0x81, 0x6a, 0x63, 0xb4, 0x54, 0x8e, 0x4d, 0x3d, 0xfa, 0xd2, 0xf4, 0x1a, 0x96, 0x42, 0x2b, 0x87,
	0xca, 0xbd, 0xb9, 0xde, 0x20, 0x9b, 0x79, 0x7e, 0x36, 0xf6, 0x5e, 0x7a, 0x83, 0x94, 0xc2, 0xbc,
	0xd4, 0x9b, 0x9e, 0xcd, 0x57, 0x24, 0x59, 0xe6, 0xbe, 0x8e, 0xaf, 0x60, 0x91, 0xa3, 0x35, 0x5a,
	0xd9, 0x03, 0x27, 0xdf, 0xf8, 0x2b, 0x84, 0x8f, 0x68, 0x6d, 0x51, 0x21, 0xbd, 0x80, 0x7f, 0x4e,
	0x1b, 0x29, 0xc6, 0xad, 0x06, 0xf1, 0x6b, 0xee, 0xf4, 0xf4, 0xdc, 0xd9, 0xc1, 0xf7, 0xee, 0x83,
	0x40, 0x90, 0xf9, 0xaf, 0xd3, 0x7b, 0x98, 0x67, 0x45, 0x5d, 0x53, 0xc6, 0x7f, 0x84, 0xc2, 0xc7,
	0x44, 0xa2, 0xcb, 0x23, 0x64, 0xd8, 0x39, 0x9e, 0xd0, 0x0c, 0x82, 0x67, 0xd7, 0x62, 0xd1, 0xfc,
	0xd1, 0x20, 0x21, 0xb7, 0x84, 0x3e, 0x40, 0xf8, 0xd4, 0x95, 0xb5, 0xb4, 0xdb, 0x23, 0x2e, 0x63,
	0x00, 0xd1, 0x49, 0x12, 0x4f, 0xca, 0xc0, 0xdf, 0x75, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0xb6,
	0x4d, 0x6e, 0xd5, 0x0e, 0x02, 0x00, 0x00,
}
