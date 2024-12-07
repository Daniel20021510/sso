package auth

import (
	"context"
	ssov1 "github.com/Daniel20021510/sso-proto/gen/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Login(ctx context.Context, request *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if err := validateLoginRequest(request); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, request.GetEmail(), request.GetPassword(), request.GetAppId())
	if err != nil {
		//if errors.Is(err, auth.ErrInvalidCredentials) {
		//	return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		//}

		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &ssov1.LoginResponse{Token: token}, nil
}

func validateLoginRequest(request *ssov1.LoginRequest) error {
	if request.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if request.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	if request.GetAppId() == 0 {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}

	return nil
}
