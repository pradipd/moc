// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_virtualnetworkinterface.proto

package network

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VirtualNetworkInterface_NetworkInterfaceType int32

const (
	VirtualNetworkInterface_Local  VirtualNetworkInterface_NetworkInterfaceType = 0
	VirtualNetworkInterface_Remote VirtualNetworkInterface_NetworkInterfaceType = 1
)

var VirtualNetworkInterface_NetworkInterfaceType_name = map[int32]string{
	0: "Local",
	1: "Remote",
}

var VirtualNetworkInterface_NetworkInterfaceType_value = map[string]int32{
	"Local":  0,
	"Remote": 1,
}

func (x VirtualNetworkInterface_NetworkInterfaceType) String() string {
	return proto.EnumName(VirtualNetworkInterface_NetworkInterfaceType_name, int32(x))
}

func (VirtualNetworkInterface_NetworkInterfaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_031864a4497d6a65, []int{3, 0}
}

type VirtualNetworkInterfaceRequest struct {
	VirtualNetworkInterfaces []*VirtualNetworkInterface `protobuf:"bytes,1,rep,name=VirtualNetworkInterfaces,proto3" json:"VirtualNetworkInterfaces,omitempty"`
	OperationType            common.Operation           `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *VirtualNetworkInterfaceRequest) Reset()         { *m = VirtualNetworkInterfaceRequest{} }
func (m *VirtualNetworkInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkInterfaceRequest) ProtoMessage()    {}
func (*VirtualNetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_031864a4497d6a65, []int{0}
}

func (m *VirtualNetworkInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkInterfaceRequest.Unmarshal(m, b)
}
func (m *VirtualNetworkInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkInterfaceRequest.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkInterfaceRequest.Merge(m, src)
}
func (m *VirtualNetworkInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkInterfaceRequest.Size(m)
}
func (m *VirtualNetworkInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkInterfaceRequest proto.InternalMessageInfo

func (m *VirtualNetworkInterfaceRequest) GetVirtualNetworkInterfaces() []*VirtualNetworkInterface {
	if m != nil {
		return m.VirtualNetworkInterfaces
	}
	return nil
}

func (m *VirtualNetworkInterfaceRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualNetworkInterfaceResponse struct {
	VirtualNetworkInterfaces []*VirtualNetworkInterface `protobuf:"bytes,1,rep,name=VirtualNetworkInterfaces,proto3" json:"VirtualNetworkInterfaces,omitempty"`
	Result                   *wrappers.BoolValue        `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                    string                     `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *VirtualNetworkInterfaceResponse) Reset()         { *m = VirtualNetworkInterfaceResponse{} }
func (m *VirtualNetworkInterfaceResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkInterfaceResponse) ProtoMessage()    {}
func (*VirtualNetworkInterfaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_031864a4497d6a65, []int{1}
}

func (m *VirtualNetworkInterfaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkInterfaceResponse.Unmarshal(m, b)
}
func (m *VirtualNetworkInterfaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkInterfaceResponse.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkInterfaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkInterfaceResponse.Merge(m, src)
}
func (m *VirtualNetworkInterfaceResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkInterfaceResponse.Size(m)
}
func (m *VirtualNetworkInterfaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkInterfaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkInterfaceResponse proto.InternalMessageInfo

func (m *VirtualNetworkInterfaceResponse) GetVirtualNetworkInterfaces() []*VirtualNetworkInterface {
	if m != nil {
		return m.VirtualNetworkInterfaces
	}
	return nil
}

func (m *VirtualNetworkInterfaceResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualNetworkInterfaceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type IpConfiguration struct {
	Ipaddress            string                    `protobuf:"bytes,1,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	Prefixlength         string                    `protobuf:"bytes,2,opt,name=prefixlength,proto3" json:"prefixlength,omitempty"`
	Subnetid             string                    `protobuf:"bytes,3,opt,name=subnetid,proto3" json:"subnetid,omitempty"`
	Primary              bool                      `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Gateway              string                    `protobuf:"bytes,5,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Allocation           common.IPAllocationMethod `protobuf:"varint,6,opt,name=allocation,proto3,enum=moc.IPAllocationMethod" json:"allocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *IpConfiguration) Reset()         { *m = IpConfiguration{} }
