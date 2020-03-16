// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/user_location_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A user location view.
//
// User Location View includes all metrics aggregated at the country level,
// one row per country. It reports metrics at the actual physical location of
// the user by targeted or not targeted location. If other segment fields are
// used, you may get more than one row per country.
type UserLocationView struct {
	// Output only. The resource name of the user location view.
	// UserLocation view resource names have the form:
	//
	// `customers/{customer_id}/userLocationViews/{country_criterion_id}~{targeting_location}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Criterion Id for the country.
	CountryCriterionId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=country_criterion_id,json=countryCriterionId,proto3" json:"country_criterion_id,omitempty"`
	// Output only. Indicates whether location was targeted or not.
	TargetingLocation    *wrappers.BoolValue `protobuf:"bytes,3,opt,name=targeting_location,json=targetingLocation,proto3" json:"targeting_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserLocationView) Reset()         { *m = UserLocationView{} }
func (m *UserLocationView) String() string { return proto.CompactTextString(m) }
func (*UserLocationView) ProtoMessage()    {}
func (*UserLocationView) Descriptor() ([]byte, []int) {
	return fileDescriptor_010a04e760a368f3, []int{0}
}

func (m *UserLocationView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLocationView.Unmarshal(m, b)
}
func (m *UserLocationView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLocationView.Marshal(b, m, deterministic)
}
func (m *UserLocationView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLocationView.Merge(m, src)
}
func (m *UserLocationView) XXX_Size() int {
	return xxx_messageInfo_UserLocationView.Size(m)
}
func (m *UserLocationView) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLocationView.DiscardUnknown(m)
}

var xxx_messageInfo_UserLocationView proto.InternalMessageInfo

func (m *UserLocationView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserLocationView) GetCountryCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CountryCriterionId
	}
	return nil
}

func (m *UserLocationView) GetTargetingLocation() *wrappers.BoolValue {
	if m != nil {
		return m.TargetingLocation
	}
	return nil
}

func init() {
	proto.RegisterType((*UserLocationView)(nil), "google.ads.googleads.v3.resources.UserLocationView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/user_location_view.proto", fileDescriptor_010a04e760a368f3)
}

var fileDescriptor_010a04e760a368f3 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x6a, 0xd4, 0x40,
	0x1c, 0x66, 0x13, 0x10, 0x8c, 0x0a, 0x1a, 0x14, 0xd6, 0x55, 0xb4, 0x15, 0x0a, 0x15, 0x61, 0x06,
	0x8d, 0x78, 0x98, 0x9e, 0x26, 0x1e, 0x4a, 0x45, 0x4a, 0x59, 0x30, 0x07, 0x59, 0x08, 0xb3, 0xc9,
	0x34, 0x1d, 0x4c, 0xe6, 0x17, 0x66, 0x26, 0xbb, 0x48, 0xe9, 0xc5, 0x47, 0xf1, 0xe8, 0xa3, 0xf8,
	0x14, 0x3d, 0xf7, 0x11, 0xf4, 0x22, 0x9b, 0xcc, 0x4c, 0x43, 0x17, 0x14, 0x6f, 0x5f, 0xf8, 0xfe,
	0xe4, 0xfb, 0x86, 0x5f, 0x44, 0x2a, 0x80, 0xaa, 0xe6, 0x98, 0x95, 0x1a, 0x0f, 0x70, 0x83, 0x56,
	0x09, 0x56, 0x5c, 0x43, 0xa7, 0x0a, 0xae, 0x71, 0xa7, 0xb9, 0xca, 0x6b, 0x28, 0x98, 0x11, 0x20,
	0xf3, 0x95, 0xe0, 0x6b, 0xd4, 0x2a, 0x30, 0x10, 0xef, 0x0e, 0x06, 0xc4, 0x4a, 0x8d, 0xbc, 0x17,
	0xad, 0x12, 0xe4, 0xbd, 0xb3, 0xe7, 0x2e, 0xbe, 0x15, 0xf8, 0x54, 0xf0, 0xba, 0xcc, 0x97, 0xfc,
	0x8c, 0xad, 0x04, 0xa8, 0x21, 0x63, 0xf6, 0x78, 0x24, 0x70, 0x36, 0x4b, 0x3d, 0xb3, 0x54, 0xff,
	0xb5, 0xec, 0x4e, 0xf1, 0x5a, 0xb1, 0xb6, 0xe5, 0x4a, 0x5b, 0xfe, 0xe9, 0xc8, 0xca, 0xa4, 0x04,
	0xd3, 0x17, 0xb4, 0xec, 0x8b, 0xdf, 0x41, 0x74, 0xff, 0x93, 0xe6, 0xea, 0xa3, 0x2d, 0x9e, 0x09,
	0xbe, 0x8e, 0xb3, 0xe8, 0x9e, 0xfb, 0x49, 0x2e, 0x59, 0xc3, 0xa7, 0x93, 0x9d, 0xc9, 0xfe, 0xed,
	0xf4, 0xf5, 0x25, 0x0d, 0x7f, 0xd1, 0x57, 0xd1, 0xcb, 0xeb, 0x15, 0x16, 0xb5, 0x42, 0xa3, 0x02,
	0x1a, 0x7c, 0x33, 0x69, 0x7e, 0xd7, 0xe5, 0x1c, 0xb3, 0x86, 0xc7, 0xf3, 0xe8, 0x61, 0x01, 0x9d,
	0x34, 0xea, 0x6b, 0x5e, 0x28, 0x61, 0xb8, 0xda, 0xbc, 0x94, 0x28, 0xa7, 0xc1, 0xce, 0x64, 0xff,
	0xce, 0x9b, 0x27, 0x36, 0x0d, 0xb9, 0x25, 0xe8, 0x48, 0x9a, 0x77, 0x6f, 0x33, 0x56, 0x77, 0x3c,
	0x0d, 0x2f, 0x69, 0x38, 0x8f, 0xad, 0xfb, 0xbd, 0x33, 0x1f, 0x95, 0xf1, 0x71, 0x14, 0x1b, 0xa6,
	0x2a, 0x6e, 0x84, 0xac, 0xfc, 0xf3, 0x4f, 0xc3, 0x3e, 0x71, 0xb6, 0x95, 0x98, 0x02, 0xd4, 0xa3,
	0xc0, 0x07, 0xde, 0xea, 0x5a, 0x93, 0x2f, 0x57, 0xf4, 0xec, 0x3f, 0x16, 0xc6, 0x07, 0x45, 0xa7,
	0x0d, 0x34, 0x5c, 0x69, 0x7c, 0xee, 0xe0, 0x45, 0x7f, 0x0c, 0x63, 0x99, 0xc6, 0xe7, 0xdb, 0xf7,
	0x71, 0x91, 0x7e, 0x0b, 0xa2, 0xbd, 0x02, 0x1a, 0xf4, 0xcf, 0x0b, 0x49, 0x1f, 0xdd, 0xfc, 0xf1,
	0xc9, 0x66, 0xd2, 0xc9, 0xe4, 0xf3, 0x07, 0xeb, 0xad, 0xa0, 0x66, 0xb2, 0x42, 0xa0, 0x2a, 0x5c,
	0x71, 0xd9, 0x0f, 0xc6, 0xd7, 0xdd, 0xff, 0x72, 0xb8, 0x07, 0x1e, 0x7d, 0x0f, 0xc2, 0x43, 0x4a,
	0x7f, 0x04, 0xbb, 0x87, 0x43, 0x24, 0x2d, 0x35, 0x1a, 0xe0, 0x06, 0x65, 0x09, 0x9a, 0x3b, 0xe5,
	0x4f, 0xa7, 0x59, 0xd0, 0x52, 0x2f, 0xbc, 0x66, 0x91, 0x25, 0x0b, 0xaf, 0xb9, 0x0a, 0xf6, 0x06,
	0x82, 0x10, 0x5a, 0x6a, 0x42, 0xbc, 0x8a, 0x90, 0x2c, 0x21, 0xc4, 0xeb, 0x96, 0xb7, 0xfa, 0xb2,
	0xc9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0xf9, 0xc2, 0x8f, 0x64, 0x03, 0x00, 0x00,
}
