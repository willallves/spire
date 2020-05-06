// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bundle.proto

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

type Bundle struct {
	// The SPIFFE ID of the trust domain the bundle belongs to, e.g.
	// spiffe://example.org.
	TrustDomainId string `protobuf:"bytes,1,opt,name=trust_domain_id,json=trustDomainId,proto3" json:"trust_domain_id,omitempty"`
	// X.509 authorities for authenticating X509-SVIDs.
	X509Authorities []*X509Certificate `protobuf:"bytes,2,rep,name=x509_authorities,json=x509Authorities,proto3" json:"x509_authorities,omitempty"`
	// JWT authorities for authenticating JWT-SVIDs.
	JwtAuthorities []*JWTKey `protobuf:"bytes,3,rep,name=jwt_authorities,json=jwtAuthorities,proto3" json:"jwt_authorities,omitempty"`
	// A hint on how often the bundle should be refreshed from the bundle
	// provider, in seconds. Can be zero (meaning no hint available).
	RefreshHint int64 `protobuf:"varint,4,opt,name=refresh_hint,json=refreshHint,proto3" json:"refresh_hint,omitempty"`
	// The sequence number of the bundle.
	SequenceNumber       uint64   `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bundle) Reset()         { *m = Bundle{} }
func (m *Bundle) String() string { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()    {}
func (*Bundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf01a1817f9fc5c2, []int{0}
}

func (m *Bundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bundle.Unmarshal(m, b)
}
func (m *Bundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bundle.Marshal(b, m, deterministic)
}
func (m *Bundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundle.Merge(m, src)
}
func (m *Bundle) XXX_Size() int {
	return xxx_messageInfo_Bundle.Size(m)
}
func (m *Bundle) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundle.DiscardUnknown(m)
}

var xxx_messageInfo_Bundle proto.InternalMessageInfo

func (m *Bundle) GetTrustDomainId() string {
	if m != nil {
		return m.TrustDomainId
	}
	return ""
}

func (m *Bundle) GetX509Authorities() []*X509Certificate {
	if m != nil {
		return m.X509Authorities
	}
	return nil
}

func (m *Bundle) GetJwtAuthorities() []*JWTKey {
	if m != nil {
		return m.JwtAuthorities
	}
	return nil
}

func (m *Bundle) GetRefreshHint() int64 {
	if m != nil {
		return m.RefreshHint
	}
	return 0
}

func (m *Bundle) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

type X509Certificate struct {
	// The ASN.1 DER encoded bytes of the X.509 certificate.
	Asn1                 []byte   `protobuf:"bytes,1,opt,name=asn1,proto3" json:"asn1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *X509Certificate) Reset()         { *m = X509Certificate{} }
func (m *X509Certificate) String() string { return proto.CompactTextString(m) }
func (*X509Certificate) ProtoMessage()    {}
func (*X509Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf01a1817f9fc5c2, []int{1}
}

func (m *X509Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X509Certificate.Unmarshal(m, b)
}
func (m *X509Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X509Certificate.Marshal(b, m, deterministic)
}
func (m *X509Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X509Certificate.Merge(m, src)
}
func (m *X509Certificate) XXX_Size() int {
	return xxx_messageInfo_X509Certificate.Size(m)
}
func (m *X509Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_X509Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_X509Certificate proto.InternalMessageInfo

func (m *X509Certificate) GetAsn1() []byte {
	if m != nil {
		return m.Asn1
	}
	return nil
}

type JWTKey struct {
	// The PKIX encoded public key.
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The key identifier.
	KeyId string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// When the key expires (seconds since Unix epoch). If zero, the key does
	// not expire.
	ExpiresAt            int64    `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTKey) Reset()         { *m = JWTKey{} }
func (m *JWTKey) String() string { return proto.CompactTextString(m) }
func (*JWTKey) ProtoMessage()    {}
func (*JWTKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf01a1817f9fc5c2, []int{2}
}

func (m *JWTKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTKey.Unmarshal(m, b)
}
func (m *JWTKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTKey.Marshal(b, m, deterministic)
}
func (m *JWTKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTKey.Merge(m, src)
}
func (m *JWTKey) XXX_Size() int {
	return xxx_messageInfo_JWTKey.Size(m)
}
func (m *JWTKey) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTKey.DiscardUnknown(m)
}

var xxx_messageInfo_JWTKey proto.InternalMessageInfo

func (m *JWTKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *JWTKey) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *JWTKey) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

type BundleMask struct {
	// trust_domain_id field mask.
	TrustDomainId bool `protobuf:"varint,1,opt,name=trust_domain_id,json=trustDomainId,proto3" json:"trust_domain_id,omitempty"`
	// x509_authorities field mask.
	X509Authorities bool `protobuf:"varint,2,opt,name=x509_authorities,json=x509Authorities,proto3" json:"x509_authorities,omitempty"`
	// jwt_authorities field mask.
	JwtAuthorities bool `protobuf:"varint,3,opt,name=jwt_authorities,json=jwtAuthorities,proto3" json:"jwt_authorities,omitempty"`
	// refresh_hint field mask.
	RefreshHint bool `protobuf:"varint,4,opt,name=refresh_hint,json=refreshHint,proto3" json:"refresh_hint,omitempty"`
	// sequence_number field mask.
	SequenceNumber       bool     `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BundleMask) Reset()         { *m = BundleMask{} }
func (m *BundleMask) String() string { return proto.CompactTextString(m) }
func (*BundleMask) ProtoMessage()    {}
func (*BundleMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf01a1817f9fc5c2, []int{3}
}

func (m *BundleMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleMask.Unmarshal(m, b)
}
func (m *BundleMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleMask.Marshal(b, m, deterministic)
}
func (m *BundleMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleMask.Merge(m, src)
}
func (m *BundleMask) XXX_Size() int {
	return xxx_messageInfo_BundleMask.Size(m)
}
func (m *BundleMask) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleMask.DiscardUnknown(m)
}

var xxx_messageInfo_BundleMask proto.InternalMessageInfo

func (m *BundleMask) GetTrustDomainId() bool {
	if m != nil {
		return m.TrustDomainId
	}
	return false
}

func (m *BundleMask) GetX509Authorities() bool {
	if m != nil {
		return m.X509Authorities
	}
	return false
}

func (m *BundleMask) GetJwtAuthorities() bool {
	if m != nil {
		return m.JwtAuthorities
	}
	return false
}

func (m *BundleMask) GetRefreshHint() bool {
	if m != nil {
		return m.RefreshHint
	}
	return false
}

func (m *BundleMask) GetSequenceNumber() bool {
	if m != nil {
		return m.SequenceNumber
	}
	return false
}

func init() {
	proto.RegisterType((*Bundle)(nil), "spire.types.Bundle")
	proto.RegisterType((*X509Certificate)(nil), "spire.types.X509Certificate")
	proto.RegisterType((*JWTKey)(nil), "spire.types.JWTKey")
	proto.RegisterType((*BundleMask)(nil), "spire.types.BundleMask")
}

func init() { proto.RegisterFile("bundle.proto", fileDescriptor_cf01a1817f9fc5c2) }

var fileDescriptor_cf01a1817f9fc5c2 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x86, 0x49, 0xd3, 0x1b, 0xda, 0x69, 0xbd, 0x91, 0x11, 0x21, 0x0b, 0x85, 0x18, 0xd0, 0x1b,
	0x17, 0x26, 0x55, 0xe9, 0xa2, 0xe0, 0xa6, 0x55, 0xd0, 0x2a, 0xba, 0x08, 0x82, 0x22, 0x48, 0xc8,
	0xc7, 0x89, 0x99, 0xa6, 0x9d, 0xc4, 0x99, 0x13, 0xda, 0xac, 0xfd, 0x8b, 0xfe, 0x20, 0xc9, 0xa4,
	0x62, 0x5b, 0x7b, 0xa1, 0xbb, 0xe1, 0x39, 0x1f, 0x9c, 0x3c, 0x6f, 0xc8, 0x38, 0xae, 0x79, 0xba,
	0x06, 0xaf, 0x12, 0x25, 0x96, 0x74, 0x24, 0x2b, 0x26, 0xc0, 0xc3, 0xa6, 0x02, 0xe9, 0xfc, 0xea,
	0x11, 0x63, 0xa1, 0xaa, 0xf4, 0x09, 0x31, 0x51, 0xd4, 0x12, 0xc3, 0xb4, 0xdc, 0x44, 0x8c, 0x87,
	0x2c, 0xb5, 0x34, 0x5b, 0x73, 0x87, 0xc1, 0x1d, 0x85, 0xdf, 0x28, 0xba, 0x4c, 0xe9, 0x5b, 0x72,
	0x77, 0x37, 0x9d, 0xcc, 0xc2, 0xa8, 0xc6, 0xbc, 0x14, 0x0c, 0x19, 0x48, 0xab, 0x67, 0xeb, 0xee,
	0xe8, 0xc5, 0x03, 0xef, 0x60, 0xb5, 0xf7, 0x75, 0x3a, 0x99, 0xbd, 0x06, 0x81, 0x2c, 0x63, 0x49,
	0x84, 0x10, 0x98, 0xed, 0xd4, 0xfc, 0xdf, 0x10, 0x7d, 0x45, 0xcc, 0xd5, 0x16, 0x8f, 0xf6, 0xe8,
	0x6a, 0xcf, 0xbd, 0xa3, 0x3d, 0xef, 0xbf, 0x7c, 0xfe, 0x00, 0x4d, 0x70, 0xbd, 0xda, 0xe2, 0xe1,
	0xf4, 0x23, 0x32, 0x16, 0x90, 0x09, 0x90, 0x79, 0x98, 0x33, 0x8e, 0x56, 0xdf, 0xd6, 0x5c, 0x3d,
	0x18, 0xed, 0xd9, 0x3b, 0xc6, 0x91, 0xde, 0x10, 0x53, 0xc2, 0xcf, 0x1a, 0x78, 0x02, 0x21, 0xaf,
	0x37, 0x31, 0x08, 0xeb, 0xca, 0xd6, 0xdc, 0x7e, 0x70, 0xfd, 0x17, 0x7f, 0x52, 0xd4, 0x79, 0x4c,
	0xcc, 0x93, 0x6b, 0x29, 0x25, 0xfd, 0x48, 0xf2, 0xe7, 0x4a, 0xc1, 0x38, 0x50, 0x6f, 0xe7, 0x3b,
	0x31, 0xba, 0x63, 0xe8, 0x43, 0x42, 0xaa, 0x3a, 0x5e, 0xb3, 0x24, 0x2c, 0xa0, 0xd9, 0xf7, 0x0c,
	0x3b, 0xd2, 0x96, 0xef, 0x13, 0xa3, 0x80, 0xa6, 0x35, 0xd8, 0x53, 0x06, 0xaf, 0x0a, 0x68, 0x96,
	0x69, 0x3b, 0x05, 0xbb, 0xf6, 0xcb, 0x64, 0x18, 0xa1, 0xa5, 0xab, 0x83, 0x87, 0x7b, 0x32, 0x47,
	0xe7, 0xb7, 0x46, 0x48, 0x97, 0xc5, 0xc7, 0x48, 0x16, 0xb7, 0xe5, 0x31, 0x38, 0xcd, 0xe3, 0xe9,
	0xd9, 0x3c, 0xda, 0xc6, 0xff, 0x8c, 0xdf, 0x9c, 0x33, 0xde, 0x76, 0x5e, 0x22, 0x77, 0x70, 0x91,
	0xdc, 0xc1, 0xa9, 0xdc, 0xc5, 0xe4, 0x9b, 0xf7, 0x83, 0x61, 0x5e, 0xc7, 0x5e, 0x52, 0x6e, 0x7c,
	0x59, 0xb1, 0x2c, 0x03, 0x5f, 0x05, 0xec, 0xab, 0x1f, 0xb2, 0x7b, 0x3f, 0xe3, 0xb0, 0x43, 0x5f,
	0x25, 0x1e, 0x1b, 0x8a, 0xbf, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xce, 0xd7, 0x1c, 0xf0, 0xb8,
	0x02, 0x00, 0x00,
}