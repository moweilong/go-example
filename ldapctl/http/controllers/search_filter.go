package controllers

import (
	"ldapctl/g"
	"ldapctl/models"

	"github.com/gin-gonic/gin"
)

func SearchFilter(c *gin.Context) {
	searchFilter := c.Param(":filter")
	results, err := models.SingleSearch(g.Config().Ldap, searchFilter)
	if err != nil {
		var failedResult MsgResult
		failedResult.Msg = err.Error()
		c.JSON(403, failedResult)
	} else {
		var successResult SearchResult
		successResult.Results = results
		successResult.Success = true
		c.JSON(200, successResult)
	}
}
