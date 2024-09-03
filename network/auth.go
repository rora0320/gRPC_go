package network

import (
	"gRPC/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (n *Network) login(c *gin.Context) {
	var req types.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
	} else if res, err := n.service.CreateAuth(req.Name); err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (n *Network) verify(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}
