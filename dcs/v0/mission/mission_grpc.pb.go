// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: dcs/mission/v0/mission.proto

package mission

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

// MissionServiceClient is the client API for MissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MissionServiceClient interface {
	// Streams DCS game generated Events.
	// See https://wiki.hoggitworld.com/view/Category:Events
	StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (MissionService_StreamEventsClient, error)
	// Streams unit updates
	// Provides similar functionality as Tacview but at a much lower update rate
	// so puts less load on the server. Suitable for things like online maps but
	// not as a Tacview replacement.
	StreamUnits(ctx context.Context, in *StreamUnitsRequest, opts ...grpc.CallOption) (MissionService_StreamUnitsClient, error)
	// Returns the mission's in-game starttime as an ISO 8601 formatted datetime
	// string.
	GetScenarioStartTime(ctx context.Context, in *GetScenarioStartTimeRequest, opts ...grpc.CallOption) (*GetScenarioStartTimeResponse, error)
	// Returns the mission's in-game current time as an ISO 8601 formatted
	// datetime string.
	GetScenarioCurrentTime(ctx context.Context, in *GetScenarioCurrentTimeRequest, opts ...grpc.CallOption) (*GetScenarioCurrentTimeResponse, error)
	// Adds a new mission command
	// See https://wiki.hoggitworld.com/view/DCS_func_addCommand
	AddMissionCommand(ctx context.Context, in *AddMissionCommandRequest, opts ...grpc.CallOption) (*AddMissionCommandResponse, error)
	// Adds a new command sub menu
	// See https://wiki.hoggitworld.com/view/DCS_func_addSubMenu
	AddMissionCommandSubMenu(ctx context.Context, in *AddMissionCommandSubMenuRequest, opts ...grpc.CallOption) (*AddMissionCommandSubMenuResponse, error)
	// Removes a registered mission command.
	// See https://wiki.hoggitworld.com/view/DCS_func_removeItem
	RemoveMissionCommandItem(ctx context.Context, in *RemoveMissionCommandItemRequest, opts ...grpc.CallOption) (*RemoveMissionCommandItemResponse, error)
	// Adds a new coalition command
	// See https://wiki.hoggitworld.com/view/DCS_func_addCommandForCoalition
	AddCoalitionCommand(ctx context.Context, in *AddCoalitionCommandRequest, opts ...grpc.CallOption) (*AddCoalitionCommandResponse, error)
	// Adds a new coalition command sub menu
	// See https://wiki.hoggitworld.com/view/DCS_func_addSubMenuForCoalition
	AddCoalitionCommandSubMenu(ctx context.Context, in *AddCoalitionCommandSubMenuRequest, opts ...grpc.CallOption) (*AddCoalitionCommandSubMenuResponse, error)
	// Removes a registered coalition command.
	// See https://wiki.hoggitworld.com/view/DCS_func_removeItemForCoalition
	RemoveCoalitionCommandItem(ctx context.Context, in *RemoveCoalitionCommandItemRequest, opts ...grpc.CallOption) (*RemoveCoalitionCommandItemResponse, error)
	// Adds a new group command
	// See https://wiki.hoggitworld.com/view/DCS_func_addCommandForGroup
	AddGroupCommand(ctx context.Context, in *AddGroupCommandRequest, opts ...grpc.CallOption) (*AddGroupCommandResponse, error)
	// Adds a new group command sub menu
	// See https://wiki.hoggitworld.com/view/DCS_func_addSubMenuForGroup
	AddGroupCommandSubMenu(ctx context.Context, in *AddGroupCommandSubMenuRequest, opts ...grpc.CallOption) (*AddGroupCommandSubMenuResponse, error)
	// Removes a group coalition command.
	// See https://wiki.hoggitworld.com/view/DCS_func_removeItemForGroup
	RemoveGroupCommandItem(ctx context.Context, in *RemoveGroupCommandItemRequest, opts ...grpc.CallOption) (*RemoveGroupCommandItemResponse, error)
	// Returns an ID for the current session.
	// The ID will change upon mission change or server restart.
	GetSessionId(ctx context.Context, in *GetSessionIdRequest, opts ...grpc.CallOption) (*GetSessionIdResponse, error)
}

type missionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMissionServiceClient(cc grpc.ClientConnInterface) MissionServiceClient {
	return &missionServiceClient{cc}
}

func (c *missionServiceClient) StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (MissionService_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MissionService_ServiceDesc.Streams[0], "/dcs.mission.v0.MissionService/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &missionServiceStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MissionService_StreamEventsClient interface {
	Recv() (*StreamEventsResponse, error)
	grpc.ClientStream
}

type missionServiceStreamEventsClient struct {
	grpc.ClientStream
}

func (x *missionServiceStreamEventsClient) Recv() (*StreamEventsResponse, error) {
	m := new(StreamEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *missionServiceClient) StreamUnits(ctx context.Context, in *StreamUnitsRequest, opts ...grpc.CallOption) (MissionService_StreamUnitsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MissionService_ServiceDesc.Streams[1], "/dcs.mission.v0.MissionService/StreamUnits", opts...)
	if err != nil {
		return nil, err
	}
	x := &missionServiceStreamUnitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MissionService_StreamUnitsClient interface {
	Recv() (*StreamUnitsResponse, error)
	grpc.ClientStream
}

type missionServiceStreamUnitsClient struct {
	grpc.ClientStream
}

func (x *missionServiceStreamUnitsClient) Recv() (*StreamUnitsResponse, error) {
	m := new(StreamUnitsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *missionServiceClient) GetScenarioStartTime(ctx context.Context, in *GetScenarioStartTimeRequest, opts ...grpc.CallOption) (*GetScenarioStartTimeResponse, error) {
	out := new(GetScenarioStartTimeResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/GetScenarioStartTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) GetScenarioCurrentTime(ctx context.Context, in *GetScenarioCurrentTimeRequest, opts ...grpc.CallOption) (*GetScenarioCurrentTimeResponse, error) {
	out := new(GetScenarioCurrentTimeResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/GetScenarioCurrentTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) AddMissionCommand(ctx context.Context, in *AddMissionCommandRequest, opts ...grpc.CallOption) (*AddMissionCommandResponse, error) {
	out := new(AddMissionCommandResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/AddMissionCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) AddMissionCommandSubMenu(ctx context.Context, in *AddMissionCommandSubMenuRequest, opts ...grpc.CallOption) (*AddMissionCommandSubMenuResponse, error) {
	out := new(AddMissionCommandSubMenuResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/AddMissionCommandSubMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) RemoveMissionCommandItem(ctx context.Context, in *RemoveMissionCommandItemRequest, opts ...grpc.CallOption) (*RemoveMissionCommandItemResponse, error) {
	out := new(RemoveMissionCommandItemResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/RemoveMissionCommandItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) AddCoalitionCommand(ctx context.Context, in *AddCoalitionCommandRequest, opts ...grpc.CallOption) (*AddCoalitionCommandResponse, error) {
	out := new(AddCoalitionCommandResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/AddCoalitionCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) AddCoalitionCommandSubMenu(ctx context.Context, in *AddCoalitionCommandSubMenuRequest, opts ...grpc.CallOption) (*AddCoalitionCommandSubMenuResponse, error) {
	out := new(AddCoalitionCommandSubMenuResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/AddCoalitionCommandSubMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) RemoveCoalitionCommandItem(ctx context.Context, in *RemoveCoalitionCommandItemRequest, opts ...grpc.CallOption) (*RemoveCoalitionCommandItemResponse, error) {
	out := new(RemoveCoalitionCommandItemResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/RemoveCoalitionCommandItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) AddGroupCommand(ctx context.Context, in *AddGroupCommandRequest, opts ...grpc.CallOption) (*AddGroupCommandResponse, error) {
	out := new(AddGroupCommandResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/AddGroupCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) AddGroupCommandSubMenu(ctx context.Context, in *AddGroupCommandSubMenuRequest, opts ...grpc.CallOption) (*AddGroupCommandSubMenuResponse, error) {
	out := new(AddGroupCommandSubMenuResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/AddGroupCommandSubMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) RemoveGroupCommandItem(ctx context.Context, in *RemoveGroupCommandItemRequest, opts ...grpc.CallOption) (*RemoveGroupCommandItemResponse, error) {
	out := new(RemoveGroupCommandItemResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/RemoveGroupCommandItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) GetSessionId(ctx context.Context, in *GetSessionIdRequest, opts ...grpc.CallOption) (*GetSessionIdResponse, error) {
	out := new(GetSessionIdResponse)
	err := c.cc.Invoke(ctx, "/dcs.mission.v0.MissionService/GetSessionId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MissionServiceServer is the server API for MissionService service.
// All implementations must embed UnimplementedMissionServiceServer
// for forward compatibility
type MissionServiceServer interface {
	// Streams DCS game generated Events.
	// See https://wiki.hoggitworld.com/view/Category:Events
	StreamEvents(*StreamEventsRequest, MissionService_StreamEventsServer) error
	// Streams unit updates
	// Provides similar functionality as Tacview but at a much lower update rate
	// so puts less load on the server. Suitable for things like online maps but
	// not as a Tacview replacement.
	StreamUnits(*StreamUnitsRequest, MissionService_StreamUnitsServer) error
	// Returns the mission's in-game starttime as an ISO 8601 formatted datetime
	// string.
	GetScenarioStartTime(context.Context, *GetScenarioStartTimeRequest) (*GetScenarioStartTimeResponse, error)
	// Returns the mission's in-game current time as an ISO 8601 formatted
	// datetime string.
	GetScenarioCurrentTime(context.Context, *GetScenarioCurrentTimeRequest) (*GetScenarioCurrentTimeResponse, error)
	// Adds a new mission command
	// See https://wiki.hoggitworld.com/view/DCS_func_addCommand
	AddMissionCommand(context.Context, *AddMissionCommandRequest) (*AddMissionCommandResponse, error)
	// Adds a new command sub menu
	// See https://wiki.hoggitworld.com/view/DCS_func_addSubMenu
	AddMissionCommandSubMenu(context.Context, *AddMissionCommandSubMenuRequest) (*AddMissionCommandSubMenuResponse, error)
	// Removes a registered mission command.
	// See https://wiki.hoggitworld.com/view/DCS_func_removeItem
	RemoveMissionCommandItem(context.Context, *RemoveMissionCommandItemRequest) (*RemoveMissionCommandItemResponse, error)
	// Adds a new coalition command
	// See https://wiki.hoggitworld.com/view/DCS_func_addCommandForCoalition
	AddCoalitionCommand(context.Context, *AddCoalitionCommandRequest) (*AddCoalitionCommandResponse, error)
	// Adds a new coalition command sub menu
	// See https://wiki.hoggitworld.com/view/DCS_func_addSubMenuForCoalition
	AddCoalitionCommandSubMenu(context.Context, *AddCoalitionCommandSubMenuRequest) (*AddCoalitionCommandSubMenuResponse, error)
	// Removes a registered coalition command.
	// See https://wiki.hoggitworld.com/view/DCS_func_removeItemForCoalition
	RemoveCoalitionCommandItem(context.Context, *RemoveCoalitionCommandItemRequest) (*RemoveCoalitionCommandItemResponse, error)
	// Adds a new group command
	// See https://wiki.hoggitworld.com/view/DCS_func_addCommandForGroup
	AddGroupCommand(context.Context, *AddGroupCommandRequest) (*AddGroupCommandResponse, error)
	// Adds a new group command sub menu
	// See https://wiki.hoggitworld.com/view/DCS_func_addSubMenuForGroup
	AddGroupCommandSubMenu(context.Context, *AddGroupCommandSubMenuRequest) (*AddGroupCommandSubMenuResponse, error)
	// Removes a group coalition command.
	// See https://wiki.hoggitworld.com/view/DCS_func_removeItemForGroup
	RemoveGroupCommandItem(context.Context, *RemoveGroupCommandItemRequest) (*RemoveGroupCommandItemResponse, error)
	// Returns an ID for the current session.
	// The ID will change upon mission change or server restart.
	GetSessionId(context.Context, *GetSessionIdRequest) (*GetSessionIdResponse, error)
	mustEmbedUnimplementedMissionServiceServer()
}

// UnimplementedMissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMissionServiceServer struct {
}

func (UnimplementedMissionServiceServer) StreamEvents(*StreamEventsRequest, MissionService_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (UnimplementedMissionServiceServer) StreamUnits(*StreamUnitsRequest, MissionService_StreamUnitsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamUnits not implemented")
}
func (UnimplementedMissionServiceServer) GetScenarioStartTime(context.Context, *GetScenarioStartTimeRequest) (*GetScenarioStartTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScenarioStartTime not implemented")
}
func (UnimplementedMissionServiceServer) GetScenarioCurrentTime(context.Context, *GetScenarioCurrentTimeRequest) (*GetScenarioCurrentTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScenarioCurrentTime not implemented")
}
func (UnimplementedMissionServiceServer) AddMissionCommand(context.Context, *AddMissionCommandRequest) (*AddMissionCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMissionCommand not implemented")
}
func (UnimplementedMissionServiceServer) AddMissionCommandSubMenu(context.Context, *AddMissionCommandSubMenuRequest) (*AddMissionCommandSubMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMissionCommandSubMenu not implemented")
}
func (UnimplementedMissionServiceServer) RemoveMissionCommandItem(context.Context, *RemoveMissionCommandItemRequest) (*RemoveMissionCommandItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMissionCommandItem not implemented")
}
func (UnimplementedMissionServiceServer) AddCoalitionCommand(context.Context, *AddCoalitionCommandRequest) (*AddCoalitionCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoalitionCommand not implemented")
}
func (UnimplementedMissionServiceServer) AddCoalitionCommandSubMenu(context.Context, *AddCoalitionCommandSubMenuRequest) (*AddCoalitionCommandSubMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoalitionCommandSubMenu not implemented")
}
func (UnimplementedMissionServiceServer) RemoveCoalitionCommandItem(context.Context, *RemoveCoalitionCommandItemRequest) (*RemoveCoalitionCommandItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCoalitionCommandItem not implemented")
}
func (UnimplementedMissionServiceServer) AddGroupCommand(context.Context, *AddGroupCommandRequest) (*AddGroupCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupCommand not implemented")
}
func (UnimplementedMissionServiceServer) AddGroupCommandSubMenu(context.Context, *AddGroupCommandSubMenuRequest) (*AddGroupCommandSubMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupCommandSubMenu not implemented")
}
func (UnimplementedMissionServiceServer) RemoveGroupCommandItem(context.Context, *RemoveGroupCommandItemRequest) (*RemoveGroupCommandItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupCommandItem not implemented")
}
func (UnimplementedMissionServiceServer) GetSessionId(context.Context, *GetSessionIdRequest) (*GetSessionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionId not implemented")
}
func (UnimplementedMissionServiceServer) mustEmbedUnimplementedMissionServiceServer() {}

// UnsafeMissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MissionServiceServer will
// result in compilation errors.
type UnsafeMissionServiceServer interface {
	mustEmbedUnimplementedMissionServiceServer()
}

func RegisterMissionServiceServer(s grpc.ServiceRegistrar, srv MissionServiceServer) {
	s.RegisterService(&MissionService_ServiceDesc, srv)
}

func _MissionService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MissionServiceServer).StreamEvents(m, &missionServiceStreamEventsServer{stream})
}

type MissionService_StreamEventsServer interface {
	Send(*StreamEventsResponse) error
	grpc.ServerStream
}

type missionServiceStreamEventsServer struct {
	grpc.ServerStream
}

func (x *missionServiceStreamEventsServer) Send(m *StreamEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MissionService_StreamUnits_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamUnitsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MissionServiceServer).StreamUnits(m, &missionServiceStreamUnitsServer{stream})
}

type MissionService_StreamUnitsServer interface {
	Send(*StreamUnitsResponse) error
	grpc.ServerStream
}

type missionServiceStreamUnitsServer struct {
	grpc.ServerStream
}

func (x *missionServiceStreamUnitsServer) Send(m *StreamUnitsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MissionService_GetScenarioStartTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScenarioStartTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).GetScenarioStartTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/GetScenarioStartTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).GetScenarioStartTime(ctx, req.(*GetScenarioStartTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_GetScenarioCurrentTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScenarioCurrentTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).GetScenarioCurrentTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/GetScenarioCurrentTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).GetScenarioCurrentTime(ctx, req.(*GetScenarioCurrentTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_AddMissionCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMissionCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).AddMissionCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/AddMissionCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).AddMissionCommand(ctx, req.(*AddMissionCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_AddMissionCommandSubMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMissionCommandSubMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).AddMissionCommandSubMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/AddMissionCommandSubMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).AddMissionCommandSubMenu(ctx, req.(*AddMissionCommandSubMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_RemoveMissionCommandItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMissionCommandItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).RemoveMissionCommandItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/RemoveMissionCommandItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).RemoveMissionCommandItem(ctx, req.(*RemoveMissionCommandItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_AddCoalitionCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCoalitionCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).AddCoalitionCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/AddCoalitionCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).AddCoalitionCommand(ctx, req.(*AddCoalitionCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_AddCoalitionCommandSubMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCoalitionCommandSubMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).AddCoalitionCommandSubMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/AddCoalitionCommandSubMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).AddCoalitionCommandSubMenu(ctx, req.(*AddCoalitionCommandSubMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_RemoveCoalitionCommandItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCoalitionCommandItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).RemoveCoalitionCommandItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/RemoveCoalitionCommandItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).RemoveCoalitionCommandItem(ctx, req.(*RemoveCoalitionCommandItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_AddGroupCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).AddGroupCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/AddGroupCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).AddGroupCommand(ctx, req.(*AddGroupCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_AddGroupCommandSubMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupCommandSubMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).AddGroupCommandSubMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/AddGroupCommandSubMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).AddGroupCommandSubMenu(ctx, req.(*AddGroupCommandSubMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_RemoveGroupCommandItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupCommandItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).RemoveGroupCommandItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/RemoveGroupCommandItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).RemoveGroupCommandItem(ctx, req.(*RemoveGroupCommandItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_GetSessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).GetSessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.mission.v0.MissionService/GetSessionId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).GetSessionId(ctx, req.(*GetSessionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MissionService_ServiceDesc is the grpc.ServiceDesc for MissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcs.mission.v0.MissionService",
	HandlerType: (*MissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScenarioStartTime",
			Handler:    _MissionService_GetScenarioStartTime_Handler,
		},
		{
			MethodName: "GetScenarioCurrentTime",
			Handler:    _MissionService_GetScenarioCurrentTime_Handler,
		},
		{
			MethodName: "AddMissionCommand",
			Handler:    _MissionService_AddMissionCommand_Handler,
		},
		{
			MethodName: "AddMissionCommandSubMenu",
			Handler:    _MissionService_AddMissionCommandSubMenu_Handler,
		},
		{
			MethodName: "RemoveMissionCommandItem",
			Handler:    _MissionService_RemoveMissionCommandItem_Handler,
		},
		{
			MethodName: "AddCoalitionCommand",
			Handler:    _MissionService_AddCoalitionCommand_Handler,
		},
		{
			MethodName: "AddCoalitionCommandSubMenu",
			Handler:    _MissionService_AddCoalitionCommandSubMenu_Handler,
		},
		{
			MethodName: "RemoveCoalitionCommandItem",
			Handler:    _MissionService_RemoveCoalitionCommandItem_Handler,
		},
		{
			MethodName: "AddGroupCommand",
			Handler:    _MissionService_AddGroupCommand_Handler,
		},
		{
			MethodName: "AddGroupCommandSubMenu",
			Handler:    _MissionService_AddGroupCommandSubMenu_Handler,
		},
		{
			MethodName: "RemoveGroupCommandItem",
			Handler:    _MissionService_RemoveGroupCommandItem_Handler,
		},
		{
			MethodName: "GetSessionId",
			Handler:    _MissionService_GetSessionId_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _MissionService_StreamEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamUnits",
			Handler:       _MissionService_StreamUnits_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dcs/mission/v0/mission.proto",
}
