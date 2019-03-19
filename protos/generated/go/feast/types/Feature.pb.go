// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/types/Feature.proto

package types // import "github.com/gojek/feast/protos/generated/go/feast/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Feature struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                *Value   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_Feature_9650e908dedbbf49, []int{0}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Feature) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Feature)(nil), "feast.types.Feature")
}

func init() { proto.RegisterFile("feast/types/Feature.proto", fileDescriptor_Feature_9650e908dedbbf49) }

var fileDescriptor_Feature_9650e908dedbbf49 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4b, 0x4d, 0x2c,
	0x2e, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x77, 0x4b, 0x4d, 0x2c, 0x29, 0x2d, 0x4a, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x4b, 0xe9, 0x81, 0xa5, 0xa4, 0xc4, 0x91, 0xd5,
	0x85, 0x25, 0xe6, 0x94, 0x42, 0x55, 0x29, 0x39, 0x73, 0xb1, 0x43, 0xb5, 0x09, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x69, 0x70, 0xb1,
	0x96, 0x81, 0x54, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xe9, 0x21, 0x19, 0xa8, 0x07,
	0x36, 0x23, 0x08, 0xa2, 0xc0, 0x29, 0x98, 0x0b, 0xd9, 0x32, 0x27, 0x1e, 0xa8, 0x89, 0x01, 0x20,
	0x1b, 0xa2, 0xcc, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3,
	0xb3, 0x52, 0xb3, 0xf5, 0x21, 0x6e, 0x01, 0xdb, 0x5f, 0xac, 0x9f, 0x9e, 0x9a, 0x97, 0x5a, 0x94,
	0x58, 0x92, 0x9a, 0xa2, 0x9f, 0x9e, 0xaf, 0x8f, 0xe4, 0xca, 0x24, 0x36, 0xb0, 0x02, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x5f, 0x2d, 0xd9, 0xe3, 0x00, 0x00, 0x00,
}