func (m *IpConfiguration) String() string { return proto.CompactTextString(m) }
func (*IpConfiguration) ProtoMessage()    {}
func (*IpConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_031864a4497d6a65, []int{2}
}

func (m *IpConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpConfiguration.Unmarshal(m, b)
}
func (m *IpConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpConfiguration.Marshal(b, m, deterministic)
}
func (m *IpConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpConfiguration.Merge(m, src)
}
func (m *IpConfiguration) XXX_Size() int {
	return xxx_messageInfo_IpConfiguration.Size(m)
}
func (m *IpConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_IpConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_IpConfiguration proto.InternalMessageInfo

func (m *IpConfiguration) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func (m *IpConfiguration) GetPrefixlength() string {
	if m != nil {
		return m.Prefixlength
	}
	return ""
}

func (m *IpConfiguration) GetSubnetid() string {
	if m != nil {
		return m.Subnetid
	}
	return ""
}

func (m *IpConfiguration) GetPrimary() bool {
	if m != nil {
		return m.Primary
	}
	return false
}

func (m *IpConfiguration) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *IpConfiguration) GetAllocation() common.IPAllocationMethod {
	if m != nil {
		return m.Allocation
	}
	return common.IPAllocationMethod_Invalid
}

type VirtualNetworkInterface struct {
	Name                 string                                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                                       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 VirtualNetworkInterface_NetworkInterfaceType `protobuf:"varint,3,opt,name=type,proto3,enum=moc.nodeagent.network.VirtualNetworkInterface_NetworkInterfaceType" json:"type,omitempty"`
	Ipconfigs            []*IpConfiguration                           `protobuf:"bytes,4,rep,name=ipconfigs,proto3" json:"ipconfigs,omitempty"`
	Macaddress           string                                       `protobuf:"bytes,5,opt,name=macaddress,proto3" json:"macaddress,omitempty"`
	DnsSettings          *common.Dns                                  `protobuf:"bytes,6,opt,name=dnsSettings,proto3" json:"dnsSettings,omitempty"`
	VirtualMachineName   string                                       `protobuf:"bytes,7,opt,name=virtualMachineName,proto3" json:"virtualMachineName,omitempty"`
	Status               *common.Status                               `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity                               `protobuf:"bytes,9,opt,name=entity,proto3" json:"entity,omitempty"`
	IovWeight            uint32                                       `protobuf:"varint,10,opt,name=iovWeight,proto3" json:"iovWeight,omitempty"`
	Tags                 *common.Tags                                 `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *VirtualNetworkInterface) Reset()         { *m = VirtualNetworkInterface{} }
func (m *VirtualNetworkInterface) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkInterface) ProtoMessage()    {}
func (*VirtualNetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_031864a4497d6a65, []int{3}
}

func (m *VirtualNetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkInterface.Unmarshal(m, b)
}
func (m *VirtualNetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkInterface.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkInterface.Merge(m, src)
}
func (m *VirtualNetworkInterface) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkInterface.Size(m)
}
func (m *VirtualNetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkInterface proto.InternalMessageInfo

func (m *VirtualNetworkInterface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualNetworkInterface) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualNetworkInterface) GetType() VirtualNetworkInterface_NetworkInterfaceType {
	if m != nil {
		return m.Type
	}
	return VirtualNetworkInterface_Local
}

func (m *VirtualNetworkInterface) GetIpconfigs() []*IpConfiguration {
	if m != nil {
		return m.Ipconfigs
	}
	return nil
}

func (m *VirtualNetworkInterface) GetMacaddress() string {
	if m != nil {
		return m.Macaddress
	}
	return ""
}

func (m *VirtualNetworkInterface) GetDnsSettings() *common.Dns {
	if m != nil {
		return m.DnsSettings
	}
	return nil
}

func (m *VirtualNetworkInterface) GetVirtualMachineName() string {
	if m != nil {
		return m.VirtualMachineName
	}
	return ""
}

func (m *VirtualNetworkInterface) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualNetworkInterface) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *VirtualNetworkInterface) GetIovWeight() uint32 {
	if m != nil {
		return m.IovWeight
	}
	return 0
}

