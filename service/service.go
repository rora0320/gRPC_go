// Package service 라우터 <-> 리포지토리 연결하는 역할
// 라우터-> 서비스 ->리포지토리에 요청-> 리포지토리 응답값을 서비스로 보내서 서비스에서 가공 -> 라우터로 가공 응답값 보냄

package service

import (
	"gRPC/config"
	"gRPC/repository"
)

type Service struct {
	cfg        *config.Config
	repository *repository.Repository
}

// NewService 서비스->리포지토리
func NewService(cfg *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{cfg: cfg, repository: repository}

	return r, nil
}
