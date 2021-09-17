// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: network/network.proto

package network

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	common "github.com/talos-systems/talos/pkg/machinery/api/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddressFamily int32

const (
	AddressFamily_AF_UNSPEC AddressFamily = 0
	AddressFamily_AF_INET   AddressFamily = 2
	AddressFamily_IPV4      AddressFamily = 2
	AddressFamily_AF_INET6  AddressFamily = 10
	AddressFamily_IPV6      AddressFamily = 10
)

// Enum value maps for AddressFamily.
var (
	AddressFamily_name = map[int32]string{
		0: "AF_UNSPEC",
		2: "AF_INET",
		// Duplicate value: 2: "IPV4",
		10: "AF_INET6",
		// Duplicate value: 10: "IPV6",
	}
	AddressFamily_value = map[string]int32{
		"AF_UNSPEC": 0,
		"AF_INET":   2,
		"IPV4":      2,
		"AF_INET6":  10,
		"IPV6":      10,
	}
)

func (x AddressFamily) Enum() *AddressFamily {
	p := new(AddressFamily)
	*p = x
	return p
}

func (x AddressFamily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressFamily) Descriptor() protoreflect.EnumDescriptor {
	return file_network_network_proto_enumTypes[0].Descriptor()
}

func (AddressFamily) Type() protoreflect.EnumType {
	return &file_network_network_proto_enumTypes[0]
}

func (x AddressFamily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressFamily.Descriptor instead.
func (AddressFamily) EnumDescriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{0}
}

type RouteProtocol int32

const (
	RouteProtocol_RTPROT_UNSPEC   RouteProtocol = 0
	RouteProtocol_RTPROT_REDIRECT RouteProtocol = 1  // Route installed by ICMP redirects
	RouteProtocol_RTPROT_KERNEL   RouteProtocol = 2  // Route installed by kernel
	RouteProtocol_RTPROT_BOOT     RouteProtocol = 3  // Route installed during boot
	RouteProtocol_RTPROT_STATIC   RouteProtocol = 4  // Route installed by administrator
	RouteProtocol_RTPROT_GATED    RouteProtocol = 8  // Route installed by gated
	RouteProtocol_RTPROT_RA       RouteProtocol = 9  // Route installed by router advertisement
	RouteProtocol_RTPROT_MRT      RouteProtocol = 10 // Route installed by Merit MRT
	RouteProtocol_RTPROT_ZEBRA    RouteProtocol = 11 // Route installed by Zebra/Quagga
	RouteProtocol_RTPROT_BIRD     RouteProtocol = 12 // Route installed by Bird
	RouteProtocol_RTPROT_DNROUTED RouteProtocol = 13 // Route installed by DECnet routing daemon
	RouteProtocol_RTPROT_XORP     RouteProtocol = 14 // Route installed by XORP
	RouteProtocol_RTPROT_NTK      RouteProtocol = 15 // Route installed by Netsukuku
	RouteProtocol_RTPROT_DHCP     RouteProtocol = 16 // Route installed by DHCP
	RouteProtocol_RTPROT_MROUTED  RouteProtocol = 17 // Route installed by Multicast daemon
	RouteProtocol_RTPROT_BABEL    RouteProtocol = 42 // Route installed by Babel daemon
)

// Enum value maps for RouteProtocol.
var (
	RouteProtocol_name = map[int32]string{
		0:  "RTPROT_UNSPEC",
		1:  "RTPROT_REDIRECT",
		2:  "RTPROT_KERNEL",
		3:  "RTPROT_BOOT",
		4:  "RTPROT_STATIC",
		8:  "RTPROT_GATED",
		9:  "RTPROT_RA",
		10: "RTPROT_MRT",
		11: "RTPROT_ZEBRA",
		12: "RTPROT_BIRD",
		13: "RTPROT_DNROUTED",
		14: "RTPROT_XORP",
		15: "RTPROT_NTK",
		16: "RTPROT_DHCP",
		17: "RTPROT_MROUTED",
		42: "RTPROT_BABEL",
	}
	RouteProtocol_value = map[string]int32{
		"RTPROT_UNSPEC":   0,
		"RTPROT_REDIRECT": 1,
		"RTPROT_KERNEL":   2,
		"RTPROT_BOOT":     3,
		"RTPROT_STATIC":   4,
		"RTPROT_GATED":    8,
		"RTPROT_RA":       9,
		"RTPROT_MRT":      10,
		"RTPROT_ZEBRA":    11,
		"RTPROT_BIRD":     12,
		"RTPROT_DNROUTED": 13,
		"RTPROT_XORP":     14,
		"RTPROT_NTK":      15,
		"RTPROT_DHCP":     16,
		"RTPROT_MROUTED":  17,
		"RTPROT_BABEL":    42,
	}
)

