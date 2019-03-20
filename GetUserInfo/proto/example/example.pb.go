// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetUserInfo

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
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

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Sessionid            string   `protobuf:"bytes,1,opt,name=Sessionid,proto3" json:"Sessionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
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

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

type Response struct {
	//错误码
	Errno string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	//用户id
	UserId int64 `protobuf:"varint,3,opt,name=User_id,json=UserId,proto3" json:"User_id,omitempty"`
	//用户名
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	//秘密
	Password string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	//手机号
	Mobile string `protobuf:"bytes,6,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	//真实姓名
	RealName string `protobuf:"bytes,7,opt,name=Real_name,json=RealName,proto3" json:"Real_name,omitempty"`
	//身份证号
	IdCard string `protobuf:"bytes,8,opt,name=Id_card,json=IdCard,proto3" json:"Id_card,omitempty"`
	//头衔地址
	AvatarUrl            string   `protobuf:"bytes,9,opt,name=Avatar_url,json=AvatarUrl,proto3" json:"Avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
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

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Response) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Response) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *Response) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *Response) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.GetUserInfo.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetUserInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetUserInfo.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.GetUserInfo.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.GetUserInfo.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.GetUserInfo.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.GetUserInfo.Pong")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x5e, 0xd6, 0x36, 0x69, 0x8f, 0x97, 0x71, 0x9a, 0x98, 0xd5, 0xc2, 0x34, 0xfc, 0x42, 0x01,
	0x29, 0x4c, 0xf0, 0x0b, 0x10, 0xaa, 0x50, 0x1f, 0x86, 0xaa, 0x4c, 0xe3, 0xb5, 0xf2, 0xea, 0x23,
	0x8a, 0x48, 0xec, 0x72, 0x4e, 0x07, 0xfc, 0x15, 0x7e, 0x24, 0xbf, 0x01, 0xd9, 0x4e, 0xd9, 0x84,
	0xc8, 0xb4, 0xa7, 0xf8, 0xfb, 0xee, 0x3b, 0xdf, 0x7d, 0x77, 0x31, 0xcc, 0xb6, 0x6c, 0x5b, 0xfb,
	0x86, 0x7e, 0xa8, 0x66, 0x5b, 0xd3, 0xfe, 0x9b, 0x07, 0x16, 0x45, 0x69, 0xf3, 0xa6, 0xda, 0xb0,
	0xcd, 0x1d, 0xdf, 0xe4, 0x1f, 0xa9, 0xbd, 0x72, 0xc4, 0x4b, 0xf3, 0xc5, 0xca, 0x19, 0x64, 0x17,
	0xe4, 0x9c, 0x2a, 0x09, 0x8f, 0x60, 0xe0, 0xd4, 0x4f, 0x91, 0x9c, 0x25, 0xf3, 0x49, 0xe1, 0x8f,
	0xf2, 0x05, 0x64, 0x05, 0x7d, 0xdb, 0x91, 0x6b, 0xf1, 0x29, 0x4c, 0x2e, 0xc9, 0xb9, 0xca, 0x9a,
	0x4a, 0x77, 0x92, 0x5b, 0x42, 0xfe, 0x4e, 0x60, 0x5c, 0x90, 0xdb, 0x5a, 0xe3, 0x08, 0x8f, 0x61,
	0xb4, 0x60, 0x36, 0xb6, 0x93, 0x45, 0x80, 0x4f, 0x20, 0x5d, 0x30, 0x37, 0xae, 0x14, 0x87, 0x81,
	0xee, 0x10, 0x9e, 0x40, 0xe6, 0x9b, 0x59, 0x57, 0x5a, 0x0c, 0xce, 0x92, 0xf9, 0xa0, 0x48, 0x43,
	0x6f, 0x1a, 0x11, 0x86, 0x9f, 0x54, 0x43, 0x62, 0x18, 0xe4, 0xe1, 0x8c, 0x53, 0x18, 0xaf, 0x94,
	0x73, 0xdf, 0x2d, 0x6b, 0x31, 0x0a, 0xfc, 0x5f, 0xec, 0x0b, 0x5c, 0xd8, 0xeb, 0xaa, 0x26, 0x91,
	0xc6, 0x02, 0x11, 0xe1, 0x0c, 0x26, 0x05, 0xa9, 0x7a, 0x6d, 0xfc, 0x65, 0x59, 0x4c, 0xf2, 0x44,
	0xb8, 0xf0, 0x04, 0xb2, 0xa5, 0x5e, 0x6f, 0x14, 0x6b, 0x31, 0x8e, 0x59, 0x4b, 0xfd, 0x41, 0xb1,
	0xc6, 0x67, 0x00, 0xef, 0x6f, 0x54, 0xab, 0x78, 0xbd, 0xe3, 0x5a, 0x4c, 0xa2, 0xe1, 0xc8, 0x5c,
	0x71, 0x2d, 0xe7, 0x70, 0x74, 0xd9, 0x32, 0xa9, 0xa6, 0x32, 0xe5, 0x7e, 0x44, 0xc7, 0x30, 0xda,
	0xd8, 0x9d, 0x69, 0x83, 0xef, 0x41, 0x11, 0x81, 0x7c, 0x09, 0x8f, 0xef, 0x28, 0x6f, 0x47, 0xf4,
	0x1f, 0xe9, 0x29, 0x0c, 0x57, 0x95, 0x29, 0xbd, 0x13, 0xd7, 0xb2, 0xfd, 0x4a, 0x5d, 0xb8, 0x43,
	0x21, 0x6e, 0xfb, 0xe3, 0x6f, 0x7f, 0x1d, 0x42, 0xb6, 0x88, 0x7b, 0xc7, 0xcf, 0xf0, 0xe8, 0xce,
	0x9a, 0xf1, 0x79, 0xde, 0xf7, 0x07, 0xe4, 0x5d, 0xfb, 0x53, 0x79, 0x9f, 0x24, 0xf6, 0x2d, 0x0f,
	0x90, 0x20, 0x8d, 0x76, 0xf0, 0x55, 0xbf, 0xfe, 0xdf, 0xd1, 0x4c, 0x5f, 0x3f, 0x48, 0xbb, 0x2f,
	0x72, 0x9e, 0xe0, 0x0a, 0xc6, 0x7e, 0x14, 0xc1, 0xee, 0x69, 0x7f, 0xb2, 0xd7, 0x4c, 0xef, 0x8b,
	0x5b, 0x53, 0xca, 0x83, 0x79, 0x72, 0x9e, 0x5c, 0xa7, 0xe1, 0x25, 0xbc, 0xfb, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0x57, 0xbc, 0x26, 0x28, 0x03, 0x00, 0x00,
}
