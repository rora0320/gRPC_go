package paseto

import (
	"gRPC/config"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

// CreateNewToken 토큰 생성
func (m *PasetoMaker) CreateNewToken() (string, error) {
	return "", nil
}

// VerifyToken 토큰 검증
func (m *PasetoMaker) VerifyToken(token string) error {
	return nil
}
