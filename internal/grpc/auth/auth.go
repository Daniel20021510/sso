package auth

import (
	"context"
	ssov1 "github.com/Daniel20021510/sso-proto/gen/go"
	"google.golang.org/grpc"
)

var _ ssov1.AuthServer = (*server)(nil)

type Auth interface {
	Login(ctx context.Context, email string, password string, appID uint32) (token string, err error)
	Register(ctx context.Context, email string, password string) (userID uint64, err error)
}

type server struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func RegisterServer(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &server{auth: auth})
}
