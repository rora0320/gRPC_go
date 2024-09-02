// Package network 라우터 역할
// 라우터 -> 서비스를 통해 ->리포지토리에 요청-> 리포지토리 응답값을 서비스로 보내서 서비스에서 가공 -> 라우터로 가공 응답값 보냄
package network

import (
	"gRPC/config"
	"gRPC/service"
)

type Network struct {
	cfg     *config.Config
	service *service.Service
}

// NewNetwork 라우터->서비스
func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	r := &Network{cfg: cfg, service: service}

	return r, nil
}
