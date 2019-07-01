// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packages.proto

package proto

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

type Package struct {
	Id                   string           `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Sku                  string           `protobuf:"bytes,2,opt,name=Sku,proto3" json:"Sku,omitempty"`
	Name                 *LocalizedString `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Image                *LocalizedString `protobuf:"bytes,4,opt,name=Image,proto3" json:"Image,omitempty"`
	ImageCover           *LocalizedString `protobuf:"bytes,5,opt,name=ImageCover,proto3" json:"ImageCover,omitempty"`
	ImageThumb           *LocalizedString `protobuf:"bytes,6,opt,name=ImageThumb,proto3" json:"ImageThumb,omitempty"`
	IsUpgradeAllowed     bool             `protobuf:"varint,7,opt,name=IsUpgradeAllowed,proto3" json:"IsUpgradeAllowed,omitempty"`
	IsEnabled            bool             `protobuf:"varint,8,opt,name=IsEnabled,proto3" json:"IsEnabled,omitempty"`
	DefaultProductId     string           `protobuf:"bytes,9,opt,name=DefaultProductId,proto3" json:"DefaultProductId,omitempty"`
	VendorId             string           `protobuf:"bytes,10,opt,name=VendorId,proto3" json:"VendorId,omitempty"`
	CreatorId            string           `protobuf:"bytes,11,opt,name=CreatorId,proto3" json:"CreatorId,omitempty"`
	Discount             uint32           `protobuf:"varint,12,opt,name=Discount,proto3" json:"Discount,omitempty"`
	DiscountOption       uint32           `protobuf:"varint,13,opt,name=DiscountOption,proto3" json:"DiscountOption,omitempty"`
	AllowedCountries     []string         `protobuf:"bytes,14,rep,name=AllowedCountries,proto3" json:"AllowedCountries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_d056f5e793167e97, []int{0}
}

func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Package) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *Package) GetName() *LocalizedString {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Package) GetImage() *LocalizedString {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Package) GetImageCover() *LocalizedString {
	if m != nil {
		return m.ImageCover
	}
	return nil
}

func (m *Package) GetImageThumb() *LocalizedString {
	if m != nil {
		return m.ImageThumb
	}
	return nil
}

func (m *Package) GetIsUpgradeAllowed() bool {
	if m != nil {
		return m.IsUpgradeAllowed
	}
	return false
}

func (m *Package) GetIsEnabled() bool {
	if m != nil {
		return m.IsEnabled
	}
	return false
}

func (m *Package) GetDefaultProductId() string {
	if m != nil {
		return m.DefaultProductId
	}
	return ""
}

func (m *Package) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *Package) GetCreatorId() string {
	if m != nil {
		return m.CreatorId
	}
	return ""
}

func (m *Package) GetDiscount() uint32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *Package) GetDiscountOption() uint32 {
	if m != nil {
		return m.DiscountOption
	}
	return 0
}

func (m *Package) GetAllowedCountries() []string {
	if m != nil {
		return m.AllowedCountries
	}
	return nil
}

func init() {
	proto.RegisterType((*Package)(nil), "proto.Package")
}

func init() { proto.RegisterFile("packages.proto", fileDescriptor_d056f5e793167e97) }

var fileDescriptor_d056f5e793167e97 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0x87, 0xe9, 0xfe, 0xef, 0x6c, 0x2b, 0x23, 0x17, 0x2f, 0x87, 0xf1, 0x5e, 0x14, 0x2f, 0xa4,
	0x0c, 0xd9, 0x85, 0x82, 0xf7, 0xb2, 0x79, 0x11, 0x10, 0x1d, 0x9d, 0x7a, 0x9f, 0x35, 0xb1, 0x96,
	0xb5, 0x49, 0x49, 0x53, 0x05, 0xbf, 0x87, 0xdf, 0x57, 0x92, 0x6a, 0x27, 0x1b, 0x0c, 0xaf, 0x7a,
	0xf2, 0xe4, 0xf9, 0x9d, 0x9e, 0x1c, 0xf0, 0x0b, 0x16, 0xef, 0x58, 0x22, 0xca, 0x45, 0xa1, 0x95,
	0x51, 0xa4, 0xeb, 0x3e, 0xb3, 0x71, 0xac, 0xf2, 0x5c, 0xc9, 0x1a, 0x9e, 0x7d, 0x76, 0xa0, 0xbf,
	0xae, 0x3d, 0xe2, 0x43, 0x8b, 0x72, 0xf4, 0x02, 0x2f, 0x1c, 0x46, 0x2d, 0xca, 0xc9, 0x14, 0xda,
	0x9b, 0x5d, 0x85, 0x2d, 0x07, 0x6c, 0x49, 0xe6, 0xd0, 0xb9, 0x67, 0xb9, 0xc0, 0x76, 0xe0, 0x85,
	0xa3, 0xcb, 0x7f, 0x75, 0x8f, 0xc5, 0x9d, 0x8a, 0x59, 0x96, 0x7e, 0x08, 0xbe, 0x31, 0x3a, 0x95,
	0x49, 0xe4, 0x1c, 0x72, 0x01, 0x5d, 0x9a, 0xb3, 0x44, 0x60, 0xe7, 0xa4, 0x5c, 0x4b, 0xe4, 0x1a,
	0xc0, 0x15, 0x4b, 0xf5, 0x26, 0x34, 0x76, 0x4f, 0x46, 0x7e, 0x99, 0x4d, 0xee, 0xf1, 0xb5, 0xca,
	0xb7, 0xd8, 0xfb, 0x43, 0xce, 0x99, 0x64, 0x0e, 0x53, 0x5a, 0x3e, 0x15, 0x89, 0x66, 0x5c, 0xdc,
	0x64, 0x99, 0x7a, 0x17, 0x1c, 0xfb, 0x81, 0x17, 0x0e, 0xa2, 0x23, 0x4e, 0xfe, 0xc3, 0x90, 0x96,
	0xb7, 0x92, 0x6d, 0x33, 0xc1, 0x71, 0xe0, 0xa4, 0x3d, 0xb0, 0x9d, 0x56, 0xe2, 0x85, 0x55, 0x99,
	0x59, 0x6b, 0xc5, 0xab, 0xd8, 0x50, 0x8e, 0x43, 0xb7, 0xb2, 0x23, 0x4e, 0x66, 0x30, 0x78, 0x16,
	0x92, 0x2b, 0x4d, 0x39, 0x82, 0x73, 0x9a, 0xb3, 0xfd, 0xcb, 0x52, 0x0b, 0x66, 0xdc, 0xe5, 0xc8,
	0x5d, 0xee, 0x81, 0x4d, 0xae, 0xd2, 0x32, 0x56, 0x95, 0x34, 0x38, 0x0e, 0xbc, 0x70, 0x12, 0x35,
	0x67, 0x72, 0x0e, 0xfe, 0x4f, 0xfd, 0x50, 0x98, 0x54, 0x49, 0x9c, 0x38, 0xe3, 0x80, 0xda, 0x49,
	0xbf, 0x9f, 0xb4, 0xb4, 0x54, 0xa7, 0xa2, 0x44, 0x3f, 0x68, 0xdb, 0x49, 0x0f, 0xf9, 0xb6, 0xe7,
	0x56, 0x78, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xf8, 0xbf, 0xcb, 0x45, 0x02, 0x00, 0x00,
}
