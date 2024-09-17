package http

import (
	"ldapctl/http/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRouters(r *gin.Engine) {
	ldapGroup := r.Group("/api/v1/ldap")
	{
		ldapGroup.GET("/", controllers.Hello)
		ldapGroup.GET("/health", controllers.Health)
		ldapGroup.POST("/auth/single", controllers.AuthSingle)
		ldapGroup.POST("/auth/multi", controllers.AuthMulti)
		ldapGroup.GET("/search/filter/:filter", controllers.SearchFilter)
		ldapGroup.GET("/search/user/:username", controllers.SearchUser)
		ldapGroup.POST("/search/multi", controllers.SearchMultiUser)
	}
}
