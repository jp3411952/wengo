// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account_proto

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

//注册用户客户端到登录服
type CL_LS_ReqRegisterAccoutMsg struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ClientType           uint32   `protobuf:"varint,3,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	PhoneNum             uint32   `protobuf:"varint,4,opt,name=phone_num,json=phoneNum,proto3" json:"phone_num,omitempty"`
	MacAddr              string   `protobuf:"bytes,5,opt,name=mac_addr,json=macAddr,proto3" json:"mac_addr,omitempty"`
	Version              uint32   `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CL_LS_ReqRegisterAccoutMsg) Reset()         { *m = CL_LS_ReqRegisterAccoutMsg{} }
func (m *CL_LS_ReqRegisterAccoutMsg) String() string { return proto.CompactTextString(m) }
func (*CL_LS_ReqRegisterAccoutMsg) ProtoMessage()    {}
func (*CL_LS_ReqRegisterAccoutMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *CL_LS_ReqRegisterAccoutMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CL_LS_ReqRegisterAccoutMsg.Unmarshal(m, b)
}
func (m *CL_LS_ReqRegisterAccoutMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CL_LS_ReqRegisterAccoutMsg.Marshal(b, m, deterministic)
}
func (m *CL_LS_ReqRegisterAccoutMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CL_LS_ReqRegisterAccoutMsg.Merge(m, src)
}
func (m *CL_LS_ReqRegisterAccoutMsg) XXX_Size() int {
	return xxx_messageInfo_CL_LS_ReqRegisterAccoutMsg.Size(m)
}
func (m *CL_LS_ReqRegisterAccoutMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CL_LS_ReqRegisterAccoutMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CL_LS_ReqRegisterAccoutMsg proto.InternalMessageInfo

func (m *CL_LS_ReqRegisterAccoutMsg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CL_LS_ReqRegisterAccoutMsg) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CL_LS_ReqRegisterAccoutMsg) GetClientType() uint32 {
	if m != nil {
		return m.ClientType
	}
	return 0
}

func (m *CL_LS_ReqRegisterAccoutMsg) GetPhoneNum() uint32 {
	if m != nil {
		return m.PhoneNum
	}
	return 0
}

func (m *CL_LS_ReqRegisterAccoutMsg) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

func (m *CL_LS_ReqRegisterAccoutMsg) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

//登陆
type CL_LS_ReqLoginMsg struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ClientType           uint32   `protobuf:"varint,3,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	MacAddr              string   `protobuf:"bytes,4,opt,name=mac_addr,json=macAddr,proto3" json:"mac_addr,omitempty"`
	Version              uint32   `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CL_LS_ReqLoginMsg) Reset()         { *m = CL_LS_ReqLoginMsg{} }
func (m *CL_LS_ReqLoginMsg) String() string { return proto.CompactTextString(m) }
func (*CL_LS_ReqLoginMsg) ProtoMessage()    {}
func (*CL_LS_ReqLoginMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *CL_LS_ReqLoginMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CL_LS_ReqLoginMsg.Unmarshal(m, b)
}
func (m *CL_LS_ReqLoginMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CL_LS_ReqLoginMsg.Marshal(b, m, deterministic)
}
func (m *CL_LS_ReqLoginMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CL_LS_ReqLoginMsg.Merge(m, src)
}
func (m *CL_LS_ReqLoginMsg) XXX_Size() int {
	return xxx_messageInfo_CL_LS_ReqLoginMsg.Size(m)
}
func (m *CL_LS_ReqLoginMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CL_LS_ReqLoginMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CL_LS_ReqLoginMsg proto.InternalMessageInfo

func (m *CL_LS_ReqLoginMsg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CL_LS_ReqLoginMsg) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CL_LS_ReqLoginMsg) GetClientType() uint32 {
	if m != nil {
		return m.ClientType
	}
	return 0
}

func (m *CL_LS_ReqLoginMsg) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

func (m *CL_LS_ReqLoginMsg) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

