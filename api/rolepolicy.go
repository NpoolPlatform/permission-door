package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/permission-door/message/npool"
	myconst "github.com/NpoolPlatform/permission-door/pkg/message/const"
	rolepolicy "github.com/NpoolPlatform/permission-door/pkg/middleware/role-policy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SetRolePolicies(ctx context.Context, in *npool.SetRolePoliciesRequest) (*npool.SetRolePoliciesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := rolepolicy.SetRolePolicies(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to set role policies: %v", err)
		return &npool.SetRolePoliciesResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) GetRolePolicies(ctx context.Context, in *npool.GetRolePoliciesRequest) (*npool.GetRolePoliciesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := rolepolicy.GetRolePolicies(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get role policies: %v", err)
		return &npool.GetRolePoliciesResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) AuthenticateRolePolicy(ctx context.Context, in *npool.AuthenticateRolePolicyRequest) (*npool.AuthenticateRolePolicyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := rolepolicy.AuthenticateRolePolicy(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to authenticate role policy: %v", err)
		return &npool.AuthenticateRolePolicyResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) AuthenticateRolesPolicy(ctx context.Context, in *npool.AuthenticateRolesPolicyRequest) (*npool.AuthenticateRolesPolicyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := rolepolicy.AuthenticateRolesPolicy(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to authenticate roles policy: %v", err)
		return &npool.AuthenticateRolesPolicyResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UnsetRolePolicies(ctx context.Context, in *npool.UnsetRolePoliciesRequest) (*npool.UnsetRolePoliciesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := rolepolicy.UnsetRolePolicies(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to unset role policies: %v", err)
		return &npool.UnsetRolePoliciesResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) DeleteRole(ctx context.Context, in *npool.DeleteRoleRequest) (*npool.DeleteRoleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := rolepolicy.DeleteRole(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to delete role: %v", err)
		return &npool.DeleteRoleResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
