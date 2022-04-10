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
	QueryNum             uint64   `protobuf:"varint,1,opt,name=QueryNum,proto3" json:"QueryNum,omitempty"`
	BucketId             []uint64 `protobuf:"varint,2,rep,packed,name=BucketId,proto3" json:"BucketId,omitempty"`
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

func (m *CuckooBucketQuery) GetQueryNum() uint64 {
	if m != nil {
		return m.QueryNum
	}
	return 0
}

func (m *CuckooBucketQuery) GetBucketId() []uint64 {
	if m != nil {
		return m.BucketId
	}
	return nil
}

type CuckooBucketResponse struct {
	ResponseNum          uint64   `protobuf:"varint,1,opt,name=ResponseNum,proto3" json:"ResponseNum,omitempty"`
	PackedBucket         []uint64 `protobuf:"varint,2,rep,packed,name=PackedBucket,proto3" json:"PackedBucket,omitempty"`
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

func (m *CuckooBucketResponse) GetResponseNum() uint64 {
	if m != nil {
		return m.ResponseNum
	}
	return 0
}

func (m *CuckooBucketResponse) GetPackedBucket() []uint64 {
	if m != nil {
		return m.PackedBucket
	}
	return nil
}

type HashTableInfoQuery struct {
	Dummy                uint64   `protobuf:"varint,1,opt,name=Dummy,proto3" json:"Dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashTableInfoQuery) Reset()         { *m = HashTableInfoQuery{} }
func (m *HashTableInfoQuery) String() string { return proto.CompactTextString(m) }
func (*HashTableInfoQuery) ProtoMessage()    {}
func (*HashTableInfoQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf5d1f792f5992be, []int{4}
}

func (m *HashTableInfoQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashTableInfoQuery.Unmarshal(m, b)
}
func (m *HashTableInfoQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashTableInfoQuery.Marshal(b, m, deterministic)
}
func (m *HashTableInfoQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashTableInfoQuery.Merge(m, src)
}
func (m *HashTableInfoQuery) XXX_Size() int {
	return xxx_messageInfo_HashTableInfoQuery.Size(m)
}
func (m *HashTableInfoQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_HashTableInfoQuery.DiscardUnknown(m)
}

var xxx_messageInfo_HashTableInfoQuery proto.InternalMessageInfo

func (m *HashTableInfoQuery) GetDummy() uint64 {
	if m != nil {
		return m.Dummy
	}
	return 0
}

type HashTableInfoResponse struct {
	Size                 uint64   `protobuf:"varint,1,opt,name=Size,proto3" json:"Size,omitempty"`
	Load                 uint64   `protobuf:"varint,2,opt,name=Load,proto3" json:"Load,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashTableInfoResponse) Reset()         { *m = HashTableInfoResponse{} }
func (m *HashTableInfoResponse) String() string { return proto.CompactTextString(m) }
func (*HashTableInfoResponse) ProtoMessage()    {}
func (*HashTableInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf5d1f792f5992be, []int{5}
}

func (m *HashTableInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashTableInfoResponse.Unmarshal(m, b)
}
func (m *HashTableInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashTableInfoResponse.Marshal(b, m, deterministic)
}
func (m *HashTableInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashTableInfoResponse.Merge(m, src)
}
func (m *HashTableInfoResponse) XXX_Size() int {
	return xxx_messageInfo_HashTableInfoResponse.Size(m)
}
func (m *HashTableInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HashTableInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HashTableInfoResponse proto.InternalMessageInfo

func (m *HashTableInfoResponse) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *HashTableInfoResponse) GetLoad() uint64 {
	if m != nil {
		return m.Load
	}
	return 0
}

func init() {
	proto.RegisterType((*SimpleQuery)(nil), "query.SimpleQuery")
	proto.RegisterType((*QueryResponse)(nil), "query.QueryResponse")
	proto.RegisterType((*CuckooBucketQuery)(nil), "query.CuckooBucketQuery")
	proto.RegisterType((*CuckooBucketResponse)(nil), "query.CuckooBucketResponse")
	proto.RegisterType((*HashTableInfoQuery)(nil), "query.HashTableInfoQuery")
	proto.RegisterType((*HashTableInfoResponse)(nil), "query.HashTableInfoResponse")
}

func init() { proto.RegisterFile("query/query.proto", fileDescriptor_bf5d1f792f5992be) }

var fileDescriptor_bf5d1f792f5992be = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0xfa, 0x09, 0xc6, 0x0c, 0x18, 0x65, 0xc5, 0x04, 0xd1, 0x03, 0x59, 0x63, 0x42, 0x3c,
	0xa0, 0xd1, 0x1f, 0x60, 0x02, 0x46, 0x25, 0x2a, 0x51, 0x6a, 0x3c, 0x18, 0x2f, 0x4b, 0x3b, 0x6a,
	0xd3, 0x76, 0xb7, 0xb6, 0x5d, 0x03, 0xfe, 0x6f, 0xef, 0x66, 0x77, 0xdb, 0x42, 0x03, 0x37, 0x2f,
	0xcd, 0xbc, 0x79, 0x6f, 0xde, 0x74, 0x66, 0x07, 0x1a, 0x9f, 0x12, 0xe3, 0xd9, 0x89, 0xfe, 0xf6,
	0xa2, 0x58, 0xa4, 0x82, 0x54, 0x35, 0xa0, 0x87, 0x50, 0xb3, 0xbd, 0x30, 0x0a, 0xf0, 0x51, 0x41,
	0xd2, 0x84, 0xea, 0x90, 0xbb, 0x38, 0x6d, 0x59, 0x1d, 0xab, 0x5b, 0x1d, 0x1b, 0x40, 0x8f, 0x60,
	0x53, 0xd3, 0x63, 0x4c, 0x22, 0xc1, 0x13, 0x54, 0xb2, 0x67, 0x16, 0x48, 0xcc, 0x65, 0x1a, 0xd0,
	0x5b, 0x68, 0x0c, 0xa4, 0xe3, 0x0b, 0xd1, 0x97, 0x8e, 0x8f, 0xa9, 0x71, 0x6c, 0xc3, 0x86, 0x0e,
	0x46, 0x32, 0xd4, 0xea, 0xca, 0xb8, 0xc0, 0x8a, 0x33, 0xd2, 0xa1, 0xdb, 0xfa, 0xdf, 0x59, 0x53,
	0x5c, 0x8e, 0xe9, 0x2b, 0x34, 0x17, 0xcd, 0x8a, 0xd6, 0x1d, 0xa8, 0xe5, 0xf1, 0xdc, 0x72, 0x31,
	0x45, 0x28, 0xd4, 0x1f, 0x98, 0xe3, 0xa3, 0x6b, 0x2a, 0x33, 0xe7, 0x52, 0x8e, 0x1e, 0x03, 0xb9,
	0x61, 0xc9, 0xc7, 0x13, 0x9b, 0x04, 0x38, 0xe4, 0x6f, 0xa2, 0x98, 0xfe, 0x52, 0x86, 0xe1, 0x2c,
	0x73, 0x35, 0x80, 0x5e, 0xc0, 0x6e, 0x49, 0x5b, 0xfc, 0x0a, 0x81, 0x8a, 0xed, 0x7d, 0x63, 0xa6,
	0xd6, 0xb1, 0xca, 0xdd, 0x09, 0xa6, 0xc6, 0xd1, 0x39, 0x15, 0x9f, 0xfd, 0x58, 0x50, 0xd7, 0x0d,
	0x6c, 0x8c, 0xbf, 0x3c, 0x07, 0xc9, 0x95, 0x5a, 0x3a, 0x7f, 0xcf, 0x97, 0xde, 0xea, 0x99, 0x87,
	0x59, 0x5a, 0x5e, 0x7b, 0x7f, 0x05, 0x93, 0xb7, 0xa7, 0xff, 0xc8, 0x08, 0xb6, 0x06, 0x82, 0xa7,
	0x1e, 0x97, 0x42, 0x26, 0x7f, 0xf3, 0xea, 0x5a, 0xa7, 0x16, 0xb9, 0x87, 0xed, 0x6b, 0x4c, 0x4b,
	0xc3, 0x92, 0xbd, 0xac, 0x6c, 0x79, 0x5d, 0xed, 0x83, 0x55, 0xd4, 0xdc, 0xb2, 0xbf, 0xf3, 0xd2,
	0xc0, 0x29, 0x53, 0xc7, 0xd5, 0x73, 0x44, 0x68, 0xae, 0x6f, 0xb2, 0xae, 0xcf, 0xef, 0xfc, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xf0, 0xb0, 0x31, 0x0b, 0x93, 0x02, 0x00, 0x00,
}
