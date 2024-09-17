package controllers

import (
	"fmt"
	"ldapctl/g"
	"ldapctl/models"

	"github.com/gin-gonic/gin"
)

func AuthSingle(c *gin.Context) {
	var user models.User
	var msgResult MsgResult

	err := c.ShouldBindJSON(&user)
	if err != nil {
		msgResult.Msg = err.Error()
		c.JSON(400, msgResult)
	}

	if user.Username == "" || user.Password == "" {
		msgResult.Msg = "Missing username or password"
		c.JSON(401, msgResult)
	}

	success, err := models.SingleAuth(g.Config().Ldap, user.Username, user.Password)
	if err != nil {
		msgResult.Msg = err.Error()

	} else {
		msgResult.Msg = fmt.Sprintf("user %s Auth Successed", user.Username)
	}

	msgResult.Success = success
	c.JSON(200, msgResult)
}