//登录服向客户端返回
type LS_CL_RespnLoginMsg struct {
	RestCode             int32    `protobuf:"varint,1,opt,name=rest_code,json=restCode,proto3" json:"rest_code,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	AccountID            uint64   `protobuf:"varint,3,opt,name=accountID,proto3" json:"accountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LS_CL_RespnLoginMsg) Reset()         { *m = LS_CL_RespnLoginMsg{} }
func (m *LS_CL_RespnLoginMsg) String() string { return proto.CompactTextString(m) }
func (*LS_CL_RespnLoginMsg) ProtoMessage()    {}
func (*LS_CL_RespnLoginMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *LS_CL_RespnLoginMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LS_CL_RespnLoginMsg.Unmarshal(m, b)
}
func (m *LS_CL_RespnLoginMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LS_CL_RespnLoginMsg.Marshal(b, m, deterministic)
}
func (m *LS_CL_RespnLoginMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LS_CL_RespnLoginMsg.Merge(m, src)
}
func (m *LS_CL_RespnLoginMsg) XXX_Size() int {
	return xxx_messageInfo_LS_CL_RespnLoginMsg.Size(m)
}
func (m *LS_CL_RespnLoginMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_LS_CL_RespnLoginMsg.DiscardUnknown(m)
}

var xxx_messageInfo_LS_CL_RespnLoginMsg proto.InternalMessageInfo

func (m *LS_CL_RespnLoginMsg) GetRestCode() int32 {
	if m != nil {
		return m.RestCode
	}
	return 0
}

func (m *LS_CL_RespnLoginMsg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LS_CL_RespnLoginMsg) GetAccountID() uint64 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

type CL_LS_ReqCreatePlayerMsg struct {
	Playername           string   `protobuf:"bytes,1,opt,name=playername,proto3" json:"playername,omitempty"`
	Sex                  int32    `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	Job                  int32    `protobuf:"varint,3,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CL_LS_ReqCreatePlayerMsg) Reset()         { *m = CL_LS_ReqCreatePlayerMsg{} }
func (m *CL_LS_ReqCreatePlayerMsg) String() string { return proto.CompactTextString(m) }
func (*CL_LS_ReqCreatePlayerMsg) ProtoMessage()    {}
func (*CL_LS_ReqCreatePlayerMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *CL_LS_ReqCreatePlayerMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CL_LS_ReqCreatePlayerMsg.Unmarshal(m, b)
}
func (m *CL_LS_ReqCreatePlayerMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CL_LS_ReqCreatePlayerMsg.Marshal(b, m, deterministic)
}
func (m *CL_LS_ReqCreatePlayerMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CL_LS_ReqCreatePlayerMsg.Merge(m, src)
}
func (m *CL_LS_ReqCreatePlayerMsg) XXX_Size() int {
	return xxx_messageInfo_CL_LS_ReqCreatePlayerMsg.Size(m)
}
func (m *CL_LS_ReqCreatePlayerMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CL_LS_ReqCreatePlayerMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CL_LS_ReqCreatePlayerMsg proto.InternalMessageInfo

func (m *CL_LS_ReqCreatePlayerMsg) GetPlayername() string {
	if m != nil {
		return m.Playername
	}
	return ""
}

func (m *CL_LS_ReqCreatePlayerMsg) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *CL_LS_ReqCreatePlayerMsg) GetJob() int32 {
	if m != nil {
		return m.Job
	}
	return 0
}

//客户端发送离开到数据中心
type ClientLeaveMsg struct {
	AccountID            uint64   `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientLeaveMsg) Reset()         { *m = ClientLeaveMsg{} }
func (m *ClientLeaveMsg) String() string { return proto.CompactTextString(m) }
func (*ClientLeaveMsg) ProtoMessage()    {}
func (*ClientLeaveMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *ClientLeaveMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientLeaveMsg.Unmarshal(m, b)
}
func (m *ClientLeaveMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientLeaveMsg.Marshal(b, m, deterministic)
}
func (m *ClientLeaveMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientLeaveMsg.Merge(m, src)
}
func (m *ClientLeaveMsg) XXX_Size() int {
	return xxx_messageInfo_ClientLeaveMsg.Size(m)
}
func (m *ClientLeaveMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientLeaveMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ClientLeaveMsg proto.InternalMessageInfo

func (m *ClientLeaveMsg) GetAccountID() uint64 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *ClientLeaveMsg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*CL_LS_ReqRegisterAccoutMsg)(nil), "account_proto.CL_LS_ReqRegisterAccoutMsg")
	proto.RegisterType((*CL_LS_ReqLoginMsg)(nil), "account_proto.CL_LS_ReqLoginMsg")
	proto.RegisterType((*LS_CL_RespnLoginMsg)(nil), "account_proto.LS_CL_RespnLoginMsg")
	proto.RegisterType((*CL_LS_ReqCreatePlayerMsg)(nil), "account_proto.CL_LS_ReqCreatePlayerMsg")
	proto.RegisterType((*ClientLeaveMsg)(nil), "account_proto.ClientLeaveMsg")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x72, 0xda, 0x30,
	0x10, 0x86, 0xc7, 0x80, 0xc1, 0xde, 0x0e, 0x6d, 0x71, 0x2f, 0x2a, 0x74, 0x5a, 0xc6, 0x27, 0x4e,
	0xf4, 0xd0, 0x27, 0xa0, 0xee, 0xa5, 0x1d, 0x37, 0x93, 0x11, 0x39, 0xe5, 0x10, 0x8d, 0xb0, 0x77,
	0x88, 0x33, 0x58, 0x72, 0x24, 0x99, 0x84, 0xc7, 0xc9, 0xd3, 0xe4, 0xb5, 0x32, 0x12, 0xc1, 0xc1,
	0xcc, 0x24, 0xb7, 0x9c, 0xac, 0xff, 0x5f, 0x7b, 0xf7, 0xfb, 0xd7, 0x82, 0x21, 0xcf, 0x32, 0x59,
	0x0b, 0x33, 0xaf, 0x94, 0x34, 0x32, 0x3a, 0x48, 0xe6, 0x64, 0xfc, 0xe8, 0xc1, 0x38, 0x49, 0x59,
	0xba, 0x64, 0x14, 0x6f, 0x29, 0xae, 0x0b, 0x6d, 0x50, 0x2d, 0xec, 0x2b, 0xe6, 0xbf, 0x5e, 0x47,
	0x63, 0x08, 0x6a, 0x8d, 0x4a, 0xf0, 0x12, 0x89, 0x37, 0xf5, 0x66, 0x21, 0x6d, 0xb4, 0xad, 0x55,
	0x5c, 0xeb, 0x3b, 0xa9, 0x72, 0xd2, 0xd9, 0xd7, 0x0e, 0x3a, 0xfa, 0x01, 0x1f, 0xb2, 0x4d, 0x81,
	0xc2, 0x30, 0xb3, 0xab, 0x90, 0x74, 0xa7, 0xde, 0x6c, 0x48, 0x61, 0x6f, 0x5d, 0xec, 0x2a, 0x8c,
	0x26, 0x10, 0x56, 0xd7, 0x52, 0x20, 0x13, 0x75, 0x49, 0x7a, 0xae, 0x1c, 0x38, 0xe3, 0xac, 0x2e,
	0xa3, 0xaf, 0x10, 0x94, 0x3c, 0x63, 0x3c, 0xcf, 0x15, 0xf1, 0x5d, 0xe7, 0x41, 0xc9, 0xb3, 0x45,
	0x9e, 0xab, 0x88, 0xc0, 0x60, 0x8b, 0x4a, 0x17, 0x52, 0x90, 0xbe, 0xfb, 0xea, 0x20, 0xe3, 0x07,
	0x0f, 0x46, 0x4d, 0x92, 0x54, 0xae, 0x0b, 0xf1, 0xae, 0x01, 0x8e, 0x19, 0x7b, 0xaf, 0x32, 0xfa,
	0x6d, 0xc6, 0x0d, 0x7c, 0x49, 0x97, 0x2c, 0x49, 0x19, 0x45, 0x5d, 0x89, 0x06, 0x72, 0x02, 0xa1,
	0x42, 0x6d, 0x58, 0x26, 0xf3, 0x3d, 0xa5, 0x4f, 0x03, 0x6b, 0x24, 0x32, 0xc7, 0x56, 0x82, 0xce,
	0x49, 0x82, 0x6f, 0x10, 0x3e, 0xff, 0xce, 0xbf, 0x7f, 0x1c, 0x63, 0x8f, 0xbe, 0x18, 0xf1, 0x15,
	0x90, 0x66, 0x21, 0x89, 0x42, 0x6e, 0xf0, 0x7c, 0xc3, 0x77, 0xa8, 0xec, 0xc8, 0xef, 0x00, 0x95,
	0x13, 0x47, 0x9b, 0x39, 0x72, 0xa2, 0xcf, 0xd0, 0xd5, 0x78, 0xef, 0x06, 0xfa, 0xd4, 0x1e, 0xad,
	0x73, 0x23, 0x57, 0x6e, 0x8a, 0x4f, 0xed, 0x31, 0xfe, 0x07, 0x1f, 0x13, 0xb7, 0x90, 0x14, 0xf9,
	0x16, 0x6d, 0xd7, 0x16, 0x8f, 0x77, 0xc2, 0xf3, 0x56, 0x92, 0xdf, 0xa3, 0xcb, 0x4f, 0xf3, 0x9f,
	0xad, 0xab, 0xb9, 0xea, 0xbb, 0xc7, 0xaf, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x04, 0x4a,
	0x3b, 0xc1, 0x02, 0x00, 0x00,
}