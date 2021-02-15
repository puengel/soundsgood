// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package soundsgood

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

// AudioStreamClient is the client API for AudioStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioStreamClient interface {
	GetStream(ctx context.Context, opts ...grpc.CallOption) (AudioStream_GetStreamClient, error)
	GetFormat(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AudioFormat, error)
}

type audioStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioStreamClient(cc grpc.ClientConnInterface) AudioStreamClient {
	return &audioStreamClient{cc}
}

func (c *audioStreamClient) GetStream(ctx context.Context, opts ...grpc.CallOption) (AudioStream_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AudioStream_ServiceDesc.Streams[0], "/AudioStream/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioStreamGetStreamClient{stream}
	return x, nil
}

type AudioStream_GetStreamClient interface {
	Send(*SampleRequest) error
	Recv() (*AudioSample, error)
	grpc.ClientStream
}

type audioStreamGetStreamClient struct {
	grpc.ClientStream
}

func (x *audioStreamGetStreamClient) Send(m *SampleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioStreamGetStreamClient) Recv() (*AudioSample, error) {
	m := new(AudioSample)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *audioStreamClient) GetFormat(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AudioFormat, error) {
	out := new(AudioFormat)
	err := c.cc.Invoke(ctx, "/AudioStream/GetFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudioStreamServer is the server API for AudioStream service.
// All implementations must embed UnimplementedAudioStreamServer
// for forward compatibility
type AudioStreamServer interface {
	GetStream(AudioStream_GetStreamServer) error
	GetFormat(context.Context, *emptypb.Empty) (*AudioFormat, error)
	mustEmbedUnimplementedAudioStreamServer()
}

// UnimplementedAudioStreamServer must be embedded to have forward compatible implementations.
type UnimplementedAudioStreamServer struct {
}

func (UnimplementedAudioStreamServer) GetStream(AudioStream_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedAudioStreamServer) GetFormat(context.Context, *emptypb.Empty) (*AudioFormat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFormat not implemented")
}
func (UnimplementedAudioStreamServer) mustEmbedUnimplementedAudioStreamServer() {}

// UnsafeAudioStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioStreamServer will
// result in compilation errors.
type UnsafeAudioStreamServer interface {
	mustEmbedUnimplementedAudioStreamServer()
}

func RegisterAudioStreamServer(s grpc.ServiceRegistrar, srv AudioStreamServer) {
	s.RegisterService(&AudioStream_ServiceDesc, srv)
}

func _AudioStream_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioStreamServer).GetStream(&audioStreamGetStreamServer{stream})
}

type AudioStream_GetStreamServer interface {
	Send(*AudioSample) error
	Recv() (*SampleRequest, error)
	grpc.ServerStream
}

type audioStreamGetStreamServer struct {
	grpc.ServerStream
}

func (x *audioStreamGetStreamServer) Send(m *AudioSample) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioStreamGetStreamServer) Recv() (*SampleRequest, error) {
	m := new(SampleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AudioStream_GetFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioStreamServer).GetFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AudioStream/GetFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioStreamServer).GetFormat(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AudioStream_ServiceDesc is the grpc.ServiceDesc for AudioStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AudioStream",
	HandlerType: (*AudioStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFormat",
			Handler:    _AudioStream_GetFormat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _AudioStream_GetStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "soundService.proto",
}