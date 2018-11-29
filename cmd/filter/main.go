package filter

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/cwen0/filter/proxy"
)

var (
	upstream   string
	listenAddr string
)

func init() {
	flag.StringVar(&upstream, "upstream", "127.0.0.1:11000", "upstream port")
	flag.StringVar(&listenAddr, "listen-addr", ":10000", "serve port")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ctx := context.Background()

	cfg := make(map[string]string)
	cfg["/helloworld.Greeter/SayHello"] = "rand(5)->delay(1000)|rand(1)->timeout()"
	proxyHandler, err := proxy.NewProxyHandler(ctx, cfg, upstream)
	if err != nil {
		log.Fatalf("failed to setup proxy: %v", err)
	}
	s := grpc.NewServer(grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxyHandler.StreamHandler()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
