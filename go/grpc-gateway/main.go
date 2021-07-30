package main


import (
	"context"
	"flag"

	"github.com/golang/glog"
	server "go/grpc-gateway/echo"
	"go/grpc-gateway/gateway"
)

var (
	addr    = flag.String("addr", ":9090", "endpoint of the gRPC service")
        endpoint   = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
        openAPIDir = flag.String("openapi_dir", "examples/internal/proto/examplepb", "path to the directory which contains OpenAPI definitions")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
        opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
		OpenAPIDir: *openAPIDir,
	}
        go func() {
          if err := gateway.Run(ctx, opts); err != nil {
            glog.Fatal(err)
          }
        }()
	if err := server.Run(ctx, *network, *addr); err != nil {
		glog.Fatal(err)
	}
}
