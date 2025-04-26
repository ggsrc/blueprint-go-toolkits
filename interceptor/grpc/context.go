package grpcinterceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ContextUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// You can extract metadata from incoming context
		// set request source
		// md := metautils.ExtractIncoming(ctx)

		// origin := md.Get("origin")
		// //nolint:golint,staticcheck
		// ctx = context.WithValue(ctx, "origin", origin)

		ret, err := handler(ctx, req)
		return ret, err
	}
}

func ContextUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			md = md.Copy()
		} else {
			md = metadata.MD{}
		}
		outgoingmd, ok := metadata.FromOutgoingContext(ctx)
		if ok {
			// explicitly declared outgoing md take precedence over transitive incoming md
			md = metadata.Join(outgoingmd, md)
		}
		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func NewContextWithGRPCMeta(ctx context.Context) context.Context {
	newCtx := context.Background()
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		newCtx = metadata.NewIncomingContext(newCtx, md)
	}
	md, ok = metadata.FromOutgoingContext(ctx)
	if ok {
		newCtx = metadata.NewOutgoingContext(newCtx, md)
	}
	return newCtx
}
