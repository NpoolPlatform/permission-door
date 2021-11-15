package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/permission-door/message/npool"
	"github.com/NpoolPlatform/permission-door/pkg/middleware/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthenticateUserPolicyByID(ctx context.Context, in *npool.AuthenticateUserPolicyByIDRequest) (*npool.AuthenticateUserPolicyByIDResponse, error) {
	resp, err := user.AuthenticateUserPolicyByID(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to authenticate user policy by user id: %v", err)
		return &npool.AuthenticateUserPolicyByIDResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
