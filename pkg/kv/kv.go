package kv

import (
	"context"
	"strings"

	"github.com/ngaut/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Filter proxy metadata from proxy
type Filter struct {
}

func (f *Filter) KVGet(ctx context.Context, fullMethodName string, codec grpc.Codec) (context.Context, *grpc.ClientConn, error) {
	if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
		return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
	}
	md, ok := metadata.FromIncomingContext(ctx)
	// Copy the inbound metadata explicitly.
	outCtx, _ := context.WithCancel(ctx)
	outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())
	if ok {
		// Decide on which backend to dial
		log.Infof("%v", md)
		if val, exists := md[":authority"]; exists && val[0] == "staging.api.example.com" {
			// Make sure we use DialContext so the dialing can be cancelled/time out together with the context.
			conn, err := grpc.DialContext(ctx, "api-service.staging.svc.local", grpc.WithCodec(codec))
			return outCtx, conn, err
		} else if val, exists := md[":authority"]; exists && val[0] == "api.example.com" {
			conn, err := grpc.DialContext(ctx, "api-service.prod.svc.local", grpc.WithCodec(codec))
			return outCtx, conn, err
		}
	}
	return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
}
