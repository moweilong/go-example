package controllers

import (
	"ldapctl/g"
	"ldapctl/models"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	var msgResult MsgResult

	success, err := models.HealthCheck(g.Config().Ldap)
	if err != nil {
		msgResult.Msg = err.Error()

	} else {
		msgResult.Msg = "ok"
	}
	msgResult.Success = success
	c.JSON(200, msgResult)
}
