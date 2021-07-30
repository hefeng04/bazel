
package echoservice

import (
	"context"
	"net"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	myecho "go/grpcgateway/proto/echoservice"
)

// Implements of EchoServiceServer

type echoServer struct{}

func newEchoServer() myecho.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *myecho.SimpleMessage) (*myecho.SimpleMessage, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *echoServer) EchoBody(ctx context.Context, msg *myecho.SimpleMessage) (*myecho.SimpleMessage, error) {
	glog.Info(msg)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return msg, nil
}

func (s *echoServer) EchoDelete(ctx context.Context, msg *myecho.SimpleMessage) (*myecho.SimpleMessage, error) {
	glog.Info(msg)
	return msg, nil
}

func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	myecho.RegisterEchoServiceServer(s, newEchoServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}
