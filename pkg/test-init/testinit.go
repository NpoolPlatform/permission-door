package testinit

import (
	"fmt"
	"path"
	"runtime"

	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"
	servicename "github.com/NpoolPlatform/permission-door/pkg/service-name"

	//nolint
	"golang.org/x/xerrors"
)

func Init() error {
	_, myPath, _, ok := runtime.Caller(1)
	if !ok {
		return xerrors.Errorf("cannot get source file path")
	}

	appName := path.Base(path.Dir(path.Dir(path.Dir(myPath))))
	configPath := fmt.Sprintf("%s/../../cmd/%v", path.Dir(myPath), appName)
	fmt.Println("config path is", configPath)

	err := app.Init(servicename.ServiceName, "", "", "", configPath, nil, nil,
		config.ServiceNameToNamespace(mysqlconst.MysqlServiceName),
		config.ServiceNameToNamespace(redisconst.RedisServiceName),
		config.ServiceNameToNamespace(rabbitmqconst.RabbitMQServiceName),
		config.ServiceNameToNamespace(applicationconst.ServiceName))
	if err != nil {
		return xerrors.Errorf("cannot init app stub: %v", err)
	}

	return nil
}
