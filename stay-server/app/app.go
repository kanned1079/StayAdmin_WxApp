package app

import (
	"stay-server/dao"
	"stay-server/routers"
)

type AppInstance struct {
	InstanceId  int32
	DaoInst     *dao.DaoInstance
	GatewayInst *routers.GatewayApp
}

func NewApp(id int32, gatewayRunMode string) *AppInstance {
	return &AppInstance{
		InstanceId:  1,
		GatewayInst: routers.NewGatewayApp(id, gatewayRunMode),
		DaoInst:     dao.NewDaoInstance(id),
	}
}
