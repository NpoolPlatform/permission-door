package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/permission-door/message/npool"
	myconst "github.com/NpoolPlatform/permission-door/pkg/message/const"
	"github.com/NpoolPlatform/permission-door/pkg/middleware/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteResource(ctx context.Context, in *npool.DeleteResourceRequest) (*npool.DeleteResourceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := resource.DeleteResource(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to delete resource: %v", err)
		return &npool.DeleteResourceResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
