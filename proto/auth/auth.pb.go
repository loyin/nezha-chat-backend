// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package auth

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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
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

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetTokenReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenReq) Reset()         { *m = GetTokenReq{} }
func (m *GetTokenReq) String() string { return proto.CompactTextString(m) }
func (*GetTokenReq) ProtoMessage()    {}
func (*GetTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *GetTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenReq.Unmarshal(m, b)
}
func (m *GetTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenReq.Marshal(b, m, deterministic)
}
func (m *GetTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenReq.Merge(m, src)
}
func (m *GetTokenReq) XXX_Size() int {
	return xxx_messageInfo_GetTokenReq.Size(m)
}
func (m *GetTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenReq proto.InternalMessageInfo

func (m *GetTokenReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTokenResp struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenResp) Reset()         { *m = GetTokenResp{} }
func (m *GetTokenResp) String() string { return proto.CompactTextString(m) }
func (*GetTokenResp) ProtoMessage()    {}
func (*GetTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *GetTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenResp.Unmarshal(m, b)
}
func (m *GetTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenResp.Marshal(b, m, deterministic)
}
func (m *GetTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenResp.Merge(m, src)
}
func (m *GetTokenResp) XXX_Size() int {
	return xxx_messageInfo_GetTokenResp.Size(m)
}
func (m *GetTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenResp proto.InternalMessageInfo

func (m *GetTokenResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetTokenResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckReq) Reset()         { *m = CheckReq{} }
func (m *CheckReq) String() string { return proto.CompactTextString(m) }
func (*CheckReq) ProtoMessage()    {}
func (*CheckReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *CheckReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckReq.Unmarshal(m, b)
}
func (m *CheckReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckReq.Marshal(b, m, deterministic)
}
func (m *CheckReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckReq.Merge(m, src)
}
func (m *CheckReq) XXX_Size() int {
	return xxx_messageInfo_CheckReq.Size(m)
}
func (m *CheckReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckReq proto.InternalMessageInfo

func (m *CheckReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckResp struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResp) Reset()         { *m = CheckResp{} }
func (m *CheckResp) String() string { return proto.CompactTextString(m) }
func (*CheckResp) ProtoMessage()    {}
func (*CheckResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{4}
}

func (m *CheckResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResp.Unmarshal(m, b)
}
func (m *CheckResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResp.Marshal(b, m, deterministic)
}
func (m *CheckResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResp.Merge(m, src)
}
func (m *CheckResp) XXX_Size() int {
	return xxx_messageInfo_CheckResp.Size(m)
}
func (m *CheckResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResp proto.InternalMessageInfo

func (m *CheckResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CheckResp) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CheckResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "nezha.chat.auth.srv.service.Error")
	proto.RegisterType((*GetTokenReq)(nil), "nezha.chat.auth.srv.service.GetTokenReq")
	proto.RegisterType((*GetTokenResp)(nil), "nezha.chat.auth.srv.service.GetTokenResp")
	proto.RegisterType((*CheckReq)(nil), "nezha.chat.auth.srv.service.CheckReq")
	proto.RegisterType((*CheckResp)(nil), "nezha.chat.auth.srv.service.CheckResp")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x9b, 0x94, 0x64, 0x2a, 0x22, 0x83, 0x42, 0x89, 0x97, 0xb0, 0xa0, 0xc4, 0x83,
	0x2b, 0xd4, 0x8b, 0x1e, 0x55, 0xc4, 0x7b, 0xf0, 0x20, 0x1e, 0x84, 0x98, 0x0c, 0x26, 0x94, 0x66,
	0xd3, 0xdd, 0x4d, 0x0f, 0xbe, 0x9e, 0x2f, 0x26, 0xbb, 0x49, 0x6a, 0x0e, 0x12, 0x0a, 0x5e, 0x96,
	0x99, 0xdd, 0xff, 0x9b, 0x99, 0x7f, 0x58, 0x38, 0xad, 0xa5, 0xd0, 0xe2, 0x3a, 0x6d, 0x74, 0x61,
	0x0f, 0x6e, 0x73, 0x3c, 0xab, 0xe8, 0xab, 0x48, 0x79, 0x56, 0xa4, 0x9a, 0xdb, 0x6b, 0x25, 0xb7,
	0x5c, 0x91, 0xdc, 0x96, 0x19, 0xb1, 0x2b, 0xf0, 0x9e, 0xa4, 0x14, 0x12, 0x11, 0xdc, 0x4c, 0xe4,
	0xb4, 0x70, 0x22, 0x27, 0xf6, 0x12, 0x1b, 0xe3, 0x31, 0x4c, 0xd7, 0xea, 0x73, 0x31, 0x89, 0x9c,
	0x38, 0x48, 0x4c, 0xc8, 0xee, 0x60, 0xfe, 0x4c, 0xfa, 0x45, 0xac, 0xa8, 0x4a, 0x68, 0x83, 0x21,
	0xf8, 0x8d, 0x22, 0x59, 0xa5, 0xeb, 0x16, 0x0c, 0x92, 0x5d, 0x8e, 0x47, 0x30, 0x29, 0xf3, 0x8e,
	0x9d, 0x94, 0x39, 0x7b, 0x87, 0xc3, 0x5f, 0x54, 0xd5, 0x78, 0x0b, 0x1e, 0x99, 0xce, 0x16, 0x9c,
	0x2f, 0x19, 0x1f, 0x19, 0x93, 0xdb, 0x19, 0x93, 0x16, 0xc0, 0x13, 0xf0, 0xb4, 0x29, 0xd3, 0x15,
	0x6f, 0x13, 0x16, 0x81, 0xff, 0x58, 0x50, 0xb6, 0x32, 0x73, 0xed, 0x14, 0xce, 0x50, 0xb1, 0x81,
	0xa0, 0x53, 0xfc, 0xab, 0xfd, 0xd0, 0xf4, 0xf4, 0x4f, 0xd3, 0x6e, 0x6f, 0x7a, 0xf9, 0xed, 0x80,
	0x7b, 0xdf, 0xe8, 0x02, 0x53, 0xf0, 0x7b, 0xf7, 0x18, 0x8f, 0xf6, 0x1a, 0xec, 0x37, 0xbc, 0xdc,
	0x53, 0xa9, 0x6a, 0x76, 0x80, 0xaf, 0xe0, 0x59, 0x7b, 0x78, 0x3e, 0x4a, 0xf5, 0x4b, 0x0a, 0x2f,
	0xf6, 0x91, 0x99, 0xca, 0x0f, 0xb3, 0x37, 0xd7, 0xbc, 0x7f, 0xcc, 0xec, 0x87, 0xba, 0xf9, 0x09,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0x0c, 0x6a, 0x10, 0x69, 0x02, 0x00, 0x00,
}