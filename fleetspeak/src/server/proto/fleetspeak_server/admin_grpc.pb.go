// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: fleetspeak/src/server/proto/fleetspeak_server/admin.proto

package fleetspeak_server

import (
	context "context"
	fleetspeak "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// CreateBroadcast creates a FS broadcast, potentially sending a message to
	// many machines in the fleet.
	CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// ListActiveBroadcasts lists the currently active FS broadcasts.
	ListActiveBroadcasts(ctx context.Context, in *ListActiveBroadcastsRequest, opts ...grpc.CallOption) (*ListActiveBroadcastsResponse, error)
	// ListClients lists the currently active FS clients.
	ListClients(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*ListClientsResponse, error)
	// StreamClientIds lists the currently active FS clients as a stream.
	StreamClientIds(ctx context.Context, in *StreamClientIdsRequest, opts ...grpc.CallOption) (Admin_StreamClientIdsClient, error)
	// ListClientContacts lists the contact history for a client.
	ListClientContacts(ctx context.Context, in *ListClientContactsRequest, opts ...grpc.CallOption) (*ListClientContactsResponse, error)
	// StreamClientContacts lists the contact history for a client as a stream.
	StreamClientContacts(ctx context.Context, in *StreamClientContactsRequest, opts ...grpc.CallOption) (Admin_StreamClientContactsClient, error)
	// GetMessageStatus retrieves the current status of a message.
	GetMessageStatus(ctx context.Context, in *GetMessageStatusRequest, opts ...grpc.CallOption) (*GetMessageStatusResponse, error)
	// InsertMessage inserts a message into the Fleetspeak system to be processed
	// by the server or delivered to a client.
	// TODO: Have this method return the message that is written to the
	// datastore (or at least the id).
	InsertMessage(ctx context.Context, in *fleetspeak.Message, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// DeletePendingMessages clears message queues for given clients.
	DeletePendingMessages(ctx context.Context, in *DeletePendingMessagesRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// Returns the pending messages for given clients.
	GetPendingMessages(ctx context.Context, in *GetPendingMessagesRequest, opts ...grpc.CallOption) (*GetPendingMessagesResponse, error)
	// Returns the number of pending messages for the given clients.
	GetPendingMessageCount(ctx context.Context, in *GetPendingMessageCountRequest, opts ...grpc.CallOption) (*GetPendingMessageCountResponse, error)
	// StoreFile inserts a file into the Fleetspeak system.
	StoreFile(ctx context.Context, in *StoreFileRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// KeepAlive does as little as possible.
	KeepAlive(ctx context.Context, in *fleetspeak.EmptyMessage, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// BlacklistClient marks a client_id as invalid, forcing all Fleetspeak
	// clients using it to rekey.
	BlacklistClient(ctx context.Context, in *BlacklistClientRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error)
	// FetchClientResourceUsageRecords returns all resource usage records for a
	// single client with a given limit on the number of records.
	FetchClientResourceUsageRecords(ctx context.Context, in *FetchClientResourceUsageRecordsRequest, opts ...grpc.CallOption) (*FetchClientResourceUsageRecordsResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/CreateBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListActiveBroadcasts(ctx context.Context, in *ListActiveBroadcastsRequest, opts ...grpc.CallOption) (*ListActiveBroadcastsResponse, error) {
	out := new(ListActiveBroadcastsResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/ListActiveBroadcasts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListClients(ctx context.Context, in *ListClientsRequest, opts ...grpc.CallOption) (*ListClientsResponse, error) {
	out := new(ListClientsResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/ListClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StreamClientIds(ctx context.Context, in *StreamClientIdsRequest, opts ...grpc.CallOption) (Admin_StreamClientIdsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Admin_ServiceDesc.Streams[0], "/fleetspeak.server.Admin/StreamClientIds", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminStreamClientIdsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Admin_StreamClientIdsClient interface {
	Recv() (*StreamClientIdsResponse, error)
	grpc.ClientStream
}

type adminStreamClientIdsClient struct {
	grpc.ClientStream
}

func (x *adminStreamClientIdsClient) Recv() (*StreamClientIdsResponse, error) {
	m := new(StreamClientIdsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminClient) ListClientContacts(ctx context.Context, in *ListClientContactsRequest, opts ...grpc.CallOption) (*ListClientContactsResponse, error) {
	out := new(ListClientContactsResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/ListClientContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StreamClientContacts(ctx context.Context, in *StreamClientContactsRequest, opts ...grpc.CallOption) (Admin_StreamClientContactsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Admin_ServiceDesc.Streams[1], "/fleetspeak.server.Admin/StreamClientContacts", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminStreamClientContactsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Admin_StreamClientContactsClient interface {
	Recv() (*StreamClientContactsResponse, error)
	grpc.ClientStream
}

type adminStreamClientContactsClient struct {
	grpc.ClientStream
}

func (x *adminStreamClientContactsClient) Recv() (*StreamClientContactsResponse, error) {
	m := new(StreamClientContactsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminClient) GetMessageStatus(ctx context.Context, in *GetMessageStatusRequest, opts ...grpc.CallOption) (*GetMessageStatusResponse, error) {
	out := new(GetMessageStatusResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/GetMessageStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) InsertMessage(ctx context.Context, in *fleetspeak.Message, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/InsertMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeletePendingMessages(ctx context.Context, in *DeletePendingMessagesRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/DeletePendingMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetPendingMessages(ctx context.Context, in *GetPendingMessagesRequest, opts ...grpc.CallOption) (*GetPendingMessagesResponse, error) {
	out := new(GetPendingMessagesResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/GetPendingMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetPendingMessageCount(ctx context.Context, in *GetPendingMessageCountRequest, opts ...grpc.CallOption) (*GetPendingMessageCountResponse, error) {
	out := new(GetPendingMessageCountResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/GetPendingMessageCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StoreFile(ctx context.Context, in *StoreFileRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/StoreFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) KeepAlive(ctx context.Context, in *fleetspeak.EmptyMessage, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) BlacklistClient(ctx context.Context, in *BlacklistClientRequest, opts ...grpc.CallOption) (*fleetspeak.EmptyMessage, error) {
	out := new(fleetspeak.EmptyMessage)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/BlacklistClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) FetchClientResourceUsageRecords(ctx context.Context, in *FetchClientResourceUsageRecordsRequest, opts ...grpc.CallOption) (*FetchClientResourceUsageRecordsResponse, error) {
	out := new(FetchClientResourceUsageRecordsResponse)
	err := c.cc.Invoke(ctx, "/fleetspeak.server.Admin/FetchClientResourceUsageRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// CreateBroadcast creates a FS broadcast, potentially sending a message to
	// many machines in the fleet.
	CreateBroadcast(context.Context, *CreateBroadcastRequest) (*fleetspeak.EmptyMessage, error)
	// ListActiveBroadcasts lists the currently active FS broadcasts.
	ListActiveBroadcasts(context.Context, *ListActiveBroadcastsRequest) (*ListActiveBroadcastsResponse, error)
	// ListClients lists the currently active FS clients.
	ListClients(context.Context, *ListClientsRequest) (*ListClientsResponse, error)
	// StreamClientIds lists the currently active FS clients as a stream.
	StreamClientIds(*StreamClientIdsRequest, Admin_StreamClientIdsServer) error
	// ListClientContacts lists the contact history for a client.
	ListClientContacts(context.Context, *ListClientContactsRequest) (*ListClientContactsResponse, error)
	// StreamClientContacts lists the contact history for a client as a stream.
	StreamClientContacts(*StreamClientContactsRequest, Admin_StreamClientContactsServer) error
	// GetMessageStatus retrieves the current status of a message.
	GetMessageStatus(context.Context, *GetMessageStatusRequest) (*GetMessageStatusResponse, error)
	// InsertMessage inserts a message into the Fleetspeak system to be processed
	// by the server or delivered to a client.
	// TODO: Have this method return the message that is written to the
	// datastore (or at least the id).
	InsertMessage(context.Context, *fleetspeak.Message) (*fleetspeak.EmptyMessage, error)
	// DeletePendingMessages clears message queues for given clients.
	DeletePendingMessages(context.Context, *DeletePendingMessagesRequest) (*fleetspeak.EmptyMessage, error)
	// Returns the pending messages for given clients.
	GetPendingMessages(context.Context, *GetPendingMessagesRequest) (*GetPendingMessagesResponse, error)
	// Returns the number of pending messages for the given clients.
	GetPendingMessageCount(context.Context, *GetPendingMessageCountRequest) (*GetPendingMessageCountResponse, error)
	// StoreFile inserts a file into the Fleetspeak system.
	StoreFile(context.Context, *StoreFileRequest) (*fleetspeak.EmptyMessage, error)
	// KeepAlive does as little as possible.
	KeepAlive(context.Context, *fleetspeak.EmptyMessage) (*fleetspeak.EmptyMessage, error)
	// BlacklistClient marks a client_id as invalid, forcing all Fleetspeak
	// clients using it to rekey.
	BlacklistClient(context.Context, *BlacklistClientRequest) (*fleetspeak.EmptyMessage, error)
	// FetchClientResourceUsageRecords returns all resource usage records for a
	// single client with a given limit on the number of records.
	FetchClientResourceUsageRecords(context.Context, *FetchClientResourceUsageRecordsRequest) (*FetchClientResourceUsageRecordsResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) CreateBroadcast(context.Context, *CreateBroadcastRequest) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBroadcast not implemented")
}
func (UnimplementedAdminServer) ListActiveBroadcasts(context.Context, *ListActiveBroadcastsRequest) (*ListActiveBroadcastsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActiveBroadcasts not implemented")
}
func (UnimplementedAdminServer) ListClients(context.Context, *ListClientsRequest) (*ListClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClients not implemented")
}
func (UnimplementedAdminServer) StreamClientIds(*StreamClientIdsRequest, Admin_StreamClientIdsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClientIds not implemented")
}
func (UnimplementedAdminServer) ListClientContacts(context.Context, *ListClientContactsRequest) (*ListClientContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClientContacts not implemented")
}
func (UnimplementedAdminServer) StreamClientContacts(*StreamClientContactsRequest, Admin_StreamClientContactsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClientContacts not implemented")
}
func (UnimplementedAdminServer) GetMessageStatus(context.Context, *GetMessageStatusRequest) (*GetMessageStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageStatus not implemented")
}
func (UnimplementedAdminServer) InsertMessage(context.Context, *fleetspeak.Message) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertMessage not implemented")
}
func (UnimplementedAdminServer) DeletePendingMessages(context.Context, *DeletePendingMessagesRequest) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePendingMessages not implemented")
}
func (UnimplementedAdminServer) GetPendingMessages(context.Context, *GetPendingMessagesRequest) (*GetPendingMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingMessages not implemented")
}
func (UnimplementedAdminServer) GetPendingMessageCount(context.Context, *GetPendingMessageCountRequest) (*GetPendingMessageCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingMessageCount not implemented")
}
func (UnimplementedAdminServer) StoreFile(context.Context, *StoreFileRequest) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFile not implemented")
}
func (UnimplementedAdminServer) KeepAlive(context.Context, *fleetspeak.EmptyMessage) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (UnimplementedAdminServer) BlacklistClient(context.Context, *BlacklistClientRequest) (*fleetspeak.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlacklistClient not implemented")
}
func (UnimplementedAdminServer) FetchClientResourceUsageRecords(context.Context, *FetchClientResourceUsageRecordsRequest) (*FetchClientResourceUsageRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClientResourceUsageRecords not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_CreateBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/CreateBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateBroadcast(ctx, req.(*CreateBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListActiveBroadcasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActiveBroadcastsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListActiveBroadcasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/ListActiveBroadcasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListActiveBroadcasts(ctx, req.(*ListActiveBroadcastsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/ListClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListClients(ctx, req.(*ListClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StreamClientIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamClientIdsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServer).StreamClientIds(m, &adminStreamClientIdsServer{stream})
}

type Admin_StreamClientIdsServer interface {
	Send(*StreamClientIdsResponse) error
	grpc.ServerStream
}

type adminStreamClientIdsServer struct {
	grpc.ServerStream
}

func (x *adminStreamClientIdsServer) Send(m *StreamClientIdsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Admin_ListClientContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListClientContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/ListClientContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListClientContacts(ctx, req.(*ListClientContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StreamClientContacts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamClientContactsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServer).StreamClientContacts(m, &adminStreamClientContactsServer{stream})
}

type Admin_StreamClientContactsServer interface {
	Send(*StreamClientContactsResponse) error
	grpc.ServerStream
}

type adminStreamClientContactsServer struct {
	grpc.ServerStream
}

func (x *adminStreamClientContactsServer) Send(m *StreamClientContactsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Admin_GetMessageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetMessageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/GetMessageStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetMessageStatus(ctx, req.(*GetMessageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_InsertMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fleetspeak.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).InsertMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/InsertMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).InsertMessage(ctx, req.(*fleetspeak.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeletePendingMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePendingMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeletePendingMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/DeletePendingMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeletePendingMessages(ctx, req.(*DeletePendingMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetPendingMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetPendingMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/GetPendingMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetPendingMessages(ctx, req.(*GetPendingMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetPendingMessageCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingMessageCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetPendingMessageCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/GetPendingMessageCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetPendingMessageCount(ctx, req.(*GetPendingMessageCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StoreFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).StoreFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/StoreFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).StoreFile(ctx, req.(*StoreFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fleetspeak.EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).KeepAlive(ctx, req.(*fleetspeak.EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_BlacklistClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlacklistClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).BlacklistClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/BlacklistClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).BlacklistClient(ctx, req.(*BlacklistClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_FetchClientResourceUsageRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchClientResourceUsageRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).FetchClientResourceUsageRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fleetspeak.server.Admin/FetchClientResourceUsageRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).FetchClientResourceUsageRecords(ctx, req.(*FetchClientResourceUsageRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fleetspeak.server.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBroadcast",
			Handler:    _Admin_CreateBroadcast_Handler,
		},
		{
			MethodName: "ListActiveBroadcasts",
			Handler:    _Admin_ListActiveBroadcasts_Handler,
		},
		{
			MethodName: "ListClients",
			Handler:    _Admin_ListClients_Handler,
		},
		{
			MethodName: "ListClientContacts",
			Handler:    _Admin_ListClientContacts_Handler,
		},
		{
			MethodName: "GetMessageStatus",
			Handler:    _Admin_GetMessageStatus_Handler,
		},
		{
			MethodName: "InsertMessage",
			Handler:    _Admin_InsertMessage_Handler,
		},
		{
			MethodName: "DeletePendingMessages",
			Handler:    _Admin_DeletePendingMessages_Handler,
		},
		{
			MethodName: "GetPendingMessages",
			Handler:    _Admin_GetPendingMessages_Handler,
		},
		{
			MethodName: "GetPendingMessageCount",
			Handler:    _Admin_GetPendingMessageCount_Handler,
		},
		{
			MethodName: "StoreFile",
			Handler:    _Admin_StoreFile_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _Admin_KeepAlive_Handler,
		},
		{
			MethodName: "BlacklistClient",
			Handler:    _Admin_BlacklistClient_Handler,
		},
		{
			MethodName: "FetchClientResourceUsageRecords",
			Handler:    _Admin_FetchClientResourceUsageRecords_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClientIds",
			Handler:       _Admin_StreamClientIds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamClientContacts",
			Handler:       _Admin_StreamClientContacts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fleetspeak/src/server/proto/fleetspeak_server/admin.proto",
}
