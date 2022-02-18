// Routes for wwapi.
// TODO: Try protoc-gen-doc for generating documentation.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: routes.proto

package wwapiv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName []string `protobuf:"bytes,1,rep,name=nodeName,proto3" json:"nodeName,omitempty"`
}

func (x *NodeNames) Reset() {
	*x = NodeNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeNames) ProtoMessage() {}

func (x *NodeNames) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeNames.ProtoReflect.Descriptor instead.
func (*NodeNames) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{0}
}

func (x *NodeNames) GetNodeName() []string {
	if x != nil {
		return x.NodeName
	}
	return nil
}

type NodeField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//string fieldName = 1;
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // TODO: Variable name okay? Also this is weird since it could be a bool.
}

func (x *NodeField) Reset() {
	*x = NodeField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeField) ProtoMessage() {}

func (x *NodeField) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeField.ProtoReflect.Descriptor instead.
func (*NodeField) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{1}
}

func (x *NodeField) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *NodeField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// NetDev is network devices (NICs) on a node.
type NetDev struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device  *NodeField `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Hwaddr  *NodeField `protobuf:"bytes,2,opt,name=hwaddr,proto3" json:"hwaddr,omitempty"`
	Ipaddr  *NodeField `protobuf:"bytes,3,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	Netmask *NodeField `protobuf:"bytes,4,opt,name=netmask,proto3" json:"netmask,omitempty"`
	Gateway *NodeField `protobuf:"bytes,5,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Type    *NodeField `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Onboot  *NodeField `protobuf:"bytes,7,opt,name=onboot,proto3" json:"onboot,omitempty"`
	Default *NodeField `protobuf:"bytes,8,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *NetDev) Reset() {
	*x = NetDev{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetDev) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetDev) ProtoMessage() {}

func (x *NetDev) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetDev.ProtoReflect.Descriptor instead.
func (*NetDev) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{2}
}

func (x *NetDev) GetDevice() *NodeField {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *NetDev) GetHwaddr() *NodeField {
	if x != nil {
		return x.Hwaddr
	}
	return nil
}

func (x *NetDev) GetIpaddr() *NodeField {
	if x != nil {
		return x.Ipaddr
	}
	return nil
}

func (x *NetDev) GetNetmask() *NodeField {
	if x != nil {
		return x.Netmask
	}
	return nil
}

func (x *NetDev) GetGateway() *NodeField {
	if x != nil {
		return x.Gateway
	}
	return nil
}

func (x *NetDev) GetType() *NodeField {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *NetDev) GetOnboot() *NodeField {
	if x != nil {
		return x.Onboot
	}
	return nil
}

func (x *NetDev) GetDefault() *NodeField {
	if x != nil {
		return x.Default
	}
	return nil
}

// TODO: Needed?
type NodeTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*NodeField `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *NodeTags) Reset() {
	*x = NodeTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeTags) ProtoMessage() {}

func (x *NodeTags) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeTags.ProtoReflect.Descriptor instead.
func (*NodeTags) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{3}
}

