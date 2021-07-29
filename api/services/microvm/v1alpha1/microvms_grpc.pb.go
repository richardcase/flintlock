// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

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

// MicroVMClient is the client API for MicroVM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroVMClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	ListStream(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (MicroVM_ListStreamClient, error)
}

type microVMClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroVMClient(cc grpc.ClientConnInterface) MicroVMClient {
	return &microVMClient{cc}
}

func (c *microVMClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/microvm.services.api.v1alpha1.MicroVM/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/microvm.services.api.v1alpha1.MicroVM/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/microvm.services.api.v1alpha1.MicroVM/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/microvm.services.api.v1alpha1.MicroVM/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/microvm.services.api.v1alpha1.MicroVM/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) ListStream(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (MicroVM_ListStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MicroVM_ServiceDesc.Streams[0], "/microvm.services.api.v1alpha1.MicroVM/ListStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &microVMListStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MicroVM_ListStreamClient interface {
	Recv() (*ListMessage, error)
	grpc.ClientStream
}

type microVMListStreamClient struct {
	grpc.ClientStream
}

func (x *microVMListStreamClient) Recv() (*ListMessage, error) {
	m := new(ListMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MicroVMServer is the server API for MicroVM service.
// All implementations should embed UnimplementedMicroVMServer
// for forward compatibility
type MicroVMServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	ListStream(*ListRequest, MicroVM_ListStreamServer) error
}

// UnimplementedMicroVMServer should be embedded to have forward compatible implementations.
type UnimplementedMicroVMServer struct {
}

func (UnimplementedMicroVMServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMicroVMServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMicroVMServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMicroVMServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMicroVMServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMicroVMServer) ListStream(*ListRequest, MicroVM_ListStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStream not implemented")
}

// UnsafeMicroVMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroVMServer will
// result in compilation errors.
type UnsafeMicroVMServer interface {
	mustEmbedUnimplementedMicroVMServer()
}

func RegisterMicroVMServer(s grpc.ServiceRegistrar, srv MicroVMServer) {
	s.RegisterService(&MicroVM_ServiceDesc, srv)
}

func _MicroVM_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microvm.services.api.v1alpha1.MicroVM/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microvm.services.api.v1alpha1.MicroVM/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microvm.services.api.v1alpha1.MicroVM/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microvm.services.api.v1alpha1.MicroVM/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microvm.services.api.v1alpha1.MicroVM/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_ListStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MicroVMServer).ListStream(m, &microVMListStreamServer{stream})
}

type MicroVM_ListStreamServer interface {
	Send(*ListMessage) error
	grpc.ServerStream
}

type microVMListStreamServer struct {
	grpc.ServerStream
}

func (x *microVMListStreamServer) Send(m *ListMessage) error {
	return x.ServerStream.SendMsg(m)
}

// MicroVM_ServiceDesc is the grpc.ServiceDesc for MicroVM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroVM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microvm.services.api.v1alpha1.MicroVM",
	HandlerType: (*MicroVMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MicroVM_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MicroVM_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MicroVM_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MicroVM_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MicroVM_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStream",
			Handler:       _MicroVM_ListStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/microvm/v1alpha1/microvms.proto",
}
