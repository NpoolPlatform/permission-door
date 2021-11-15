package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/poolPlatform/permission-door/message/npool"
	"github.com/poolPlatform/permission-door/pkg/middleware/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteResource(ctx context.Context, in *npool.DeleteResourceRequest) (*npool.DeleteResourceResponse, error) {
	resp, err := resource.DeleteResource(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to delete resource: %v", err)
		return &npool.DeleteResourceResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
