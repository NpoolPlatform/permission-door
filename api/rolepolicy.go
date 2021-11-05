package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/poolPlatform/permission-door/message/npool"
	"github.com/poolPlatform/permission-door/pkg/casbin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SetRolePolicies(ctx context.Context, in *npool.SetRolePoliciesRequest) (*npool.SetRolePoliciesResponse, error) {
	for _, policy := range in.ResourcePolicies {
		_, err := casbin.Enforcer().AddPolicy(in.RoleId, in.AppId, policy.ResourceId, policy.Action)
		if err != nil {
			logger.Sugar().Error("set role policy error: %w", err)
			return &npool.SetRolePoliciesResponse{}, status.Error(codes.Internal, "internal server error")
		}
	}
	return &npool.SetRolePoliciesResponse{
		Info: "set policies success!",
	}, status.Error(codes.OK, codes.OK.String())
}

func (s *Server) GetRolePolicies(ctx context.Context, in *npool.GetRolePoliciesRequest) (*npool.GetRolePoliciesResponse, error) {
	fmt.Printf("***************input is: %v, role id is: %v, app id is: %v\n", in, in.RoleId, in.AppId)
	policies := casbin.Enforcer().GetFilteredPolicy(0, in.RoleId, in.AppId)

	npoolPolicy := []*npool.ResourcePolicy{}
	for _, policy := range policies {
		npoolPolicy = append(npoolPolicy, &npool.ResourcePolicy{
			ResourceId: policy[2],
			Action:     policy[3],
		})
	}
	response := &npool.GetRolePoliciesResponse{
		Infos: &npool.RolePolicies{
			RoleId:   in.RoleId,
			AppId:    in.AppId,
			Policies: npoolPolicy,
		},
	}
	return response, status.Error(codes.OK, codes.OK.String())
}

func (s *Server) AuthenticateRolePolicy(ctx context.Context, in *npool.AuthenticateRolePolicyRequest) (*npool.AuthenticateRolePolicyResponse, error) {
	ok, err := casbin.Enforcer().Enforce(in.RoleId, in.AppId, in.Policy.ResourceId, in.Policy.Action)
	if err != nil {
		logger.Sugar().Errorf("authenticate role policy error: %v", err)
		return &npool.AuthenticateRolePolicyResponse{}, err
	}
	if !ok {
		return &npool.AuthenticateRolePolicyResponse{
			Info: "the role doesn't have this policy",
		}, status.Error(codes.PermissionDenied, "role doesn't have permission")
	}
	return &npool.AuthenticateRolePolicyResponse{
		Info: "ok",
	}, status.Error(codes.OK, codes.OK.String())
}

func (s *Server) AuthenticateRolesPolicy(ctx context.Context, in *npool.AuthenticateRolesPolicyRequest) (*npool.AuthenticateRolesPolicyResponse, error) {
	for _, role := range in.RoleIds {
		ok, err := casbin.Enforcer().Enforce(role, in.AppId, in.Policy.ResourceId, in.Policy.Action)
		if err != nil {
			logger.Sugar().Errorf("authenticate roles policy error: %v", err)
			return &npool.AuthenticateRolesPolicyResponse{}, err
		}
		if !ok {
			continue
		}
		return &npool.AuthenticateRolesPolicyResponse{
			Info: fmt.Sprintf("role %v has the permission", role),
		}, status.Error(codes.OK, codes.OK.String())
	}
	return &npool.AuthenticateRolesPolicyResponse{
		Info: "roles doesn't have the permission",
	}, status.Error(codes.PermissionDenied, "roles doesn't have the permission")
}

func (s *Server) UnsetRolePolicies(ctx context.Context, in *npool.UnsetRolePoliciesRequest) (*emptypb.Empty, error) {
	for _, policy := range in.Policies {
		ok, err := casbin.Enforcer().RemovePolicy(in.RoleId, in.AppId, policy.ResourceId, policy.Action)
		if err != nil {
			logger.Sugar().Errorf("unset role policy error: %v", err)
			return new(emptypb.Empty), err
		}
		if !ok {
			return new(emptypb.Empty), status.Error(codes.NotFound, codes.NotFound.String())
		}
	}
	return new(emptypb.Empty), status.Error(codes.OK, "unset successfully")
}

func (s *Server) DeleteRole(ctx context.Context, in *npool.DeleteRoleRequest) (*npool.DeleteRoleResponse, error) {
	for _, role := range in.RoleIds {
		ok, err := casbin.Enforcer().RemoveFilteredPolicy(0, role, in.AppId)
		if err != nil {
			logger.Sugar().Errorf("delete role error: %v", err)
			return &npool.DeleteRoleResponse{}, err
		}
		if !ok {
			return &npool.DeleteRoleResponse{}, status.Error(codes.Aborted, fmt.Sprintf("can not delete role %v", role))
		}
	}
	return &npool.DeleteRoleResponse{
		Info: "delete role successfully",
	}, status.Error(codes.OK, codes.OK.String())
}
