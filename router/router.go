package router

import (
	"net/http"

	"github.com/Ungine-Tech/redirector/bootstrap"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, bootstrap.Target)
	})

	return r
}
