package proxy

import (
	"github.com/ngaut/log"
	"google.golang.org/grpc"
)

// KVFilter proxy metadata from proxy
type KVFilter struct {
}

func (k *KVFilter) KVGet(src grpc.ServerStream, dst grpc.ClientStream) error {
	f := &frame{}
	err := src.RecvMsg(f)
	if err != nil {
		// can not use error.Trace for eof
		return err
	}

	log.Infof("%s", string(f.payload))

	return dst.SendMsg(f)
}
