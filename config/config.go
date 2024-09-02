// 환경 변수 관리

package config

import (
	"github.com/naoina/toml"
	"os"
)

// Config 구조체 -> 파싱할 toml 파일의 구조를 따라감
type Config struct{}

// NewConfig path 각 서비스에 필요한 환경변수 저장된 파일 경로
func NewConfig(path string) *Config {
	c := new(Config)
	//nil -> null(zero value) 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수타입의 zero value
	//go에서 제공되는 오픈패키지로 해당 경로의 파일을 열고
	if file, err := os.Open(path); err != nil {
		panic(err)
	} else {
		defer file.Close()

		if err = toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