func (x *NodeTags) GetTags() []*NodeField {
	if x != nil {
		return x.Tags
	}
	return nil
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName       string                `protobuf:"bytes,1,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	Id             *NodeField            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Comment        *NodeField            `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Cluster        *NodeField            `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Profiles       *NodeField            `protobuf:"bytes,5,opt,name=profiles,proto3" json:"profiles,omitempty"`
	Discoverable   *NodeField            `protobuf:"bytes,6,opt,name=discoverable,proto3" json:"discoverable,omitempty"`
	Container      *NodeField            `protobuf:"bytes,7,opt,name=container,proto3" json:"container,omitempty"`
	Kernel         *NodeField            `protobuf:"bytes,8,opt,name=kernel,proto3" json:"kernel,omitempty"`
	KernelArgs     *NodeField            `protobuf:"bytes,9,opt,name=kernelArgs,proto3" json:"kernelArgs,omitempty"`
	SystemOverlay  *NodeField            `protobuf:"bytes,10,opt,name=systemOverlay,proto3" json:"systemOverlay,omitempty"`
	RuntimeOverlay *NodeField            `protobuf:"bytes,11,opt,name=runtimeOverlay,proto3" json:"runtimeOverlay,omitempty"`
	Ipxe           *NodeField            `protobuf:"bytes,12,opt,name=ipxe,proto3" json:"ipxe,omitempty"`
	Init           *NodeField            `protobuf:"bytes,13,opt,name=init,proto3" json:"init,omitempty"`
	Root           *NodeField            `protobuf:"bytes,14,opt,name=root,proto3" json:"root,omitempty"`
	AssetKey       *NodeField            `protobuf:"bytes,15,opt,name=assetKey,proto3" json:"assetKey,omitempty"`
	IpmiIpaddr     *NodeField            `protobuf:"bytes,16,opt,name=ipmiIpaddr,proto3" json:"ipmiIpaddr,omitempty"`
	IpmiNetmask    *NodeField            `protobuf:"bytes,17,opt,name=ipmiNetmask,proto3" json:"ipmiNetmask,omitempty"`
	IpmiPort       *NodeField            `protobuf:"bytes,18,opt,name=ipmiPort,proto3" json:"ipmiPort,omitempty"`
	IpmiGateway    *NodeField            `protobuf:"bytes,19,opt,name=ipmiGateway,proto3" json:"ipmiGateway,omitempty"`
	IpmiUserName   *NodeField            `protobuf:"bytes,20,opt,name=ipmiUserName,proto3" json:"ipmiUserName,omitempty"`
	IpmiInterface  *NodeField            `protobuf:"bytes,21,opt,name=ipmiInterface,proto3" json:"ipmiInterface,omitempty"`
	NetDevs        map[string]*NetDev    `protobuf:"bytes,22,rep,name=NetDevs,proto3" json:"NetDevs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags           map[string]*NodeField `protobuf:"bytes,23,rep,name=Tags,proto3" json:"Tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Keys           map[string]*NodeField `protobuf:"bytes,24,rep,name=Keys,proto3" json:"Keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // TODO: We may not need this. Tags may be it. Ask Greg.
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{4}
}

func (x *NodeInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *NodeInfo) GetId() *NodeField {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *NodeInfo) GetComment() *NodeField {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *NodeInfo) GetCluster() *NodeField {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *NodeInfo) GetProfiles() *NodeField {
	if x != nil {
		return x.Profiles
	}
	return nil
}

func (x *NodeInfo) GetDiscoverable() *NodeField {
	if x != nil {
		return x.Discoverable
	}
	return nil
}

func (x *NodeInfo) GetContainer() *NodeField {
	if x != nil {
		return x.Container
	}
	return nil
}

func (x *NodeInfo) GetKernel() *NodeField {
	if x != nil {
		return x.Kernel
	}
	return nil
}

func (x *NodeInfo) GetKernelArgs() *NodeField {
	if x != nil {
		return x.KernelArgs
	}
	return nil
}

func (x *NodeInfo) GetSystemOverlay() *NodeField {
	if x != nil {
		return x.SystemOverlay
	}
	return nil
}

func (x *NodeInfo) GetRuntimeOverlay() *NodeField {
	if x != nil {
		return x.RuntimeOverlay
	}
	return nil
}

func (x *NodeInfo) GetIpxe() *NodeField {
	if x != nil {
		return x.Ipxe
	}
	return nil
}

func (x *NodeInfo) GetInit() *NodeField {
	if x != nil {
		return x.Init
	}
	return nil
}

func (x *NodeInfo) GetRoot() *NodeField {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *NodeInfo) GetAssetKey() *NodeField {
	if x != nil {
		return x.AssetKey
	}
	return nil
}

func (x *NodeInfo) GetIpmiIpaddr() *NodeField {
	if x != nil {
		return x.IpmiIpaddr
	}
	return nil
}

func (x *NodeInfo) GetIpmiNetmask() *NodeField {
	if x != nil {
		return x.IpmiNetmask
	}
	return nil
}

func (x *NodeInfo) GetIpmiPort() *NodeField {
	if x != nil {
		return x.IpmiPort
	}
	return nil
}

func (x *NodeInfo) GetIpmiGateway() *NodeField {
	if x != nil {
		return x.IpmiGateway
	}
	return nil
}

func (x *NodeInfo) GetIpmiUserName() *NodeField {
	if x != nil {
		return x.IpmiUserName
	}
	return nil
}

func (x *NodeInfo) GetIpmiInterface() *NodeField {
	if x != nil {
		return x.IpmiInterface
	}
	return nil
}

func (x *NodeInfo) GetNetDevs() map[string]*NetDev {
	if x != nil {
		return x.NetDevs
	}
	return nil
}

func (x *NodeInfo) GetTags() map[string]*NodeField {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *NodeInfo) GetKeys() map[string]*NodeField {
	if x != nil {
		return x.Keys
	}
	return nil
}

type NodeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *NodeListResponse) Reset() {
	*x = NodeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeListResponse) ProtoMessage() {}

func (x *NodeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeListResponse.ProtoReflect.Descriptor instead.
func (*NodeListResponse) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{5}
}

func (x *NodeListResponse) GetNodes() []*NodeInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiPrefix       string `protobuf:"bytes,1,opt,name=apiPrefix,proto3" json:"apiPrefix,omitempty"`
	ApiVersion      string `protobuf:"bytes,2,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	WarewulfVersion string `protobuf:"bytes,3,opt,name=warewulfVersion,proto3" json:"warewulfVersion,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{6}
}

