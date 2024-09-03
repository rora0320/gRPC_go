package paseto

import (
	"gRPC/config"
	auth "gRPC/gRPC/proto"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

// CreateNewToken 토큰 생성 proto를 통해서
func (m *PasetoMaker) CreateNewToken(clientAuth *auth.AuthData) (string, error) {
	return "", nil
}

// VerifyToken 토큰 검증
func (m *PasetoMaker) VerifyToken(token string) error {
	return nil
}
