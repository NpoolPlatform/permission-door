package grpc

import (
	"context"
	"strings"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
)

const (
	ApplicationService         = "application-management.npool.top"
	ApplicationServiceGRPCPort = "50081"
)

func newApplicationGrpcClient() (*grpc.ClientConn, error) {
	serviceAgent, err := config.PeekService(ApplicationService)
	if err != nil {
		return nil, err
	}

	myAddress := []string{}
	for _, address := range strings.Split(serviceAgent.Address, ",") {
		myAddress = append(myAddress, address+":50081")
	}

	conn, err := mygrpc.GetGRPCConn(strings.Join(myAddress, ","))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetUserRoleIDs(userID, appID string) ([]string, error) {
	conn, err := newApplicationGrpcClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbApplication.NewApplicationManagementClient(conn)
	resp, err := client.GetUserRole(context.Background(), &pbApplication.GetUserRoleRequest{
		UserID: userID,
		AppID:  appID,
	})
	if err != nil {
		return nil, err
	}

	roleIDs := []string{}
	for _, info := range resp.Info.Infos {
		roleIDs = append(roleIDs, info.ID)
	}
	return roleIDs, nil
}
