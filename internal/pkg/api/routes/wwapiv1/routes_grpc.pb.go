// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: routes.proto

package wwapiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WWApiClient is the client API for WWApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WWApiClient interface {
	// ContainerBuild builds zero or more containers.
	ContainerBuild(ctx context.Context, in *ContainerBuildParameter, opts ...grpc.CallOption) (*ContainerListResponse, error)
	// ContainerDelete removes one or more container from Warewulf management.
	ContainerDelete(ctx context.Context, in *ContainerDeleteParameter, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ContainerImport imports a container to Warewulf.
	ContainerImport(ctx context.Context, in *ContainerImportParameter, opts ...grpc.CallOption) (*ContainerListResponse, error)
	// ContainerList lists ContainerInfo for each container.
	ContainerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ContainerListResponse, error)
	// ContainerShow lists ContainerShow for each container.
	ContainerShow(ctx context.Context, in *ContainerShowParameter, opts ...grpc.CallOption) (*ContainerShowResponse, error)
	// NodeAdd adds one or more nodes for management by Warewulf and returns
	// the added nodes. Node fields may be shimmed in per profiles.
	NodeAdd(ctx context.Context, in *NodeAddParameter, opts ...grpc.CallOption) (*NodeListResponse, error)
	// NodeDelete removes one or more nodes from Warewulf management.
	NodeDelete(ctx context.Context, in *NodeDeleteParameter, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// NodeList lists some or all nodes managed by Warewulf.
	NodeList(ctx context.Context, in *NodeNames, opts ...grpc.CallOption) (*NodeListResponse, error)
	// NodeSet updates node fields for one or more nodes.
	NodeSet(ctx context.Context, in *NodeSetParameter, opts ...grpc.CallOption) (*NodeListResponse, error)
	// NodeStatus returns the imaging state for nodes.
	// This requires warewulfd.
	NodeStatus(ctx context.Context, in *NodeNames, opts ...grpc.CallOption) (*NodeStatusResponse, error)
	// Version returns the wwapi version, the api prefix, and the Warewulf
	// version. This is also useful for testing if the service is up.
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type wWApiClient struct {
	cc grpc.ClientConnInterface
}

func NewWWApiClient(cc grpc.ClientConnInterface) WWApiClient {
	return &wWApiClient{cc}
}

func (c *wWApiClient) ContainerBuild(ctx context.Context, in *ContainerBuildParameter, opts ...grpc.CallOption) (*ContainerListResponse, error) {
	out := new(ContainerListResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/ContainerBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) ContainerDelete(ctx context.Context, in *ContainerDeleteParameter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/ContainerDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) ContainerImport(ctx context.Context, in *ContainerImportParameter, opts ...grpc.CallOption) (*ContainerListResponse, error) {
	out := new(ContainerListResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/ContainerImport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) ContainerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ContainerListResponse, error) {
	out := new(ContainerListResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/ContainerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) ContainerShow(ctx context.Context, in *ContainerShowParameter, opts ...grpc.CallOption) (*ContainerShowResponse, error) {
	out := new(ContainerShowResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/ContainerShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) NodeAdd(ctx context.Context, in *NodeAddParameter, opts ...grpc.CallOption) (*NodeListResponse, error) {
	out := new(NodeListResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/NodeAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) NodeDelete(ctx context.Context, in *NodeDeleteParameter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/NodeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) NodeList(ctx context.Context, in *NodeNames, opts ...grpc.CallOption) (*NodeListResponse, error) {
	out := new(NodeListResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/NodeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) NodeSet(ctx context.Context, in *NodeSetParameter, opts ...grpc.CallOption) (*NodeListResponse, error) {
	out := new(NodeListResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/NodeSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) NodeStatus(ctx context.Context, in *NodeNames, opts ...grpc.CallOption) (*NodeStatusResponse, error) {
	out := new(NodeStatusResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/NodeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wWApiClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/wwapi.v1.WWApi/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WWApiServer is the server API for WWApi service.
// All implementations must embed UnimplementedWWApiServer
// for forward compatibility
type WWApiServer interface {
	// ContainerBuild builds zero or more containers.
	ContainerBuild(context.Context, *ContainerBuildParameter) (*ContainerListResponse, error)
	// ContainerDelete removes one or more container from Warewulf management.
	ContainerDelete(context.Context, *ContainerDeleteParameter) (*emptypb.Empty, error)
	// ContainerImport imports a container to Warewulf.
	ContainerImport(context.Context, *ContainerImportParameter) (*ContainerListResponse, error)
	// ContainerList lists ContainerInfo for each container.
	ContainerList(context.Context, *emptypb.Empty) (*ContainerListResponse, error)
	// ContainerShow lists ContainerShow for each container.
	ContainerShow(context.Context, *ContainerShowParameter) (*ContainerShowResponse, error)
	// NodeAdd adds one or more nodes for management by Warewulf and returns
	// the added nodes. Node fields may be shimmed in per profiles.
	NodeAdd(context.Context, *NodeAddParameter) (*NodeListResponse, error)
	// NodeDelete removes one or more nodes from Warewulf management.
	NodeDelete(context.Context, *NodeDeleteParameter) (*emptypb.Empty, error)
	// NodeList lists some or all nodes managed by Warewulf.
	NodeList(context.Context, *NodeNames) (*NodeListResponse, error)
	// NodeSet updates node fields for one or more nodes.
	NodeSet(context.Context, *NodeSetParameter) (*NodeListResponse, error)
	// NodeStatus returns the imaging state for nodes.
	// This requires warewulfd.
	NodeStatus(context.Context, *NodeNames) (*NodeStatusResponse, error)
	// Version returns the wwapi version, the api prefix, and the Warewulf
	// version. This is also useful for testing if the service is up.
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	mustEmbedUnimplementedWWApiServer()
}

// UnimplementedWWApiServer must be embedded to have forward compatible implementations.
type UnimplementedWWApiServer struct {
}

func (UnimplementedWWApiServer) ContainerBuild(context.Context, *ContainerBuildParameter) (*ContainerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerBuild not implemented")
}
func (UnimplementedWWApiServer) ContainerDelete(context.Context, *ContainerDeleteParameter) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerDelete not implemented")
}
func (UnimplementedWWApiServer) ContainerImport(context.Context, *ContainerImportParameter) (*ContainerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerImport not implemented")
}
func (UnimplementedWWApiServer) ContainerList(context.Context, *emptypb.Empty) (*ContainerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerList not implemented")
}
func (UnimplementedWWApiServer) ContainerShow(context.Context, *ContainerShowParameter) (*ContainerShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerShow not implemented")
}
func (UnimplementedWWApiServer) NodeAdd(context.Context, *NodeAddParameter) (*NodeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeAdd not implemented")
}
func (UnimplementedWWApiServer) NodeDelete(context.Context, *NodeDeleteParameter) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeDelete not implemented")
}
func (UnimplementedWWApiServer) NodeList(context.Context, *NodeNames) (*NodeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeList not implemented")
}
func (UnimplementedWWApiServer) NodeSet(context.Context, *NodeSetParameter) (*NodeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeSet not implemented")
}
func (UnimplementedWWApiServer) NodeStatus(context.Context, *NodeNames) (*NodeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeStatus not implemented")
}
func (UnimplementedWWApiServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedWWApiServer) mustEmbedUnimplementedWWApiServer() {}

// UnsafeWWApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WWApiServer will
// result in compilation errors.
type UnsafeWWApiServer interface {
	mustEmbedUnimplementedWWApiServer()
}

func RegisterWWApiServer(s grpc.ServiceRegistrar, srv WWApiServer) {
	s.RegisterService(&WWApi_ServiceDesc, srv)
}

func _WWApi_ContainerBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerBuildParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).ContainerBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/ContainerBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).ContainerBuild(ctx, req.(*ContainerBuildParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_ContainerDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerDeleteParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).ContainerDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/ContainerDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).ContainerDelete(ctx, req.(*ContainerDeleteParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_ContainerImport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerImportParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).ContainerImport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/ContainerImport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).ContainerImport(ctx, req.(*ContainerImportParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_ContainerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).ContainerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/ContainerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).ContainerList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_ContainerShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerShowParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).ContainerShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/ContainerShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).ContainerShow(ctx, req.(*ContainerShowParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_NodeAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeAddParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).NodeAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/NodeAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).NodeAdd(ctx, req.(*NodeAddParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_NodeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeDeleteParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).NodeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/NodeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).NodeDelete(ctx, req.(*NodeDeleteParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_NodeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeNames)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).NodeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/NodeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).NodeList(ctx, req.(*NodeNames))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_NodeSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeSetParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).NodeSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/NodeSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).NodeSet(ctx, req.(*NodeSetParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_NodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeNames)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).NodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/NodeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).NodeStatus(ctx, req.(*NodeNames))
	}
	return interceptor(ctx, in, info, handler)
}

