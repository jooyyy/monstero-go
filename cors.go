package monstera

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

// CorsMiddleware is a middleware for gin to fix cors problems
func CorsMiddleware(whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		// filterHost 做过滤器，防止不合法的域名访问
		var isAccess = false
		for _, v := range whitelist {
			match, _ := regexp.MatchString(v, origin)
			if match {
				isAccess = true
			}
		}
		if isAccess {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}