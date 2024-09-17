package http

import (
	"ldapctl/g"

	"github.com/gin-gonic/gin"
)

func Start() {
	gSrv := gin.Default()
	ConfigRouters(gSrv)
	gSrv.Run(g.Config().Http.Listen)
}