func _WWApi_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WWApiServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwapi.v1.WWApi/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WWApiServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// WWApi_ServiceDesc is the grpc.ServiceDesc for WWApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WWApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wwapi.v1.WWApi",
	HandlerType: (*WWApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContainerBuild",
			Handler:    _WWApi_ContainerBuild_Handler,
		},
		{
			MethodName: "ContainerDelete",
			Handler:    _WWApi_ContainerDelete_Handler,
		},
		{
			MethodName: "ContainerImport",
			Handler:    _WWApi_ContainerImport_Handler,
		},
		{
			MethodName: "ContainerList",
			Handler:    _WWApi_ContainerList_Handler,
		},
		{
			MethodName: "ContainerShow",
			Handler:    _WWApi_ContainerShow_Handler,
		},
		{
			MethodName: "NodeAdd",
			Handler:    _WWApi_NodeAdd_Handler,
		},
		{
			MethodName: "NodeDelete",
			Handler:    _WWApi_NodeDelete_Handler,
		},
		{
			MethodName: "NodeList",
			Handler:    _WWApi_NodeList_Handler,
		},
		{
			MethodName: "NodeSet",
			Handler:    _WWApi_NodeSet_Handler,
		},
		{
			MethodName: "NodeStatus",
			Handler:    _WWApi_NodeStatus_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _WWApi_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "routes.proto",
}
