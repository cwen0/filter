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

func (f *Filter) KVGet(ctx context.Context, fullMethodName string) error {
	if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
		return grpc.Errorf(codes.Unimplemented, "Unknown method")
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// Decide on which backend to dial
		// if val, exists := md[":authority"]; exists && val[0] == "staging.api.example.com" {
		// 	// Make sure we use DialContext so the dialing can be cancelled/time out together with the context.
		// 	// return grpc.DialContext(ctx, "api-service.staging.svc.local", grpc.WithCodec(proxy.Codec()))

		// } else if val, exists := md[":authority"]; exists && val[0] == "api.example.com" {
		// 	// return grpc.DialContext(ctx, "api-service.prod.svc.local", grpc.WithCodec(proxy.Codec()))
		// }
		log.Infof("%v", md)
	}
	return grpc.Errorf(codes.Unimplemented, "Unknown method")
}