func (x RouteProtocol) Enum() *RouteProtocol {
	p := new(RouteProtocol)
	*p = x
	return p
}

func (x RouteProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_network_network_proto_enumTypes[1].Descriptor()
}

func (RouteProtocol) Type() protoreflect.EnumType {
	return &file_network_network_proto_enumTypes[1]
}

func (x RouteProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteProtocol.Descriptor instead.
func (RouteProtocol) EnumDescriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{1}
}

type InterfaceFlags int32

const (
	InterfaceFlags_FLAG_UNKNOWN        InterfaceFlags = 0
	InterfaceFlags_FLAG_UP             InterfaceFlags = 1
	InterfaceFlags_FLAG_BROADCAST      InterfaceFlags = 2
	InterfaceFlags_FLAG_LOOPBACK       InterfaceFlags = 3
	InterfaceFlags_FLAG_POINT_TO_POINT InterfaceFlags = 4
	InterfaceFlags_FLAG_MULTICAST      InterfaceFlags = 5
)

// Enum value maps for InterfaceFlags.
var (
	InterfaceFlags_name = map[int32]string{
		0: "FLAG_UNKNOWN",
		1: "FLAG_UP",
		2: "FLAG_BROADCAST",
		3: "FLAG_LOOPBACK",
		4: "FLAG_POINT_TO_POINT",
		5: "FLAG_MULTICAST",
	}
	InterfaceFlags_value = map[string]int32{
		"FLAG_UNKNOWN":        0,
		"FLAG_UP":             1,
		"FLAG_BROADCAST":      2,
		"FLAG_LOOPBACK":       3,
		"FLAG_POINT_TO_POINT": 4,
		"FLAG_MULTICAST":      5,
	}
)

func (x InterfaceFlags) Enum() *InterfaceFlags {
	p := new(InterfaceFlags)
	*p = x
	return p
}

func (x InterfaceFlags) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InterfaceFlags) Descriptor() protoreflect.EnumDescriptor {
	return file_network_network_proto_enumTypes[2].Descriptor()
}

func (InterfaceFlags) Type() protoreflect.EnumType {
	return &file_network_network_proto_enumTypes[2]
}

func (x InterfaceFlags) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InterfaceFlags.Descriptor instead.
func (InterfaceFlags) EnumDescriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{2}
}

// The messages message containing the routes.
type RoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Routes `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *RoutesResponse) Reset() {
	*x = RoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutesResponse) ProtoMessage() {}

func (x *RoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutesResponse.ProtoReflect.Descriptor instead.
func (*RoutesResponse) Descriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{0}
}

func (x *RoutesResponse) GetMessages() []*Routes {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Routes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *common.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Routes   []*Route         `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *Routes) Reset() {
	*x = Routes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Routes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Routes) ProtoMessage() {}

func (x *Routes) ProtoReflect() protoreflect.Message {
	mi := &file_network_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Routes.ProtoReflect.Descriptor instead.
func (*Routes) Descriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{1}
}

func (x *Routes) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Routes) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

// The messages message containing a route.
type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interface is the interface over which traffic to this destination should be sent
	Interface string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	// Destination is the network prefix CIDR which this route provides
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Gateway is the gateway address to which traffic to this destination should be sent
	Gateway string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// Metric is the priority of the route, where lower metrics have higher priorities
	Metric uint32 `protobuf:"varint,4,opt,name=metric,proto3" json:"metric,omitempty"`
	// Scope desribes the scope of this route
	Scope uint32 `protobuf:"varint,5,opt,name=scope,proto3" json:"scope,omitempty"`
	// Source is the source prefix CIDR for the route, if one is defined
	Source string `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	// Family is the address family of the route.  Currently, the only options are AF_INET (IPV4) and AF_INET6 (IPV6).
	Family AddressFamily `protobuf:"varint,7,opt,name=family,proto3,enum=network.AddressFamily" json:"family,omitempty"`
	// Protocol is the protocol by which this route came to be in place
	Protocol RouteProtocol `protobuf:"varint,8,opt,name=protocol,proto3,enum=network.RouteProtocol" json:"protocol,omitempty"`
	// Flags indicate any special flags on the route
	Flags uint32 `protobuf:"varint,9,opt,name=flags,proto3" json:"flags,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_network_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{2}
}

