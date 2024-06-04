package auth

import (
	"context"

	authv1 "github.com/fallen-fatalist/protos/gen/go/auth-service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	authv1.UnimplementedAuthServer
}

const (
	emptyValue = 0
)

func Register(gRPC *grpc.Server) {
	authv1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *authv1.LoginRequest,
) (*authv1.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	// TODO: implement login via auth service

	return &authv1.LoginResponse{
		Token: "12345",
	}, nil
}
func (s *serverAPI) Register(
	ctx context.Context,
	req *authv1.RegisterRequest,
) (*authv1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *authv1.IsAdminRequest,
) (*authv1.IsAdminResponse, error) {
	panic("implement me")
}
