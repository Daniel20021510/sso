package auth

import (
	"context"
	ssov1 "github.com/Daniel20021510/sso-proto/gen/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Register(ctx context.Context, request *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if err := validateRegisterRequest(request); err != nil {
		return nil, err
	}

	uid, err := s.auth.Register(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		//if errors.Is(err, storage.ErrUserExists) {
		//	return nil, status.Error(codes.AlreadyExists, "user already exists")
		//}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &ssov1.RegisterResponse{UserId: uid}, nil
}

func validateRegisterRequest(request *ssov1.RegisterRequest) error {
	if request.Email == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if request.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}
