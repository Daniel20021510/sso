package auth_grpc

import (
	"context"
	ssov1 "github.com/Daniel20021510/sso-proto/gen/go"
	"google.golang.org/grpc"
)

var _ ssov1.AuthServer = (*server)(nil)

type AuthService interface {
	Login(ctx context.Context, email string, password string, appID uint32) (token string, err error)
	RegisterUser(ctx context.Context, email string, password string) (userID uint64, err error)
}

type server struct {
	ssov1.UnimplementedAuthServer
	authService AuthService
}

func RegisterServer(gRPCServer *grpc.Server, authService AuthService) {
	ssov1.RegisterAuthServer(gRPCServer, &server{authService: authService})
}
