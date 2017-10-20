package uds

import (
	"fmt"
	"net"
	"os"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/health"
)

func New(path string, register func(s *grpc.Server), interceptors ...grpc.UnaryServerInterceptor) (*UnixServer, error) {
	os.Remove(path)
	lis, err := net.Listen("unix", path)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(recursiveInterceptor(interceptors...)))
	//	healthapi.RegisterHealthServer(s, health.NewServer())
	register(s)
	return &UnixServer{lis: lis, server: s}, nil
}

type UnixServer struct {
	lis    net.Listener
	server *grpc.Server
}

func (uds *UnixServer) Serve() error {
	err := uds.server.Serve(uds.lis)
	if err != nil {
		glog.Errorf("server stopped: %v", err)
	}
	return err
}

func (uds *UnixServer) Stop() {
	uds.server.Stop()
}
