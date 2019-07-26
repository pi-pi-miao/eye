// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/geo_target_constant_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The possible statuses of a geo target constant.
type GeoTargetConstantStatusEnum_GeoTargetConstantStatus int32

const (
	// No value has been specified.
	GeoTargetConstantStatusEnum_UNSPECIFIED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	GeoTargetConstantStatusEnum_UNKNOWN GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 1
	// The geo target constant is valid.
	GeoTargetConstantStatusEnum_ENABLED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 2
	// The geo target constant is obsolete and will be removed.
	GeoTargetConstantStatusEnum_REMOVAL_PLANNED GeoTargetConstantStatusEnum_GeoTargetConstantStatus = 3
)

var GeoTargetConstantStatusEnum_GeoTargetConstantStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVAL_PLANNED",
}

var GeoTargetConstantStatusEnum_GeoTargetConstantStatus_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"ENABLED":         2,
	"REMOVAL_PLANNED": 3,
}

func (x GeoTargetConstantStatusEnum_GeoTargetConstantStatus) String() string {
	return proto.EnumName(GeoTargetConstantStatusEnum_GeoTargetConstantStatus_name, int32(x))
}

func (GeoTargetConstantStatusEnum_GeoTargetConstantStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67dba27e8548264d, []int{0, 0}
}

// Container for describing the status of a geo target constant.
type GeoTargetConstantStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoTargetConstantStatusEnum) Reset()         { *m = GeoTargetConstantStatusEnum{} }
func (m *GeoTargetConstantStatusEnum) String() string { return proto.CompactTextString(m) }
func (*GeoTargetConstantStatusEnum) ProtoMessage()    {}
func (*GeoTargetConstantStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_67dba27e8548264d, []int{0}
}

func (m *GeoTargetConstantStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoTargetConstantStatusEnum.Unmarshal(m, b)
}
func (m *GeoTargetConstantStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoTargetConstantStatusEnum.Marshal(b, m, deterministic)
}
func (m *GeoTargetConstantStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoTargetConstantStatusEnum.Merge(m, src)
}
func (m *GeoTargetConstantStatusEnum) XXX_Size() int {
	return xxx_messageInfo_GeoTargetConstantStatusEnum.Size(m)
}
func (m *GeoTargetConstantStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoTargetConstantStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GeoTargetConstantStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus", GeoTargetConstantStatusEnum_GeoTargetConstantStatus_name, GeoTargetConstantStatusEnum_GeoTargetConstantStatus_value)
	proto.RegisterType((*GeoTargetConstantStatusEnum)(nil), "google.ads.googleads.v1.enums.GeoTargetConstantStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/geo_target_constant_status.proto", fileDescriptor_67dba27e8548264d)
}

var fileDescriptor_67dba27e8548264d = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x6a, 0xf2, 0x30,
	0x18, 0xfd, 0xad, 0xf0, 0x0f, 0xe2, 0x85, 0xa5, 0xbb, 0x18, 0x6c, 0x7a, 0xa1, 0x0f, 0x90, 0x52,
	0x76, 0x97, 0xc1, 0x20, 0xd5, 0x4c, 0x64, 0x2e, 0x96, 0x39, 0x3b, 0x36, 0x0a, 0x25, 0xb3, 0x21,
	0x08, 0x9a, 0x4f, 0x4c, 0x94, 0x3d, 0xcf, 0x2e, 0xf7, 0x28, 0x7b, 0x94, 0xdd, 0xed, 0x0d, 0x46,
	0x53, 0xf5, 0xae, 0xbb, 0x09, 0x27, 0x39, 0xdf, 0x39, 0x39, 0xdf, 0x41, 0xb7, 0x0a, 0x40, 0xad,
	0x64, 0x28, 0x0a, 0x13, 0x56, 0xb0, 0x44, 0xfb, 0x28, 0x94, 0x7a, 0xb7, 0x36, 0xa1, 0x92, 0x90,
	0x5b, 0xb1, 0x55, 0xd2, 0xe6, 0x0b, 0xd0, 0xc6, 0x0a, 0x6d, 0x73, 0x63, 0x85, 0xdd, 0x19, 0xbc,
	0xd9, 0x82, 0x85, 0xa0, 0x5b, 0x89, 0xb0, 0x28, 0x0c, 0x3e, 0xe9, 0xf1, 0x3e, 0xc2, 0x4e, 0x7f,
	0xd9, 0x39, 0xda, 0x6f, 0x96, 0xa1, 0xd0, 0x1a, 0xac, 0xb0, 0x4b, 0xd0, 0x07, 0x71, 0xff, 0x1d,
	0x5d, 0x8d, 0x24, 0x3c, 0x39, 0xff, 0xc1, 0xc1, 0x7e, 0xe6, 0xdc, 0x99, 0xde, 0xad, 0xfb, 0x2f,
	0xe8, 0xa2, 0x86, 0x0e, 0xda, 0xa8, 0x35, 0xe7, 0xb3, 0x84, 0x0d, 0xc6, 0x77, 0x63, 0x36, 0xf4,
	0xff, 0x05, 0x2d, 0x74, 0x36, 0xe7, 0xf7, 0x7c, 0xfa, 0xcc, 0xfd, 0x46, 0x79, 0x61, 0x9c, 0xc6,
	0x13, 0x36, 0xf4, 0xbd, 0xe0, 0x1c, 0xb5, 0x1f, 0xd9, 0xc3, 0x34, 0xa5, 0x93, 0x3c, 0x99, 0x50,
	0xce, 0xd9, 0xd0, 0x6f, 0xc6, 0x3f, 0x0d, 0xd4, 0x5b, 0xc0, 0x1a, 0xff, 0x99, 0x3e, 0xee, 0xd4,
	0x7c, 0x9f, 0x94, 0xe9, 0x93, 0xc6, 0x6b, 0x7c, 0x90, 0x2b, 0x58, 0x09, 0xad, 0x30, 0x6c, 0x55,
	0xa8, 0xa4, 0x76, 0xbb, 0x1d, 0xcb, 0xdc, 0x2c, 0x4d, 0x4d, 0xb7, 0x37, 0xee, 0xfc, 0xf0, 0x9a,
	0x23, 0x4a, 0x3f, 0xbd, 0xee, 0xa8, 0xb2, 0xa2, 0x85, 0xc1, 0x15, 0x2c, 0x51, 0x1a, 0xe1, 0xb2,
	0x09, 0xf3, 0x75, 0xe4, 0x33, 0x5a, 0x98, 0xec, 0xc4, 0x67, 0x69, 0x94, 0x39, 0xfe, 0xdb, 0xeb,
	0x55, 0x8f, 0x84, 0xd0, 0xc2, 0x10, 0x72, 0x9a, 0x20, 0x24, 0x8d, 0x08, 0x71, 0x33, 0x6f, 0xff,
	0x5d, 0xb0, 0xeb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x6d, 0x90, 0x4b, 0xf3, 0x01, 0x00,
	0x00,
}