func (m *VirtualNetworkInterface) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.nodeagent.network.VirtualNetworkInterface_NetworkInterfaceType", VirtualNetworkInterface_NetworkInterfaceType_name, VirtualNetworkInterface_NetworkInterfaceType_value)
	proto.RegisterType((*VirtualNetworkInterfaceRequest)(nil), "moc.nodeagent.network.VirtualNetworkInterfaceRequest")
	proto.RegisterType((*VirtualNetworkInterfaceResponse)(nil), "moc.nodeagent.network.VirtualNetworkInterfaceResponse")
	proto.RegisterType((*IpConfiguration)(nil), "moc.nodeagent.network.IpConfiguration")
	proto.RegisterType((*VirtualNetworkInterface)(nil), "moc.nodeagent.network.VirtualNetworkInterface")
}

func init() {
	proto.RegisterFile("moc_nodeagent_virtualnetworkinterface.proto", fileDescriptor_031864a4497d6a65)
}

var fileDescriptor_031864a4497d6a65 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xc6, 0x10, 0x42, 0x72, 0xb2, 0xb0, 0xec, 0x88, 0x5d, 0xbc, 0x59, 0xc8, 0x46, 0x59, 0x69,
	0x15, 0xb5, 0xaa, 0x53, 0xa5, 0x7f, 0xd7, 0xfc, 0x49, 0x8d, 0x54, 0x68, 0x35, 0x20, 0x90, 0x7a,
	0x83, 0x26, 0xf6, 0x89, 0x33, 0xc5, 0x9e, 0x71, 0x67, 0xc6, 0xd0, 0xbc, 0x44, 0x1f, 0xaa, 0x2f,
	0xd1, 0x9b, 0xde, 0xf5, 0x45, 0x2a, 0x8f, 0x9d, 0xc4, 0x50, 0x72, 0xc1, 0x4d, 0xaf, 0x92, 0x39,
	0xdf, 0x77, 0xbe, 0xf3, 0xeb, 0x03, 0x8f, 0x63, 0xe9, 0x5f, 0x0a, 0x19, 0x20, 0x0b, 0x51, 0x98,
	0xcb, 0x6b, 0xae, 0x4c, 0xca, 0x22, 0x81, 0xe6, 0x46, 0xaa, 0x2b, 0x2e, 0x0c, 0xaa, 0x11, 0xf3,
	0xd1, 0x4b, 0x94, 0x34, 0x92, 0xfc, 0x19, 0x4b, 0xdf, 0x9b, 0x91, 0xbd, 0x82, 0xd5, 0x6c, 0x85,
	0x52, 0x86, 0x11, 0xf6, 0x2c, 0x69, 0x98, 0x8e, 0x7a, 0x37, 0x8a, 0x25, 0x09, 0x2a, 0x9d, 0xbb,
	0x35, 0xff, 0xb9, 0x8b, 0x63, 0x9c, 0x98, 0x49, 0x01, 0x6e, 0x67, 0x09, 0xf8, 0x32, 0x8e, 0xa5,
	0x28, 0x7e, 0x0a, 0x60, 0xb7, 0x04, 0x08, 0x69, 0xf8, 0x88, 0xfb, 0xcc, 0xf0, 0x19, 0xdc, 0x2a,
	0xc3, 0x79, 0x22, 0x65, 0xf7, 0xce, 0x17, 0x07, 0x5a, 0xe7, 0x79, 0x35, 0x27, 0x39, 0x3c, 0x98,
	0x56, 0x43, 0xf1, 0x63, 0x8a, 0xda, 0x90, 0x0f, 0xe0, 0x2e, 0x60, 0x68, 0xd7, 0x69, 0xaf, 0x74,
	0x1b, 0x7d, 0xcf, 0xbb, 0xb7, 0x62, 0x6f, 0x91, 0xf0, 0x42, 0x3d, 0xf2, 0x1c, 0xd6, 0xdf, 0x26,
	0xa8, 0x6c, 0x05, 0x67, 0x93, 0x04, 0xdd, 0xe5, 0xb6, 0xd3, 0xdd, 0xe8, 0x6f, 0xd8, 0x00, 0x33,
	0x84, 0xde, 0x26, 0x75, 0xbe, 0x3a, 0xf0, 0xef, 0xc2, 0x22, 0x74, 0x22, 0x85, 0xc6, 0x5f, 0x5a,
	0x45, 0x1f, 0xaa, 0x14, 0x75, 0x1a, 0x19, 0x9b, 0x7e, 0xa3, 0xdf, 0xf4, 0xf2, 0xd1, 0x7a, 0xd3,
	0xd1, 0x7a, 0xfb, 0x52, 0x46, 0xe7, 0x2c, 0x4a, 0x91, 0x16, 0x4c, 0xb2, 0x05, 0xab, 0x47, 0x4a,
	0x49, 0xe5, 0xae, 0xb4, 0x9d, 0x6e, 0x9d, 0xe6, 0x8f, 0xce, 0x37, 0x07, 0x7e, 0x1f, 0x24, 0x07,
	0x52, 0x8c, 0x78, 0x98, 0xe6, 0x15, 0x93, 0x1d, 0xa8, 0xf3, 0x84, 0x05, 0x81, 0x42, 0x9d, 0xa5,
	0x9e, 0xb1, 0xe7, 0x06, 0xd2, 0x81, 0xdf, 0x12, 0x85, 0x23, 0xfe, 0x29, 0x42, 0x11, 0x9a, 0xb1,
	0xcd, 0xa0, 0x4e, 0x6f, 0xd9, 0x48, 0x13, 0x6a, 0x3a, 0x1d, 0x0a, 0x34, 0x3c, 0x28, 0xc2, 0xcd,
	0xde, 0xc4, 0x85, 0xb5, 0x44, 0xf1, 0x98, 0xa9, 0x89, 0x5b, 0x69, 0x3b, 0xdd, 0x1a, 0x9d, 0x3e,
	0x33, 0x24, 0x64, 0x06, 0x6f, 0xd8, 0xc4, 0x5d, 0xb5, 0x4e, 0xd3, 0x27, 0x79, 0x05, 0xc0, 0xa2,
	0x48, 0xe6, 0x8b, 0xe7, 0x56, 0xed, 0xc8, 0xb6, 0x6d, 0x37, 0x07, 0xef, 0xf6, 0x66, 0xc0, 0x31,
	0x9a, 0xb1, 0x0c, 0x68, 0x89, 0xda, 0xf9, 0x5c, 0x81, 0xed, 0x05, 0x5d, 0x24, 0x04, 0x2a, 0x82,
	0xc5, 0x58, 0x54, 0x68, 0xff, 0x93, 0x0d, 0x58, 0xe6, 0x41, 0x51, 0xd2, 0x32, 0x0f, 0xc8, 0x05,
	0x54, 0x4c, 0xb6, 0x25, 0x2b, 0x36, 0xe4, 0xc1, 0xc3, 0x06, 0xe8, 0xdd, 0x35, 0x64, 0xbb, 0x44,
	0xad, 0x20, 0x39, 0xcc, 0x7a, 0xec, 0xdb, 0xb6, 0x6b, 0xb7, 0x62, 0xd7, 0xe3, 0xff, 0x05, 0xea,
	0x77, 0xc6, 0x43, 0xe7, 0x8e, 0xa4, 0x05, 0x10, 0x33, 0x7f, 0x3a, 0xaa, 0xbc, 0x69, 0x25, 0x0b,
	0x79, 0x04, 0x8d, 0x40, 0xe8, 0x53, 0x34, 0x86, 0x8b, 0x50, 0xdb, 0xc6, 0x35, 0xfa, 0x35, 0x1b,
	0xe7, 0x50, 0x68, 0x5a, 0x06, 0x89, 0x07, 0xa4, 0xb8, 0x3a, 0xc7, 0xcc, 0x1f, 0x73, 0x81, 0x27,
	0x59, 0x73, 0xd6, 0xac, 0xe6, 0x3d, 0x08, 0xf9, 0x0f, 0xaa, 0xda, 0x30, 0x93, 0x6a, 0xb7, 0x66,
	0x65, 0x1b, 0x56, 0xf6, 0xd4, 0x9a, 0x68, 0x01, 0x65, 0x24, 0x14, 0x86, 0x9b, 0x89, 0x5b, 0x2f,
	0x91, 0x8e, 0xac, 0x89, 0x16, 0x90, 0xdd, 0x37, 0x79, 0x7d, 0x81, 0x3c, 0x1c, 0x1b, 0x17, 0xda,
	0x4e, 0x77, 0x9d, 0xce, 0x0d, 0x64, 0x17, 0x2a, 0x86, 0x85, 0xda, 0x6d, 0x58, 0x81, 0xba, 0x15,
	0x38, 0x63, 0xa1, 0xa6, 0xd6, 0xdc, 0x79, 0x02, 0x5b, 0xf7, 0xb5, 0x99, 0xd4, 0x61, 0xf5, 0x8d,
	0xf4, 0x59, 0xb4, 0xb9, 0x44, 0x20, 0xfb, 0x5a, 0x62, 0x69, 0x70, 0xd3, 0xe9, 0x7f, 0x77, 0x60,
	0x67, 0xc1, 0xb8, 0xf6, 0xb2, 0xae, 0x93, 0x09, 0x54, 0x07, 0xe2, 0x5a, 0x5e, 0x21, 0x79, 0xf1,
	0xc0, 0xcf, 0x35, 0xbf, 0x66, 0xcd, 0x97, 0x0f, 0x75, 0xcb, 0xef, 0x47, 0x67, 0x89, 0xbc, 0x86,
	0x3f, 0x0e, 0xc6, 0xe8, 0x5f, 0x9d, 0x94, 0xae, 0x2c, 0xf9, 0xeb, 0xa7, 0x4f, 0xfb, 0x28, 0xbb,
	0xda, 0xcd, 0xbf, 0x6d, 0x98, 0x32, 0x75, 0xae, 0xb4, 0xff, 0xf4, 0xbd, 0x17, 0x72, 0x33, 0x4e,
	0x87, 0x9e, 0x2f, 0xe3, 0x5e, 0xcc, 0x7d, 0x25, 0xb5, 0x1c, 0x99, 0x5e, 0x2c, 0xfd, 0x9e, 0x4a,
	0xfc, 0xde, 0x2c, 0xbb, 0x5e, 0x91, 0xdd, 0xb0, 0x6a, 0xe5, 0x9f, 0xfd, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x27, 0x6c, 0x0f, 0x85, 0x88, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualNetworkInterfaceAgentClient is the client API for VirtualNetworkInterfaceAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualNetworkInterfaceAgentClient interface {
	Invoke(ctx context.Context, in *VirtualNetworkInterfaceRequest, opts ...grpc.CallOption) (*VirtualNetworkInterfaceResponse, error)
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
}

type virtualNetworkInterfaceAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualNetworkInterfaceAgentClient(cc *grpc.ClientConn) VirtualNetworkInterfaceAgentClient {
	return &virtualNetworkInterfaceAgentClient{cc}
}

func (c *virtualNetworkInterfaceAgentClient) Invoke(ctx context.Context, in *VirtualNetworkInterfaceRequest, opts ...grpc.CallOption) (*VirtualNetworkInterfaceResponse, error) {
	out := new(VirtualNetworkInterfaceResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.network.VirtualNetworkInterfaceAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualNetworkInterfaceAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.network.VirtualNetworkInterfaceAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualNetworkInterfaceAgentServer is the server API for VirtualNetworkInterfaceAgent service.
type VirtualNetworkInterfaceAgentServer interface {
	Invoke(context.Context, *VirtualNetworkInterfaceRequest) (*VirtualNetworkInterfaceResponse, error)
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
}

// UnimplementedVirtualNetworkInterfaceAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualNetworkInterfaceAgentServer struct {
}

func (*UnimplementedVirtualNetworkInterfaceAgentServer) Invoke(ctx context.Context, req *VirtualNetworkInterfaceRequest) (*VirtualNetworkInterfaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedVirtualNetworkInterfaceAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}

func RegisterVirtualNetworkInterfaceAgentServer(s *grpc.Server, srv VirtualNetworkInterfaceAgentServer) {
	s.RegisterService(&_VirtualNetworkInterfaceAgent_serviceDesc, srv)
}

func _VirtualNetworkInterfaceAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualNetworkInterfaceAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.network.VirtualNetworkInterfaceAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualNetworkInterfaceAgentServer).Invoke(ctx, req.(*VirtualNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualNetworkInterfaceAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualNetworkInterfaceAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.network.VirtualNetworkInterfaceAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualNetworkInterfaceAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualNetworkInterfaceAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.network.VirtualNetworkInterfaceAgent",
	HandlerType: (*VirtualNetworkInterfaceAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualNetworkInterfaceAgent_Invoke_Handler,
		},
		{
			MethodName: "CheckNotification",
			Handler:    _VirtualNetworkInterfaceAgent_CheckNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_virtualnetworkinterface.proto",
}