func (x *Route) GetInterface() string {
	if x != nil {
		return x.Interface
	}
	return ""
}

func (x *Route) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Route) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *Route) GetMetric() uint32 {
	if x != nil {
		return x.Metric
	}
	return 0
}

func (x *Route) GetScope() uint32 {
	if x != nil {
		return x.Scope
	}
	return 0
}

func (x *Route) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Route) GetFamily() AddressFamily {
	if x != nil {
		return x.Family
	}
	return AddressFamily_AF_UNSPEC
}

func (x *Route) GetProtocol() RouteProtocol {
	if x != nil {
		return x.Protocol
	}
	return RouteProtocol_RTPROT_UNSPEC
}

func (x *Route) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

type InterfacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Interfaces `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *InterfacesResponse) Reset() {
	*x = InterfacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesResponse) ProtoMessage() {}

func (x *InterfacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesResponse.ProtoReflect.Descriptor instead.
func (*InterfacesResponse) Descriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{3}
}

func (x *InterfacesResponse) GetMessages() []*Interfaces {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Interfaces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata   *common.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Interfaces []*Interface     `protobuf:"bytes,2,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
}

func (x *Interfaces) Reset() {
	*x = Interfaces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interfaces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interfaces) ProtoMessage() {}

func (x *Interfaces) ProtoReflect() protoreflect.Message {
	mi := &file_network_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interfaces.ProtoReflect.Descriptor instead.
func (*Interfaces) Descriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{4}
}

func (x *Interfaces) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Interfaces) GetInterfaces() []*Interface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

