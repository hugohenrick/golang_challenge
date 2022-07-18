// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: coin.proto

package proto

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

// CoinServiceClient is the client API for CoinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoinServiceClient interface {
	CreateCoin(ctx context.Context, in *Coin, opts ...grpc.CallOption) (*CoinId, error)
	ReadCoin(ctx context.Context, in *CoinId, opts ...grpc.CallOption) (*Coin, error)
	UpdateCoin(ctx context.Context, in *Coin, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCoin(ctx context.Context, in *CoinId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListCoins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CoinService_ListCoinsClient, error)
}

type coinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinServiceClient(cc grpc.ClientConnInterface) CoinServiceClient {
	return &coinServiceClient{cc}
}

func (c *coinServiceClient) CreateCoin(ctx context.Context, in *Coin, opts ...grpc.CallOption) (*CoinId, error) {
	out := new(CoinId)
	err := c.cc.Invoke(ctx, "/coin.CoinService/CreateCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinServiceClient) ReadCoin(ctx context.Context, in *CoinId, opts ...grpc.CallOption) (*Coin, error) {
	out := new(Coin)
	err := c.cc.Invoke(ctx, "/coin.CoinService/ReadCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinServiceClient) UpdateCoin(ctx context.Context, in *Coin, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/coin.CoinService/UpdateCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinServiceClient) DeleteCoin(ctx context.Context, in *CoinId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/coin.CoinService/DeleteCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinServiceClient) ListCoins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CoinService_ListCoinsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoinService_ServiceDesc.Streams[0], "/coin.CoinService/ListCoins", opts...)
	if err != nil {
		return nil, err
	}
	x := &coinServiceListCoinsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CoinService_ListCoinsClient interface {
	Recv() (*Coin, error)
	grpc.ClientStream
}

type coinServiceListCoinsClient struct {
	grpc.ClientStream
}

func (x *coinServiceListCoinsClient) Recv() (*Coin, error) {
	m := new(Coin)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CoinServiceServer is the server API for CoinService service.
// All implementations must embed UnimplementedCoinServiceServer
// for forward compatibility
type CoinServiceServer interface {
	CreateCoin(context.Context, *Coin) (*CoinId, error)
	ReadCoin(context.Context, *CoinId) (*Coin, error)
	UpdateCoin(context.Context, *Coin) (*emptypb.Empty, error)
	DeleteCoin(context.Context, *CoinId) (*emptypb.Empty, error)
	ListCoins(*emptypb.Empty, CoinService_ListCoinsServer) error
	mustEmbedUnimplementedCoinServiceServer()
}

// UnimplementedCoinServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoinServiceServer struct {
}

func (UnimplementedCoinServiceServer) CreateCoin(context.Context, *Coin) (*CoinId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoin not implemented")
}
func (UnimplementedCoinServiceServer) ReadCoin(context.Context, *CoinId) (*Coin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCoin not implemented")
}
func (UnimplementedCoinServiceServer) UpdateCoin(context.Context, *Coin) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoin not implemented")
}
func (UnimplementedCoinServiceServer) DeleteCoin(context.Context, *CoinId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoin not implemented")
}
func (UnimplementedCoinServiceServer) ListCoins(*emptypb.Empty, CoinService_ListCoinsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCoins not implemented")
}
func (UnimplementedCoinServiceServer) mustEmbedUnimplementedCoinServiceServer() {}

// UnsafeCoinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoinServiceServer will
// result in compilation errors.
type UnsafeCoinServiceServer interface {
	mustEmbedUnimplementedCoinServiceServer()
}

func RegisterCoinServiceServer(s grpc.ServiceRegistrar, srv CoinServiceServer) {
	s.RegisterService(&CoinService_ServiceDesc, srv)
}

func _CoinService_CreateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Coin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServiceServer).CreateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coin.CoinService/CreateCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServiceServer).CreateCoin(ctx, req.(*Coin))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinService_ReadCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServiceServer).ReadCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coin.CoinService/ReadCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServiceServer).ReadCoin(ctx, req.(*CoinId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinService_UpdateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Coin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServiceServer).UpdateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coin.CoinService/UpdateCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServiceServer).UpdateCoin(ctx, req.(*Coin))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinService_DeleteCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServiceServer).DeleteCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coin.CoinService/DeleteCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServiceServer).DeleteCoin(ctx, req.(*CoinId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinService_ListCoins_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoinServiceServer).ListCoins(m, &coinServiceListCoinsServer{stream})
}

type CoinService_ListCoinsServer interface {
	Send(*Coin) error
	grpc.ServerStream
}

type coinServiceListCoinsServer struct {
	grpc.ServerStream
}

func (x *coinServiceListCoinsServer) Send(m *Coin) error {
	return x.ServerStream.SendMsg(m)
}

// CoinService_ServiceDesc is the grpc.ServiceDesc for CoinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coin.CoinService",
	HandlerType: (*CoinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoin",
			Handler:    _CoinService_CreateCoin_Handler,
		},
		{
			MethodName: "ReadCoin",
			Handler:    _CoinService_ReadCoin_Handler,
		},
		{
			MethodName: "UpdateCoin",
			Handler:    _CoinService_UpdateCoin_Handler,
		},
		{
			MethodName: "DeleteCoin",
			Handler:    _CoinService_DeleteCoin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCoins",
			Handler:       _CoinService_ListCoins_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "coin.proto",
}
