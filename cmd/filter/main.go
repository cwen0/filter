package main

import (
	"context"
	"flag"
	"net"

	"github.com/cwen0/filter/proxy"
	"github.com/ngaut/log"
	"google.golang.org/grpc"
)

var (
	upstream   string
	listenAddr string
	logFile    string
	logLevel   string
)

func init() {
	flag.StringVar(&upstream, "upstream", "127.0.0.1:11000", "upstream port")
	flag.StringVar(&listenAddr, "listen-addr", "127.0.0.1:11111", "serve port")
	flag.StringVar(&logFile, "log-file", "", "agent log file")
	flag.StringVar(&logLevel, "log-level", "info", "agent log level: info, warn, fatal, error")
}

func initLogger() {
	log.SetLevelByString(logLevel)

	if len(logFile) > 0 {
		log.SetOutputByName(logFile)
		log.SetRotateByDay()
	}
}

func main() {
	flag.Parse()

	initLogger()

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
