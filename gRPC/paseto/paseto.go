package paseto

import (
	"crypto/rand"
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
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	//Encrypt 암호화 ( 키, 값 , 랜덤값)
	return m.Pt.Encrypt(m.Key, clientAuth, randomBytes)
}

// VerifyToken 토큰 검증
func (m *PasetoMaker) VerifyToken(token string) error {
	var clientVerifyAuth *auth.AuthData
	//Decrypt 복호화 토큰, 키값, 페이로드
	return m.Pt.Decrypt(token, m.Key, clientVerifyAuth, nil)
}
