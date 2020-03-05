// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entry.proto

package types

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

type Entry struct {
	// Globally unique ID for the entry.
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	// The SPIFFE ID of the identity described by this entry.
	SpiffeId string `protobuf:"bytes,3,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// Who the entry is delegated to. If the entry describes a node, this is
	// set to the SPIFFE ID of the SPIRE server of the trust domain (e.g.
	// spiffe://example.org/spire/server). Otherwise, it will be set to a node
	// SPIFFE ID.
	DelegatedTo string `protobuf:"bytes,2,opt,name=delegated_to,json=delegatedTo,proto3" json:"delegated_to,omitempty"`
	// The selectors which identify which entities match this entry. If this is
	// an entry for a node, these selectors represent selectors produced by
	// node attestation. Otherwise, these selectors represent those produced by
	// workload attestation.
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// The time to live for identities issued for this entry.
	Ttl int32 `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// A list of trust domains the identity described by this entry federates
	// with.
	FederatesWith []string `protobuf:"bytes,5,rep,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
	// Whether or not the identity described by this entry is an administrative
	// workload. Administrative workloads are granted additional access to
	// various managerial server APIs, such as entry registration.
	Admin bool `protobuf:"varint,7,opt,name=admin,proto3" json:"admin,omitempty"`
	// Whether or not the identity described by this entry represents a
	// downstream SPIRE server. Downstream SPIRE servers have additional access
	// to various signing APIs, such as those used to sign X.509 CA
	// certificates and publish JWT signing keys.
	Downstream bool `protobuf:"varint,8,opt,name=downstream,proto3" json:"downstream,omitempty"`
	// When the entry expires, in seconds since unix epoch.
	ExpiresAt int64 `protobuf:"varint,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// A list of DNS names associated with the identity described by this entry.
	DnsNames             []string `protobuf:"bytes,10,rep,name=dns_names,json=dnsNames,proto3" json:"dns_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_daa6c5b6c627940f, []int{0}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Entry) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *Entry) GetDelegatedTo() string {
	if m != nil {
		return m.DelegatedTo
	}
	return ""
}

func (m *Entry) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *Entry) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Entry) GetFederatesWith() []string {
	if m != nil {
		return m.FederatesWith
	}
	return nil
}

func (m *Entry) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

func (m *Entry) GetDownstream() bool {
	if m != nil {
		return m.Downstream
	}
	return false
}

func (m *Entry) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Entry) GetDnsNames() []string {
	if m != nil {
		return m.DnsNames
	}
	return nil
}

// Field mask for Entry fields
type EntryMask struct {
	// id field mask
	Id bool `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	// spiffe_id field mask
	SpiffeId bool `protobuf:"varint,3,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// delegated_to field mask
	DelegatedTo bool `protobuf:"varint,2,opt,name=delegated_to,json=delegatedTo,proto3" json:"delegated_to,omitempty"`
	// selectors field mask
	Selectors bool `protobuf:"varint,1,opt,name=selectors,proto3" json:"selectors,omitempty"`
	// ttl field mask
	Ttl bool `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// federates_with field mask
	FederatesWith bool `protobuf:"varint,5,opt,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
	// admin field mask
	Admin bool `protobuf:"varint,7,opt,name=admin,proto3" json:"admin,omitempty"`
	// downstream field mask
	Downstream bool `protobuf:"varint,8,opt,name=downstream,proto3" json:"downstream,omitempty"`
	// expires_at field mask
	ExpiresAt bool `protobuf:"varint,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// dns_names field mask
	DnsNames             bool     `protobuf:"varint,10,opt,name=dns_names,json=dnsNames,proto3" json:"dns_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryMask) Reset()         { *m = EntryMask{} }
func (m *EntryMask) String() string { return proto.CompactTextString(m) }
func (*EntryMask) ProtoMessage()    {}
func (*EntryMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_daa6c5b6c627940f, []int{1}
}

func (m *EntryMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryMask.Unmarshal(m, b)
}
func (m *EntryMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryMask.Marshal(b, m, deterministic)
}
func (m *EntryMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryMask.Merge(m, src)
}
func (m *EntryMask) XXX_Size() int {
	return xxx_messageInfo_EntryMask.Size(m)
}
func (m *EntryMask) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryMask.DiscardUnknown(m)
}

var xxx_messageInfo_EntryMask proto.InternalMessageInfo

func (m *EntryMask) GetId() bool {
	if m != nil {
		return m.Id
	}
	return false
}

func (m *EntryMask) GetSpiffeId() bool {
	if m != nil {
		return m.SpiffeId
	}
	return false
}

func (m *EntryMask) GetDelegatedTo() bool {
	if m != nil {
		return m.DelegatedTo
	}
	return false
}

func (m *EntryMask) GetSelectors() bool {
	if m != nil {
		return m.Selectors
	}
	return false
}

func (m *EntryMask) GetTtl() bool {
	if m != nil {
		return m.Ttl
	}
	return false
}

func (m *EntryMask) GetFederatesWith() bool {
	if m != nil {
		return m.FederatesWith
	}
	return false
}

func (m *EntryMask) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

func (m *EntryMask) GetDownstream() bool {
	if m != nil {
		return m.Downstream
	}
	return false
}

func (m *EntryMask) GetExpiresAt() bool {
	if m != nil {
		return m.ExpiresAt
	}
	return false
}

func (m *EntryMask) GetDnsNames() bool {
	if m != nil {
		return m.DnsNames
	}
	return false
}

func init() {
	proto.RegisterType((*Entry)(nil), "spire.types.Entry")
	proto.RegisterType((*EntryMask)(nil), "spire.types.EntryMask")
}

func init() { proto.RegisterFile("entry.proto", fileDescriptor_daa6c5b6c627940f) }

var fileDescriptor_daa6c5b6c627940f = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xdd, 0x4a, 0xe3, 0x40,
	0x14, 0xc7, 0x49, 0xb2, 0xe9, 0x4e, 0x4e, 0x76, 0xcb, 0x32, 0xec, 0xc2, 0xb0, 0x7e, 0xc5, 0x82,
	0x90, 0xab, 0x44, 0xec, 0x13, 0x28, 0x78, 0xe1, 0x85, 0x5e, 0x44, 0x41, 0xf0, 0x26, 0xa4, 0x9d,
	0xd3, 0x76, 0xb0, 0xf9, 0x20, 0x73, 0xa4, 0xed, 0xfb, 0xf8, 0x0e, 0xbe, 0x9e, 0x64, 0xd2, 0xd6,
	0x56, 0xaa, 0x5e, 0x79, 0x97, 0xf3, 0xfb, 0x9f, 0x70, 0xce, 0xfc, 0x38, 0xe0, 0x63, 0x41, 0xf5,
	0x22, 0xaa, 0xea, 0x92, 0x4a, 0xee, 0xeb, 0x4a, 0xd5, 0x18, 0xd1, 0xa2, 0x42, 0xfd, 0xff, 0xc8,
	0x14, 0x71, 0x81, 0x73, 0x8a, 0x0d, 0x89, 0x35, 0x4e, 0x71, 0x48, 0x65, 0xdd, 0x76, 0xf7, 0x5e,
	0x6c, 0x70, 0x2f, 0x9b, 0xbf, 0x79, 0x17, 0x6c, 0x25, 0x45, 0x27, 0xb0, 0x42, 0x2f, 0xb1, 0x95,
	0xe4, 0x7b, 0xe0, 0xe9, 0x4a, 0x8d, 0x46, 0x98, 0x2a, 0x29, 0x1c, 0x83, 0x59, 0x0b, 0xae, 0x24,
	0x3f, 0x86, 0x5f, 0x12, 0xa7, 0x38, 0xce, 0x08, 0x65, 0x4a, 0xa5, 0xb0, 0x4d, 0xee, 0xaf, 0xd9,
	0x5d, 0xc9, 0xfb, 0xe0, 0xad, 0x66, 0x69, 0x61, 0x05, 0x4e, 0xe8, 0x9f, 0xfd, 0x8b, 0x36, 0x76,
	0x8b, 0x6e, 0x97, 0x69, 0xf2, 0xd6, 0xc7, 0xff, 0x80, 0x43, 0x34, 0x15, 0x3f, 0x02, 0x2b, 0x74,
	0x93, 0xe6, 0x93, 0x9f, 0x40, 0x77, 0x84, 0x12, 0xeb, 0x8c, 0x50, 0xa7, 0x33, 0x45, 0x13, 0xe1,
	0x06, 0x4e, 0xe8, 0x25, 0xbf, 0xd7, 0xf4, 0x5e, 0xd1, 0x84, 0xff, 0x05, 0x37, 0x93, 0xb9, 0x2a,
	0xc4, 0xcf, 0xc0, 0x0a, 0x59, 0xd2, 0x16, 0xfc, 0x10, 0x40, 0x96, 0xb3, 0x42, 0x53, 0x8d, 0x59,
	0x2e, 0x98, 0x89, 0x36, 0x08, 0x3f, 0x00, 0xc0, 0x79, 0xb3, 0x92, 0x4e, 0x33, 0x12, 0x5e, 0x60,
	0x85, 0x4e, 0xe2, 0x2d, 0xc9, 0x39, 0x35, 0x0a, 0x64, 0xa1, 0xd3, 0x22, 0xcb, 0x51, 0x0b, 0x30,
	0x63, 0x99, 0x2c, 0xf4, 0x4d, 0x53, 0xf7, 0x9e, 0x6d, 0xf0, 0x8c, 0xb9, 0xeb, 0x4c, 0x3f, 0x6e,
	0xd8, 0x63, 0xbb, 0xed, 0xb1, 0x2f, 0xec, 0xb1, 0x6d, 0x7b, 0xfb, 0xdb, 0xf6, 0x9a, 0x7c, 0xb7,
	0x26, 0xf6, 0xb1, 0xa6, 0x26, 0xfc, 0x26, 0x4d, 0xec, 0x13, 0x4d, 0xe6, 0xad, 0x2b, 0x4d, 0x17,
	0xa7, 0x0f, 0xd1, 0x58, 0xd1, 0xe4, 0x69, 0x10, 0x0d, 0xcb, 0x3c, 0x6e, 0x15, 0xc4, 0xed, 0x55,
	0x9a, 0x0b, 0x8c, 0xdf, 0x5f, 0xe8, 0xa0, 0x63, 0x78, 0xff, 0x35, 0x00, 0x00, 0xff, 0xff, 0x46,
	0xb2, 0x6d, 0x37, 0xd6, 0x02, 0x00, 0x00,
}
