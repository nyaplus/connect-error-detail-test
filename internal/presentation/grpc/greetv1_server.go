package grpc

import (
	"context"
	"fmt"

	greetv1 "github.com/2yanpath/connect-error-detail-test/proto/greet/v1"
	"github.com/2yanpath/connect-error-detail-test/proto/greet/v1/greetv1connect"
	"github.com/bufbuild/connect-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type greetV1Server struct {
}

// NewGreetV1Server
func NewGreetV1Server() greetv1connect.GreetServiceHandler {
	return &greetV1Server{}
}

// Greet
// grpcurl -plaintext -d '{"name": ""}' localhost:8082 greet.v1.GreetService.Greet
func (s *greetV1Server) Greet(_ context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {

	if req.Msg.Name == "" {
		badRequestDetail := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "username",
					Description: "should not empty",
				},
			},
		}
		newErr := connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("name is required"))
		if detail, detailErr := connect.NewErrorDetail(badRequestDetail); detailErr == nil {
			newErr.AddDetail(detail)
		}
		return nil, newErr
	}

	greetResp := &greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	}
	resp := connect.NewResponse(greetResp)

	return resp, nil
}
