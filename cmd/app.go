// 리포지토리, 서비스, 네트워크에 대한 객체 값을 가지고 있음

package cmd

import (
	"gRPC/config"
	"gRPC/network"
	"gRPC/repository"
	"gRPC/service"
)

type App struct {
	cfg *config.Config

	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

// NewApp 리포지토리, 서비스, 네트워크에 대한 객체값이 필요해- 앱
func NewApp(cfg *config.Config) {

	var err error

	a := &App{cfg: cfg}

	//리포지토리 - 디비연결 설정
	if a.repository, err = repository.NewRepository(cfg); err != nil {
		panic(err)

		//서비스 - 리포지토리로 연결
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)

		//네트워크(라우터) - 서비스로 연결
	} else if a.network, err = network.NewNetwork(cfg, a.service); err != nil {
		panic(err)
	} else {
		//TODO -> start server
		a.network.StartServer()
	}

}
