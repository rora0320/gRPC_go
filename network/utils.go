package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// verifyLogin 미들웨어
func (n *Network) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//1. Bearer token 가져오기
		BearerToken := getAuthToken(c)

		// 토큰 없으면 권한없음으로 거절
		if BearerToken == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		} else {
			// 토큰 있어도 토큰에 문제가 있으면 권한없음으로 거정
			if _, err := n.gRPCClient.VerifyAuth(BearerToken); err != nil {
				c.JSON(http.StatusUnauthorized, err)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}

// getAuthToken 토큰 가져오는 함수
func getAuthToken(c *gin.Context) string {
	var token string

	authToken := c.Request.Header.Get("Authorization")
	//Bearer ~~~~~ // split 결과 [Bearer,~~~~~~]
	authSliced := strings.Split(authToken, " ")

	if len(authSliced) > 1 {
		token = authSliced[1]
	}

	return token
}
