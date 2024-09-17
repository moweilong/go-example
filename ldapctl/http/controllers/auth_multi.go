package controllers

import (
	"ldapctl/g"
	"ldapctl/models"

	"github.com/gin-gonic/gin"
)

func AuthMulti(c *gin.Context) {
	var userList []models.User
	var msgResult MsgResult
	var authResult AuthResult

	err := c.ShouldBindJSON(&userList)
	if err != nil {
		msgResult.Msg = err.Error()
		c.JSON(400, msgResult)
		return
	}

	result, err := models.MultiAuth(g.Config().Ldap, userList)
	if err != nil {
		msgResult.Msg = err.Error()
		c.JSON(400, msgResult)
		return
	}

	authResult.Success = true
	authResult.Result = result
	c.JSON(200, authResult)
}
