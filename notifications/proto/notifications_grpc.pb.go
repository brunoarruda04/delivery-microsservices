// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: notifications.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	// Envia uma notificação para o usuário
	Send(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*Notification, error)
	// Lista todas as notificações do usuário
	List(ctx context.Context, in *User, opts ...grpc.CallOption) (NotificationService_ListClient, error)
	// Marca uma ou mais notificações como lidas
	MarkAsRead(ctx context.Context, opts ...grpc.CallOption) (NotificationService_MarkAsReadClient, error)
	// Marca todas as notificações como lidas
	MarkAllAsRead(ctx context.Context, opts ...grpc.CallOption) (NotificationService_MarkAllAsReadClient, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Send(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*Notification, error) {
	out := new(Notification)
	err := c.cc.Invoke(ctx, "/notification_pb.NotificationService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) List(ctx context.Context, in *User, opts ...grpc.CallOption) (NotificationService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[0], "/notification_pb.NotificationService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_ListClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type notificationServiceListClient struct {
	grpc.ClientStream
}

func (x *notificationServiceListClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) MarkAsRead(ctx context.Context, opts ...grpc.CallOption) (NotificationService_MarkAsReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[1], "/notification_pb.NotificationService/MarkAsRead", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceMarkAsReadClient{stream}
	return x, nil
}

type NotificationService_MarkAsReadClient interface {
	Send(*Notification) error
	Recv() (*Notification, error)
	grpc.ClientStream
}

type notificationServiceMarkAsReadClient struct {
	grpc.ClientStream
}

func (x *notificationServiceMarkAsReadClient) Send(m *Notification) error {
	return x.ClientStream.SendMsg(m)
}

func (x *notificationServiceMarkAsReadClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) MarkAllAsRead(ctx context.Context, opts ...grpc.CallOption) (NotificationService_MarkAllAsReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[2], "/notification_pb.NotificationService/MarkAllAsRead", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceMarkAllAsReadClient{stream}
	return x, nil
}

type NotificationService_MarkAllAsReadClient interface {
	Send(*Notification) error
	Recv() (*Notification, error)
	grpc.ClientStream
}

type notificationServiceMarkAllAsReadClient struct {
	grpc.ClientStream
}

func (x *notificationServiceMarkAllAsReadClient) Send(m *Notification) error {
	return x.ClientStream.SendMsg(m)
}

func (x *notificationServiceMarkAllAsReadClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	// Envia uma notificação para o usuário
	Send(context.Context, *Notification) (*Notification, error)
	// Lista todas as notificações do usuário
	List(*User, NotificationService_ListServer) error
	// Marca uma ou mais notificações como lidas
	MarkAsRead(NotificationService_MarkAsReadServer) error
	// Marca todas as notificações como lidas
	MarkAllAsRead(NotificationService_MarkAllAsReadServer) error
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) Send(context.Context, *Notification) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedNotificationServiceServer) List(*User, NotificationService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNotificationServiceServer) MarkAsRead(NotificationService_MarkAsReadServer) error {
	return status.Errorf(codes.Unimplemented, "method MarkAsRead not implemented")
}
func (UnimplementedNotificationServiceServer) MarkAllAsRead(NotificationService_MarkAllAsReadServer) error {
	return status.Errorf(codes.Unimplemented, "method MarkAllAsRead not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification_pb.NotificationService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Send(ctx, req.(*Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).List(m, &notificationServiceListServer{stream})
}

type NotificationService_ListServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type notificationServiceListServer struct {
	grpc.ServerStream
}

func (x *notificationServiceListServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationService_MarkAsRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NotificationServiceServer).MarkAsRead(&notificationServiceMarkAsReadServer{stream})
}

type NotificationService_MarkAsReadServer interface {
	Send(*Notification) error
	Recv() (*Notification, error)
	grpc.ServerStream
}

type notificationServiceMarkAsReadServer struct {
	grpc.ServerStream
}

func (x *notificationServiceMarkAsReadServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func (x *notificationServiceMarkAsReadServer) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NotificationService_MarkAllAsRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NotificationServiceServer).MarkAllAsRead(&notificationServiceMarkAllAsReadServer{stream})
}

type NotificationService_MarkAllAsReadServer interface {
	Send(*Notification) error
	Recv() (*Notification, error)
	grpc.ServerStream
}

type notificationServiceMarkAllAsReadServer struct {
	grpc.ServerStream
}

func (x *notificationServiceMarkAllAsReadServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func (x *notificationServiceMarkAllAsReadServer) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification_pb.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _NotificationService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _NotificationService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MarkAsRead",
			Handler:       _NotificationService_MarkAsRead_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MarkAllAsRead",
			Handler:       _NotificationService_MarkAllAsRead_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "notifications.proto",
}