func (x *VersionResponse) GetApiPrefix() string {
	if x != nil {
		return x.ApiPrefix
	}
	return ""
}

func (x *VersionResponse) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *VersionResponse) GetWarewulfVersion() string {
	if x != nil {
		return x.WarewulfVersion
	}
	return ""
}

var File_routes_proto protoreflect.FileDescriptor

var file_routes_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x09,
	0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf2, 0x02, 0x0a, 0x06, 0x4e, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x68, 0x77, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x68, 0x77, 0x61, 0x64, 0x64, 0x72, 0x12, 0x2b, 0x0a, 0x06,
	0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77,
	0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x06, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x12, 0x2d, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x2d, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x07,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2b, 0x0a, 0x06, 0x6f, 0x6e, 0x62, 0x6f, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x6f, 0x6e, 0x62, 0x6f, 0x6f, 0x74, 0x12, 0x2d, 0x0a,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x33, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x54, 0x61, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x22, 0x9d, 0x0b, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x37,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x6b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x12, 0x33, 0x0a, 0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x41, 0x72, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x0d,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x12, 0x3b, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x76, 0x65,
	0x72, 0x6c, 0x61, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x70, 0x78, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x69, 0x70, 0x78, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12,
	0x2f, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x33, 0x0a, 0x0a, 0x69, 0x70, 0x6d, 0x69, 0x49, 0x70, 0x61, 0x64, 0x64, 0x72, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0a, 0x69, 0x70, 0x6d, 0x69, 0x49,
	0x70, 0x61, 0x64, 0x64, 0x72, 0x12, 0x35, 0x0a, 0x0b, 0x69, 0x70, 0x6d, 0x69, 0x4e, 0x65, 0x74,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x0b, 0x69, 0x70, 0x6d, 0x69, 0x4e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x2f, 0x0a, 0x08,
	0x69, 0x70, 0x6d, 0x69, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x08, 0x69, 0x70, 0x6d, 0x69, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x35, 0x0a,
	0x0b, 0x69, 0x70, 0x6d, 0x69, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x69, 0x70, 0x6d, 0x69, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0x37, 0x0a, 0x0c, 0x69, 0x70, 0x6d, 0x69, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x0c, 0x69, 0x70, 0x6d, 0x69, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x0d, 0x69, 0x70, 0x6d, 0x69, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0d, 0x69, 0x70, 0x6d, 0x69, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x77, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4e, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x18, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x4c, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x44, 0x65, 0x76, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4c, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x4c, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x3c, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22,
	0x79, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x0f, 0x77, 0x61, 0x72, 0x65, 0x77, 0x75, 0x6c, 0x66, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x61, 0x72, 0x65, 0x77,
	0x75, 0x6c, 0x66, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xa6, 0x01, 0x0a, 0x05, 0x57,
	0x57, 0x41, 0x70, 0x69, 0x12, 0x4d, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x13, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x1a, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x4e, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x77, 0x77, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x29, 0x5a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x77,
	0x77, 0x61, 0x70, 0x69, 0x76, 0x31, 0x3b, 0x77, 0x77, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routes_proto_rawDescOnce sync.Once
	file_routes_proto_rawDescData = file_routes_proto_rawDesc
)

