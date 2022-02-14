// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/collection.proto

package entities

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

type Collection struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TransactionIds       [][]byte `protobuf:"bytes,2,rep,name=transaction_ids,json=transactionIds,proto3" json:"transaction_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_b302551359ed99bf, []int{0}
}

func (m *Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collection.Unmarshal(m, b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return xxx_messageInfo_Collection.Size(m)
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Collection) GetTransactionIds() [][]byte {
	if m != nil {
		return m.TransactionIds
	}
	return nil
}

type CollectionGuarantee struct {
	CollectionId         []byte   `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Signatures           [][]byte `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionGuarantee) Reset()         { *m = CollectionGuarantee{} }
func (m *CollectionGuarantee) String() string { return proto.CompactTextString(m) }
func (*CollectionGuarantee) ProtoMessage()    {}
func (*CollectionGuarantee) Descriptor() ([]byte, []int) {
	return fileDescriptor_b302551359ed99bf, []int{1}
}

func (m *CollectionGuarantee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionGuarantee.Unmarshal(m, b)
}
func (m *CollectionGuarantee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionGuarantee.Marshal(b, m, deterministic)
}
func (m *CollectionGuarantee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionGuarantee.Merge(m, src)
}
func (m *CollectionGuarantee) XXX_Size() int {
	return xxx_messageInfo_CollectionGuarantee.Size(m)
}
func (m *CollectionGuarantee) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionGuarantee.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionGuarantee proto.InternalMessageInfo

func (m *CollectionGuarantee) GetCollectionId() []byte {
	if m != nil {
		return m.CollectionId
	}
	return nil
}

func (m *CollectionGuarantee) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func init() {
	proto.RegisterType((*Collection)(nil), "flow.entities.Collection")
	proto.RegisterType((*CollectionGuarantee)(nil), "flow.entities.CollectionGuarantee")
}

func init() { proto.RegisterFile("flow/entities/collection.proto", fileDescriptor_b302551359ed99bf) }

var fileDescriptor_b302551359ed99bf = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xd9, 0x0a, 0x1e, 0x86, 0xee, 0x0a, 0xf1, 0xb2, 0x07, 0x59, 0x96, 0xf5, 0x60, 0x4f,
	0x89, 0xe0, 0x37, 0x50, 0x44, 0x7a, 0x93, 0x1e, 0x7b, 0x29, 0x69, 0x93, 0xc6, 0x81, 0x9a, 0x91,
	0x64, 0x82, 0x5f, 0x5f, 0x8c, 0xf4, 0x8f, 0x97, 0x39, 0xbc, 0xdf, 0xe3, 0xf1, 0xde, 0xc0, 0x69,
	0x9c, 0xe8, 0x5b, 0x59, 0xcf, 0xc8, 0x68, 0xa3, 0x1a, 0x68, 0x9a, 0xec, 0xc0, 0x48, 0x5e, 0x7e,
	0x05, 0x62, 0x12, 0xfb, 0x5f, 0x2e, 0x67, 0x7e, 0x79, 0x05, 0x78, 0x59, 0x2c, 0xe2, 0x00, 0x05,
	0x9a, 0xe3, 0xee, 0xbc, 0xab, 0xca, 0xa6, 0x40, 0x23, 0x1e, 0xe0, 0x86, 0x83, 0xf6, 0x51, 0x67,
	0xdc, 0xa1, 0x89, 0xc7, 0xe2, 0x7c, 0x55, 0x95, 0xcd, 0x61, 0x23, 0xd7, 0x26, 0x5e, 0x5a, 0xb8,
	0x5d, 0x63, 0xde, 0x92, 0x0e, 0xda, 0xb3, 0xb5, 0xe2, 0x1e, 0xf6, 0x6b, 0x81, 0x6e, 0x89, 0x2e,
	0x57, 0xb1, 0x36, 0xe2, 0x04, 0x10, 0xd1, 0x79, 0xcd, 0x29, 0xd8, 0x39, 0x7f, 0xa3, 0x3c, 0xbf,
	0xc3, 0x1d, 0x05, 0x27, 0xc9, 0xe7, 0xe6, 0x79, 0x45, 0x9f, 0xc6, 0x65, 0x42, 0xfb, 0xe8, 0x90,
	0x3f, 0x52, 0x2f, 0x07, 0xfa, 0x54, 0x7f, 0x26, 0x95, 0xcf, 0xec, 0x54, 0x8e, 0xd4, 0xbf, 0xa7,
	0xf4, 0xd7, 0x19, 0x3d, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x06, 0xa1, 0x22, 0x2c, 0x01,
	0x00, 0x00,
}
