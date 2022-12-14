// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: interview.proto

package interview

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

// LogicalInterviewClient is the client API for LogicalInterview service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogicalInterviewClient interface {
	GetResult(ctx context.Context, in *User, opts ...grpc.CallOption) (*Result, error)
	GetNextQuestion(ctx context.Context, in *User, opts ...grpc.CallOption) (*Question, error)
	SendAnswer(ctx context.Context, in *Ansver, opts ...grpc.CallOption) (*AnswerResponce, error)
	GetSummary(ctx context.Context, in *User, opts ...grpc.CallOption) (*Summary, error)
}

type logicalInterviewClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicalInterviewClient(cc grpc.ClientConnInterface) LogicalInterviewClient {
	return &logicalInterviewClient{cc}
}

func (c *logicalInterviewClient) GetResult(ctx context.Context, in *User, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/logicalInterview.LogicalInterview/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicalInterviewClient) GetNextQuestion(ctx context.Context, in *User, opts ...grpc.CallOption) (*Question, error) {
	out := new(Question)
	err := c.cc.Invoke(ctx, "/logicalInterview.LogicalInterview/GetNextQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicalInterviewClient) SendAnswer(ctx context.Context, in *Ansver, opts ...grpc.CallOption) (*AnswerResponce, error) {
	out := new(AnswerResponce)
	err := c.cc.Invoke(ctx, "/logicalInterview.LogicalInterview/SendAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicalInterviewClient) GetSummary(ctx context.Context, in *User, opts ...grpc.CallOption) (*Summary, error) {
	out := new(Summary)
	err := c.cc.Invoke(ctx, "/logicalInterview.LogicalInterview/GetSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicalInterviewServer is the server API for LogicalInterview service.
// All implementations must embed UnimplementedLogicalInterviewServer
// for forward compatibility
type LogicalInterviewServer interface {
	GetResult(context.Context, *User) (*Result, error)
	GetNextQuestion(context.Context, *User) (*Question, error)
	SendAnswer(context.Context, *Ansver) (*AnswerResponce, error)
	GetSummary(context.Context, *User) (*Summary, error)
	mustEmbedUnimplementedLogicalInterviewServer()
}

// UnimplementedLogicalInterviewServer must be embedded to have forward compatible implementations.
type UnimplementedLogicalInterviewServer struct {
}

func (UnimplementedLogicalInterviewServer) GetResult(context.Context, *User) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedLogicalInterviewServer) GetNextQuestion(context.Context, *User) (*Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextQuestion not implemented")
}
func (UnimplementedLogicalInterviewServer) SendAnswer(context.Context, *Ansver) (*AnswerResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAnswer not implemented")
}
func (UnimplementedLogicalInterviewServer) GetSummary(context.Context, *User) (*Summary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummary not implemented")
}
func (UnimplementedLogicalInterviewServer) mustEmbedUnimplementedLogicalInterviewServer() {}

// UnsafeLogicalInterviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogicalInterviewServer will
// result in compilation errors.
type UnsafeLogicalInterviewServer interface {
	mustEmbedUnimplementedLogicalInterviewServer()
}

func RegisterLogicalInterviewServer(s grpc.ServiceRegistrar, srv LogicalInterviewServer) {
	s.RegisterService(&LogicalInterview_ServiceDesc, srv)
}

func _LogicalInterview_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicalInterviewServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logicalInterview.LogicalInterview/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicalInterviewServer).GetResult(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicalInterview_GetNextQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicalInterviewServer).GetNextQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logicalInterview.LogicalInterview/GetNextQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicalInterviewServer).GetNextQuestion(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicalInterview_SendAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ansver)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicalInterviewServer).SendAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logicalInterview.LogicalInterview/SendAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicalInterviewServer).SendAnswer(ctx, req.(*Ansver))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicalInterview_GetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicalInterviewServer).GetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logicalInterview.LogicalInterview/GetSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicalInterviewServer).GetSummary(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// LogicalInterview_ServiceDesc is the grpc.ServiceDesc for LogicalInterview service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogicalInterview_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logicalInterview.LogicalInterview",
	HandlerType: (*LogicalInterviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResult",
			Handler:    _LogicalInterview_GetResult_Handler,
		},
		{
			MethodName: "GetNextQuestion",
			Handler:    _LogicalInterview_GetNextQuestion_Handler,
		},
		{
			MethodName: "SendAnswer",
			Handler:    _LogicalInterview_SendAnswer_Handler,
		},
		{
			MethodName: "GetSummary",
			Handler:    _LogicalInterview_GetSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interview.proto",
}
