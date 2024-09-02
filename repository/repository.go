//DB에 대한 접근 설정

package repository

import "gRPC/config"

type Repository struct {
	cfg *config.Config
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}

	return r, nil
}
