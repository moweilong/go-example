package controllers

import (
	"ldapctl/g"
	"ldapctl/models"

	"github.com/gin-gonic/gin"
)

type SearchMultiResult struct {
	Success bool                         `json:"success"`
	Result  models.MultiSearchUserResult `json:"result"`
}

func SearchMultiUser(c *gin.Context) {
	var msgResult MsgResult
	var searchMultiResult SearchMultiResult
	var userList []string
	err := c.BindJSON(&userList)
	if err != nil {
		msgResult.Msg = err.Error()
		c.JSON(400, msgResult)
	}

	result, err := models.MultiSearchUser(g.Config().Ldap, userList)
	if err != nil {
		msgResult.Msg = err.Error()
		c.JSON(400, msgResult)
	} else {
		searchMultiResult.Success = true
		searchMultiResult.Result = result
		c.JSON(200, searchMultiResult)
	}
}
