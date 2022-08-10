package router

import (
	"net/http"
	"strings"

	"github.com/Ungine-Tech/redirector/bootstrap"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		urlObj := c.Request.URL
		path := strings.TrimLeft(urlObj.RequestURI(), "/")
		location := strings.Replace(bootstrap.Target, "{{path}}", path, 1)
		c.Redirect(http.StatusMovedPermanently, location)
	})

	return r
}