func file_routes_proto_rawDescGZIP() []byte {
	file_routes_proto_rawDescOnce.Do(func() {
		file_routes_proto_rawDescData = protoimpl.X.CompressGZIP(file_routes_proto_rawDescData)
	})
	return file_routes_proto_rawDescData
}

var file_routes_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_routes_proto_goTypes = []interface{}{
	(*NodeNames)(nil),        // 0: wwapi.v1.NodeNames
	(*NodeField)(nil),        // 1: wwapi.v1.NodeField
	(*NetDev)(nil),           // 2: wwapi.v1.NetDev
	(*NodeTags)(nil),         // 3: wwapi.v1.NodeTags
	(*NodeInfo)(nil),         // 4: wwapi.v1.NodeInfo
	(*NodeListResponse)(nil), // 5: wwapi.v1.NodeListResponse
	(*VersionResponse)(nil),  // 6: wwapi.v1.VersionResponse
	nil,                      // 7: wwapi.v1.NodeInfo.NetDevsEntry
	nil,                      // 8: wwapi.v1.NodeInfo.TagsEntry
	nil,                      // 9: wwapi.v1.NodeInfo.KeysEntry
	(*emptypb.Empty)(nil),    // 10: google.protobuf.Empty
}
var file_routes_proto_depIdxs = []int32{
	1,  // 0: wwapi.v1.NetDev.device:type_name -> wwapi.v1.NodeField
	1,  // 1: wwapi.v1.NetDev.hwaddr:type_name -> wwapi.v1.NodeField
	1,  // 2: wwapi.v1.NetDev.ipaddr:type_name -> wwapi.v1.NodeField
	1,  // 3: wwapi.v1.NetDev.netmask:type_name -> wwapi.v1.NodeField
	1,  // 4: wwapi.v1.NetDev.gateway:type_name -> wwapi.v1.NodeField
	1,  // 5: wwapi.v1.NetDev.type:type_name -> wwapi.v1.NodeField
	1,  // 6: wwapi.v1.NetDev.onboot:type_name -> wwapi.v1.NodeField
	1,  // 7: wwapi.v1.NetDev.default:type_name -> wwapi.v1.NodeField
	1,  // 8: wwapi.v1.NodeTags.tags:type_name -> wwapi.v1.NodeField
	1,  // 9: wwapi.v1.NodeInfo.id:type_name -> wwapi.v1.NodeField
	1,  // 10: wwapi.v1.NodeInfo.comment:type_name -> wwapi.v1.NodeField
	1,  // 11: wwapi.v1.NodeInfo.cluster:type_name -> wwapi.v1.NodeField
	1,  // 12: wwapi.v1.NodeInfo.profiles:type_name -> wwapi.v1.NodeField
	1,  // 13: wwapi.v1.NodeInfo.discoverable:type_name -> wwapi.v1.NodeField
	1,  // 14: wwapi.v1.NodeInfo.container:type_name -> wwapi.v1.NodeField
	1,  // 15: wwapi.v1.NodeInfo.kernel:type_name -> wwapi.v1.NodeField
	1,  // 16: wwapi.v1.NodeInfo.kernelArgs:type_name -> wwapi.v1.NodeField
	1,  // 17: wwapi.v1.NodeInfo.systemOverlay:type_name -> wwapi.v1.NodeField
	1,  // 18: wwapi.v1.NodeInfo.runtimeOverlay:type_name -> wwapi.v1.NodeField
	1,  // 19: wwapi.v1.NodeInfo.ipxe:type_name -> wwapi.v1.NodeField
	1,  // 20: wwapi.v1.NodeInfo.init:type_name -> wwapi.v1.NodeField
	1,  // 21: wwapi.v1.NodeInfo.root:type_name -> wwapi.v1.NodeField
	1,  // 22: wwapi.v1.NodeInfo.assetKey:type_name -> wwapi.v1.NodeField
	1,  // 23: wwapi.v1.NodeInfo.ipmiIpaddr:type_name -> wwapi.v1.NodeField
	1,  // 24: wwapi.v1.NodeInfo.ipmiNetmask:type_name -> wwapi.v1.NodeField
	1,  // 25: wwapi.v1.NodeInfo.ipmiPort:type_name -> wwapi.v1.NodeField
	1,  // 26: wwapi.v1.NodeInfo.ipmiGateway:type_name -> wwapi.v1.NodeField
	1,  // 27: wwapi.v1.NodeInfo.ipmiUserName:type_name -> wwapi.v1.NodeField
	1,  // 28: wwapi.v1.NodeInfo.ipmiInterface:type_name -> wwapi.v1.NodeField
	7,  // 29: wwapi.v1.NodeInfo.NetDevs:type_name -> wwapi.v1.NodeInfo.NetDevsEntry
	8,  // 30: wwapi.v1.NodeInfo.Tags:type_name -> wwapi.v1.NodeInfo.TagsEntry
	9,  // 31: wwapi.v1.NodeInfo.Keys:type_name -> wwapi.v1.NodeInfo.KeysEntry
	4,  // 32: wwapi.v1.NodeListResponse.nodes:type_name -> wwapi.v1.NodeInfo
	2,  // 33: wwapi.v1.NodeInfo.NetDevsEntry.value:type_name -> wwapi.v1.NetDev
	1,  // 34: wwapi.v1.NodeInfo.TagsEntry.value:type_name -> wwapi.v1.NodeField
	1,  // 35: wwapi.v1.NodeInfo.KeysEntry.value:type_name -> wwapi.v1.NodeField
	0,  // 36: wwapi.v1.WWApi.NodeList:input_type -> wwapi.v1.NodeNames
	10, // 37: wwapi.v1.WWApi.Version:input_type -> google.protobuf.Empty
	5,  // 38: wwapi.v1.WWApi.NodeList:output_type -> wwapi.v1.NodeListResponse
	6,  // 39: wwapi.v1.WWApi.Version:output_type -> wwapi.v1.VersionResponse
	38, // [38:40] is the sub-list for method output_type
	36, // [36:38] is the sub-list for method input_type
	36, // [36:36] is the sub-list for extension type_name
	36, // [36:36] is the sub-list for extension extendee
	0,  // [0:36] is the sub-list for field type_name
}

func init() { file_routes_proto_init() }
func file_routes_proto_init() {
	if File_routes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeNames); i {
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
		file_routes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeField); i {
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
		file_routes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetDev); i {
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
		file_routes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeTags); i {
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
		file_routes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_routes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeListResponse); i {
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
		file_routes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
			RawDescriptor: file_routes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routes_proto_goTypes,
		DependencyIndexes: file_routes_proto_depIdxs,
		MessageInfos:      file_routes_proto_msgTypes,
	}.Build()
	File_routes_proto = out.File
	file_routes_proto_rawDesc = nil
	file_routes_proto_goTypes = nil
	file_routes_proto_depIdxs = nil
}