package controllers

import (
	"ldapctl/g"
	"ldapctl/models"

	"github.com/gin-gonic/gin"
)

type SearchUserResult struct {
	User    models.LADP_RESULT `json:"user"`
	Success bool               `json:"success"`
}

func SearchUser(c *gin.Context) {
	username := c.Param(":username")
	user, err := models.SingleSearchUser(g.Config().Ldap, username)
	if err != nil {
		var failedResult MsgResult
		failedResult.Msg = err.Error()
		c.JSON(500, failedResult)
	} else {
		var successResult SearchUserResult
		successResult.User = user
		successResult.Success = true
		c.JSON(200, successResult)
	}
}
