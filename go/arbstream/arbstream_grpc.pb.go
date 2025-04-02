// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: arbstream.proto

package arbstream

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
	ArbStreamServer_StreamArbData_FullMethodName = "/com.bloxroute.bsc.mev.arbstream.ArbStreamServer/StreamArbData"
)

// ArbStreamServerClient is the client API for ArbStreamServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// gRPC service definition
type ArbStreamServerClient interface {
	StreamArbData(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Order], error)
}

type arbStreamServerClient struct {
	cc grpc.ClientConnInterface
}

func NewArbStreamServerClient(cc grpc.ClientConnInterface) ArbStreamServerClient {
	return &arbStreamServerClient{cc}
}

func (c *arbStreamServerClient) StreamArbData(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Order], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ArbStreamServer_ServiceDesc.Streams[0], ArbStreamServer_StreamArbData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamRequest, Order]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ArbStreamServer_StreamArbDataClient = grpc.ServerStreamingClient[Order]

// ArbStreamServerServer is the server API for ArbStreamServer service.
// All implementations must embed UnimplementedArbStreamServerServer
// for forward compatibility.
//
// gRPC service definition
type ArbStreamServerServer interface {
	StreamArbData(*StreamRequest, grpc.ServerStreamingServer[Order]) error
	mustEmbedUnimplementedArbStreamServerServer()
}

// UnimplementedArbStreamServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedArbStreamServerServer struct{}

func (UnimplementedArbStreamServerServer) StreamArbData(*StreamRequest, grpc.ServerStreamingServer[Order]) error {
	return status.Errorf(codes.Unimplemented, "method StreamArbData not implemented")
}
func (UnimplementedArbStreamServerServer) mustEmbedUnimplementedArbStreamServerServer() {}
func (UnimplementedArbStreamServerServer) testEmbeddedByValue()                         {}

// UnsafeArbStreamServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArbStreamServerServer will
// result in compilation errors.
type UnsafeArbStreamServerServer interface {
	mustEmbedUnimplementedArbStreamServerServer()
}

func RegisterArbStreamServerServer(s grpc.ServiceRegistrar, srv ArbStreamServerServer) {
	// If the following call pancis, it indicates UnimplementedArbStreamServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ArbStreamServer_ServiceDesc, srv)
}

func _ArbStreamServer_StreamArbData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArbStreamServerServer).StreamArbData(m, &grpc.GenericServerStream[StreamRequest, Order]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ArbStreamServer_StreamArbDataServer = grpc.ServerStreamingServer[Order]

// ArbStreamServer_ServiceDesc is the grpc.ServiceDesc for ArbStreamServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArbStreamServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.bloxroute.bsc.mev.arbstream.ArbStreamServer",
	HandlerType: (*ArbStreamServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamArbData",
			Handler:       _ArbStreamServer_StreamArbData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "arbstream.proto",
}
