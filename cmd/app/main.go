package main

import (
	"net/http"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/2yanpath/connect-error-detail-test/internal/presentation/grpc"
	"github.com/2yanpath/connect-error-detail-test/proto/greet/v1/greetv1connect"
)

func main() {
	greeter := grpc.NewGreetV1Server()
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	reflector := grpcreflect.NewStaticReflector(
		greetv1connect.GreetServiceName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	http.ListenAndServe(
		"localhost:8082",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
