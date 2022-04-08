// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query/query.proto

package query

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

type SimpleQuery struct {
	Index                int32    `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleQuery) Reset()         { *m = SimpleQuery{} }
func (m *SimpleQuery) String() string { return proto.CompactTextString(m) }
func (*SimpleQuery) ProtoMessage()    {}
func (*SimpleQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf5d1f792f5992be, []int{0}
}

func (m *SimpleQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleQuery.Unmarshal(m, b)
}
func (m *SimpleQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleQuery.Marshal(b, m, deterministic)
}
func (m *SimpleQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleQuery.Merge(m, src)
}
func (m *SimpleQuery) XXX_Size() int {
	return xxx_messageInfo_SimpleQuery.Size(m)
}
func (m *SimpleQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleQuery proto.InternalMessageInfo

func (m *SimpleQuery) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type QueryResponse struct {
	Value                int32    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf5d1f792f5992be, []int{1}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type CuckooBucketQuery struct {
	BucketId             uint64   `protobuf:"varint,1,opt,name=BucketId,proto3" json:"BucketId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CuckooBucketQuery) Reset()         { *m = CuckooBucketQuery{} }
func (m *CuckooBucketQuery) String() string { return proto.CompactTextString(m) }
func (*CuckooBucketQuery) ProtoMessage()    {}
func (*CuckooBucketQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf5d1f792f5992be, []int{2}
}

func (m *CuckooBucketQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CuckooBucketQuery.Unmarshal(m, b)
}
func (m *CuckooBucketQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CuckooBucketQuery.Marshal(b, m, deterministic)
}
func (m *CuckooBucketQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CuckooBucketQuery.Merge(m, src)
}
func (m *CuckooBucketQuery) XXX_Size() int {
	return xxx_messageInfo_CuckooBucketQuery.Size(m)
}
func (m *CuckooBucketQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CuckooBucketQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CuckooBucketQuery proto.InternalMessageInfo

func (m *CuckooBucketQuery) GetBucketId() uint64 {
	if m != nil {
		return m.BucketId
	}
	return 0
}

type CuckooBucketResponse struct {
	Bucket               []uint64 `protobuf:"varint,1,rep,packed,name=Bucket,proto3" json:"Bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CuckooBucketResponse) Reset()         { *m = CuckooBucketResponse{} }
func (m *CuckooBucketResponse) String() string { return proto.CompactTextString(m) }
func (*CuckooBucketResponse) ProtoMessage()    {}
func (*CuckooBucketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf5d1f792f5992be, []int{3}
}

func (m *CuckooBucketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CuckooBucketResponse.Unmarshal(m, b)
}
func (m *CuckooBucketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CuckooBucketResponse.Marshal(b, m, deterministic)
}
func (m *CuckooBucketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CuckooBucketResponse.Merge(m, src)
}
func (m *CuckooBucketResponse) XXX_Size() int {
	return xxx_messageInfo_CuckooBucketResponse.Size(m)
}
func (m *CuckooBucketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CuckooBucketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CuckooBucketResponse proto.InternalMessageInfo

func (m *CuckooBucketResponse) GetBucket() []uint64 {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleQuery)(nil), "query.SimpleQuery")
	proto.RegisterType((*QueryResponse)(nil), "query.QueryResponse")
	proto.RegisterType((*CuckooBucketQuery)(nil), "query.CuckooBucketQuery")
	proto.RegisterType((*CuckooBucketResponse)(nil), "query.CuckooBucketResponse")
}

func init() { proto.RegisterFile("query/query.proto", fileDescriptor_bf5d1f792f5992be) }

var fileDescriptor_bf5d1f792f5992be = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2c, 0x4d, 0x2d,
	0xaa, 0xd4, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x92, 0x32,
	0x17, 0x77, 0x70, 0x66, 0x6e, 0x41, 0x4e, 0x6a, 0x20, 0x88, 0x2b, 0x24, 0xc2, 0xc5, 0xea, 0x99,
	0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0xe1, 0x28, 0xa9, 0x72, 0xf1,
	0x82, 0xa5, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41, 0xca, 0xc2, 0x12, 0x73, 0x4a,
	0x53, 0x61, 0xca, 0xc0, 0x1c, 0x25, 0x7d, 0x2e, 0x41, 0xe7, 0xd2, 0xe4, 0xec, 0xfc, 0x7c, 0xa7,
	0xd2, 0xe4, 0xec, 0xd4, 0x12, 0x88, 0x89, 0x52, 0x5c, 0x1c, 0x10, 0xae, 0x67, 0x0a, 0x58, 0x35,
	0x4b, 0x10, 0x9c, 0xaf, 0xa4, 0xc7, 0x25, 0x82, 0xac, 0x01, 0x6e, 0xbc, 0x18, 0x17, 0x1b, 0x44,
	0x44, 0x82, 0x51, 0x81, 0x59, 0x83, 0x25, 0x08, 0xca, 0x33, 0x5a, 0xc6, 0xc8, 0xc5, 0x03, 0x36,
	0x35, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x0d, 0xe4, 0xfa, 0xbc, 0x74, 0x98, 0xeb,
	0x25, 0xf4, 0x20, 0x3e, 0xc4, 0x70, 0x85, 0x94, 0x34, 0x16, 0x19, 0x98, 0x75, 0x4a, 0x0c, 0x42,
	0x7e, 0x5c, 0xfc, 0xce, 0xf9, 0x79, 0x25, 0x99, 0x79, 0xa5, 0xf9, 0xa5, 0xc5, 0x94, 0x99, 0xa5,
	0xc1, 0x68, 0xc0, 0xe8, 0x24, 0x1c, 0x25, 0x98, 0x5a, 0x91, 0x08, 0x0a, 0x56, 0xbd, 0xe4, 0xfc,
	0x5c, 0x48, 0xb8, 0x27, 0xb1, 0x81, 0x03, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x27, 0xfd,
	0x74, 0xd6, 0x8d, 0x01, 0x00, 0x00,
}
