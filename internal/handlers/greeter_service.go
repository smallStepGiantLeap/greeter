package handlers

import (
	"context"

	greeterv1 "github.com/smallStepGiantLeap/greeter/client/gen/greeter/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GreeterService implements greeter.v1.GreeterService. Generated once by vikrant: this file belongs to the
// service team.
type GreeterService struct {
	greeterv1.UnimplementedGreeterServiceServer
}

// NewGreeterService returns the service implementation main registers.
func NewGreeterService() *GreeterService { return &GreeterService{} }

// Greet is unary and marked idempotent in the proto, so the mesh retries it
// on UNAVAILABLE. Keep it safe to run twice.
func (s *GreeterService) Greet(ctx context.Context, req *greeterv1.GreetRequest) (*greeterv1.GreetResponse, error) {
	return nil, status.Error(codes.Unimplemented, "greeter.v1.GreeterService/Greet is not implemented yet")
}

// Subscribe streams. Return when stream.Context() is done: on shutdown the platform
// cancels it and the client reconnects to another replica.
func (s *GreeterService) Subscribe(req *greeterv1.SubscribeRequest, stream grpc.ServerStreamingServer[greeterv1.GreetResponse]) error {
	return status.Error(codes.Unimplemented, "greeter.v1.GreeterService/Subscribe is not implemented yet")
}
