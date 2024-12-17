package interseptor

import (
	"context"
	"github.com/Daniel20021510/sso/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Logger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	raw, _ := protojson.Marshal((req).(proto.Message))
	logger.Infow(ctx, "request", "method", info.FullMethod, "req", string(raw))

	if resp, err = handler(ctx, req); err != nil {
		logger.Errorw(ctx, "response", "method", info.FullMethod, "err", err)
		return
	}

	rawResp, _ := protojson.Marshal((resp).(proto.Message))
	logger.Infow(ctx, "response", "method", info.FullMethod, "resp", string(rawResp))

	return
}
