package grpc

import (
	"context"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
)

func GetUserRoleIDs(userID, appID string) ([]string, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

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