// Interface represents a net.Interface
type Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index        uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Mtu          uint32         `protobuf:"varint,2,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Name         string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Hardwareaddr string         `protobuf:"bytes,4,opt,name=hardwareaddr,proto3" json:"hardwareaddr,omitempty"`
	Flags        InterfaceFlags `protobuf:"varint,5,opt,name=flags,proto3,enum=network.InterfaceFlags" json:"flags,omitempty"`
	Ipaddress    []string       `protobuf:"bytes,6,rep,name=ipaddress,proto3" json:"ipaddress,omitempty"`
}

func (x *Interface) Reset() {
	*x = Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface) ProtoMessage() {}

func (x *Interface) ProtoReflect() protoreflect.Message {
	mi := &file_network_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface.ProtoReflect.Descriptor instead.
func (*Interface) Descriptor() ([]byte, []int) {
	return file_network_network_proto_rawDescGZIP(), []int{5}
}

func (x *Interface) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Interface) GetMtu() uint32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *Interface) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Interface) GetHardwareaddr() string {
	if x != nil {
		return x.Hardwareaddr
	}
	return ""
}

func (x *Interface) GetFlags() InterfaceFlags {
	if x != nil {
		return x.Flags
	}
	return InterfaceFlags_FLAG_UNKNOWN
}

func (x *Interface) GetIpaddress() []string {
	if x != nil {
		return x.Ipaddress
	}
	return nil
}

var File_network_network_proto protoreflect.FileDescriptor

var file_network_network_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x22, 0x5e, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x22, 0xa1, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x45, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x0a,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x22, 0xb8, 0x01, 0x0a,
	0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d,
	0x74, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61,
	0x72, 0x65, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x61, 0x64, 0x64, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x51, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x46, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x46, 0x5f, 0x49, 0x4e,
	0x45, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x34, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x46, 0x5f, 0x49, 0x4e, 0x45, 0x54, 0x36, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x50, 0x56, 0x36, 0x10, 0x0a, 0x1a, 0x02, 0x10, 0x01, 0x2a, 0xaf, 0x02, 0x0a, 0x0d, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x11, 0x0a, 0x0d,
	0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x4b,
	0x45, 0x52, 0x4e, 0x45, 0x4c, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x54, 0x50, 0x52, 0x4f,
	0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x54, 0x50, 0x52,
	0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x52,
	0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0d, 0x0a,
	0x09, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x52, 0x41, 0x10, 0x09, 0x12, 0x0e, 0x0a, 0x0a,
	0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x4d, 0x52, 0x54, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x5a, 0x45, 0x42, 0x52, 0x41, 0x10, 0x0b, 0x12, 0x0f,
	0x0a, 0x0b, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x42, 0x49, 0x52, 0x44, 0x10, 0x0c, 0x12,
	0x13, 0x0a, 0x0f, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x44, 0x4e, 0x52, 0x4f, 0x55, 0x54,
	0x45, 0x44, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f, 0x58,
	0x4f, 0x52, 0x50, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f,
	0x4e, 0x54, 0x4b, 0x10, 0x0f, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54, 0x5f,
	0x44, 0x48, 0x43, 0x50, 0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x54, 0x50, 0x52, 0x4f, 0x54,
	0x5f, 0x4d, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x44, 0x10, 0x11, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x54,
	0x50, 0x52, 0x4f, 0x54, 0x5f, 0x42, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x2a, 0x2a, 0x83, 0x01, 0x0a,
	0x0e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12,
	0x10, 0x0a, 0x0c, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x55, 0x50, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x4c, 0x4f, 0x4f, 0x50, 0x42,
	0x41, 0x43, 0x4b, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x12,
	0x0a, 0x0e, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x43, 0x41, 0x53, 0x54,
	0x10, 0x05, 0x32, 0x98, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x03, 0x88, 0x02, 0x01, 0x12, 0x46, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c, 0x6f,
	0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_network_network_proto_rawDescOnce sync.Once
	file_network_network_proto_rawDescData = file_network_network_proto_rawDesc
)

func file_network_network_proto_rawDescGZIP() []byte {
	file_network_network_proto_rawDescOnce.Do(func() {
		file_network_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_network_proto_rawDescData)
	})
	return file_network_network_proto_rawDescData
}

var (
	file_network_network_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
	file_network_network_proto_msgTypes  = make([]protoimpl.MessageInfo, 6)
	file_network_network_proto_goTypes   = []interface{}{
		(AddressFamily)(0),         // 0: network.AddressFamily
		(RouteProtocol)(0),         // 1: network.RouteProtocol
		(InterfaceFlags)(0),        // 2: network.InterfaceFlags
		(*RoutesResponse)(nil),     // 3: network.RoutesResponse
		(*Routes)(nil),             // 4: network.Routes
		(*Route)(nil),              // 5: network.Route
		(*InterfacesResponse)(nil), // 6: network.InterfacesResponse
		(*Interfaces)(nil),         // 7: network.Interfaces
		(*Interface)(nil),          // 8: network.Interface
		(*common.Metadata)(nil),    // 9: common.Metadata
		(*emptypb.Empty)(nil),      // 10: google.protobuf.Empty
	}
)

var file_network_network_proto_depIdxs = []int32{
	4,  // 0: network.RoutesResponse.messages:type_name -> network.Routes
	9,  // 1: network.Routes.metadata:type_name -> common.Metadata
	5,  // 2: network.Routes.routes:type_name -> network.Route
	0,  // 3: network.Route.family:type_name -> network.AddressFamily
	1,  // 4: network.Route.protocol:type_name -> network.RouteProtocol
	7,  // 5: network.InterfacesResponse.messages:type_name -> network.Interfaces
	9,  // 6: network.Interfaces.metadata:type_name -> common.Metadata
	8,  // 7: network.Interfaces.interfaces:type_name -> network.Interface
	2,  // 8: network.Interface.flags:type_name -> network.InterfaceFlags
	10, // 9: network.NetworkService.Routes:input_type -> google.protobuf.Empty
	10, // 10: network.NetworkService.Interfaces:input_type -> google.protobuf.Empty
	3,  // 11: network.NetworkService.Routes:output_type -> network.RoutesResponse
	6,  // 12: network.NetworkService.Interfaces:output_type -> network.InterfacesResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_network_network_proto_init() }
func file_network_network_proto_init() {
	if File_network_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Routes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfacesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interfaces); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_network_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_network_network_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_network_proto_goTypes,
		DependencyIndexes: file_network_network_proto_depIdxs,
		EnumInfos:         file_network_network_proto_enumTypes,
		MessageInfos:      file_network_network_proto_msgTypes,
	}.Build()
	File_network_network_proto = out.File
	file_network_network_proto_rawDesc = nil
	file_network_network_proto_goTypes = nil
	file_network_network_proto_depIdxs = nil
}
