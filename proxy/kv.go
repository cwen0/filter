package proxy

import (
	"encoding/json"
	"github.com/ngaut/log"
	"github.com/pingcap/kvproto/pkg/kvrpcpb"
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

	// log.Infof("%s", string(f.payload))
	// kvClient := pb.New
	// tikvCli := tikvpb.RegisterTikvServer()
	mes := &kvrpcpb.GetRequest{}
	if err := json.Unmarshal(f.payload, mes); err != nil {
		return err
	}
	log.Infof("%v", mes)

	return dst.SendMsg(f)
}
