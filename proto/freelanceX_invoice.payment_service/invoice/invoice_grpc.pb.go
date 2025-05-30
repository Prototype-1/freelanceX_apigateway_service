// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: invoice.proto

package invoicepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InvoiceService_CreateInvoice_FullMethodName        = "/freelanceX.invoice.InvoiceService/CreateInvoice"
	InvoiceService_GetInvoice_FullMethodName           = "/freelanceX.invoice.InvoiceService/GetInvoice"
	InvoiceService_GetInvoicesByUser_FullMethodName    = "/freelanceX.invoice.InvoiceService/GetInvoicesByUser"
	InvoiceService_GetInvoicesByProject_FullMethodName = "/freelanceX.invoice.InvoiceService/GetInvoicesByProject"
	InvoiceService_UpdateInvoiceStatus_FullMethodName  = "/freelanceX.invoice.InvoiceService/UpdateInvoiceStatus"
	InvoiceService_GetInvoicePDF_FullMethodName        = "/freelanceX.invoice.InvoiceService/GetInvoicePDF"
)

// InvoiceServiceClient is the client API for InvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// === Service ===
type InvoiceServiceClient interface {
	CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*InvoiceResponse, error)
	GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*InvoiceResponse, error)
	GetInvoicesByUser(ctx context.Context, in *GetInvoicesByUserRequest, opts ...grpc.CallOption) (*InvoicesResponse, error)
	GetInvoicesByProject(ctx context.Context, in *GetInvoicesByProjectRequest, opts ...grpc.CallOption) (*InvoicesResponse, error)
	UpdateInvoiceStatus(ctx context.Context, in *UpdateInvoiceStatusRequest, opts ...grpc.CallOption) (*InvoiceResponse, error)
	GetInvoicePDF(ctx context.Context, in *GetInvoicePDFRequest, opts ...grpc.CallOption) (*GetInvoicePDFResponse, error)
}

type invoiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceServiceClient(cc grpc.ClientConnInterface) InvoiceServiceClient {
	return &invoiceServiceClient{cc}
}

func (c *invoiceServiceClient) CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*InvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_CreateInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*InvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_GetInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoicesByUser(ctx context.Context, in *GetInvoicesByUserRequest, opts ...grpc.CallOption) (*InvoicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvoicesResponse)
	err := c.cc.Invoke(ctx, InvoiceService_GetInvoicesByUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoicesByProject(ctx context.Context, in *GetInvoicesByProjectRequest, opts ...grpc.CallOption) (*InvoicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvoicesResponse)
	err := c.cc.Invoke(ctx, InvoiceService_GetInvoicesByProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) UpdateInvoiceStatus(ctx context.Context, in *UpdateInvoiceStatusRequest, opts ...grpc.CallOption) (*InvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_UpdateInvoiceStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoicePDF(ctx context.Context, in *GetInvoicePDFRequest, opts ...grpc.CallOption) (*GetInvoicePDFResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvoicePDFResponse)
	err := c.cc.Invoke(ctx, InvoiceService_GetInvoicePDF_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServiceServer is the server API for InvoiceService service.
// All implementations must embed UnimplementedInvoiceServiceServer
// for forward compatibility.
//
// === Service ===
type InvoiceServiceServer interface {
	CreateInvoice(context.Context, *CreateInvoiceRequest) (*InvoiceResponse, error)
	GetInvoice(context.Context, *GetInvoiceRequest) (*InvoiceResponse, error)
	GetInvoicesByUser(context.Context, *GetInvoicesByUserRequest) (*InvoicesResponse, error)
	GetInvoicesByProject(context.Context, *GetInvoicesByProjectRequest) (*InvoicesResponse, error)
	UpdateInvoiceStatus(context.Context, *UpdateInvoiceStatusRequest) (*InvoiceResponse, error)
	GetInvoicePDF(context.Context, *GetInvoicePDFRequest) (*GetInvoicePDFResponse, error)
	mustEmbedUnimplementedInvoiceServiceServer()
}

// UnimplementedInvoiceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInvoiceServiceServer struct{}

func (UnimplementedInvoiceServiceServer) CreateInvoice(context.Context, *CreateInvoiceRequest) (*InvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoice(context.Context, *GetInvoiceRequest) (*InvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoicesByUser(context.Context, *GetInvoicesByUserRequest) (*InvoicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoicesByUser not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoicesByProject(context.Context, *GetInvoicesByProjectRequest) (*InvoicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoicesByProject not implemented")
}
func (UnimplementedInvoiceServiceServer) UpdateInvoiceStatus(context.Context, *UpdateInvoiceStatusRequest) (*InvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInvoiceStatus not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoicePDF(context.Context, *GetInvoicePDFRequest) (*GetInvoicePDFResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoicePDF not implemented")
}
func (UnimplementedInvoiceServiceServer) mustEmbedUnimplementedInvoiceServiceServer() {}
func (UnimplementedInvoiceServiceServer) testEmbeddedByValue()                        {}

// UnsafeInvoiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServiceServer will
// result in compilation errors.
type UnsafeInvoiceServiceServer interface {
	mustEmbedUnimplementedInvoiceServiceServer()
}

func RegisterInvoiceServiceServer(s grpc.ServiceRegistrar, srv InvoiceServiceServer) {
	// If the following call pancis, it indicates UnimplementedInvoiceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InvoiceService_ServiceDesc, srv)
}

func _InvoiceService_CreateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_CreateInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoice(ctx, req.(*CreateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_GetInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoice(ctx, req.(*GetInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoicesByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoicesByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoicesByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_GetInvoicesByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoicesByUser(ctx, req.(*GetInvoicesByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoicesByProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoicesByProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoicesByProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_GetInvoicesByProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoicesByProject(ctx, req.(*GetInvoicesByProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_UpdateInvoiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInvoiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).UpdateInvoiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_UpdateInvoiceStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).UpdateInvoiceStatus(ctx, req.(*UpdateInvoiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoicePDF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoicePDFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoicePDF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_GetInvoicePDF_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoicePDF(ctx, req.(*GetInvoicePDFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InvoiceService_ServiceDesc is the grpc.ServiceDesc for InvoiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvoiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "freelanceX.invoice.InvoiceService",
	HandlerType: (*InvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvoice",
			Handler:    _InvoiceService_CreateInvoice_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _InvoiceService_GetInvoice_Handler,
		},
		{
			MethodName: "GetInvoicesByUser",
			Handler:    _InvoiceService_GetInvoicesByUser_Handler,
		},
		{
			MethodName: "GetInvoicesByProject",
			Handler:    _InvoiceService_GetInvoicesByProject_Handler,
		},
		{
			MethodName: "UpdateInvoiceStatus",
			Handler:    _InvoiceService_UpdateInvoiceStatus_Handler,
		},
		{
			MethodName: "GetInvoicePDF",
			Handler:    _InvoiceService_GetInvoicePDF_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice.proto",
}
