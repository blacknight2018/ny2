package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"ny2/dorm"
	"ny2/school"
	"ny2/user"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
		c.Done()
	}
}

func main() {
	g := gin.New()
	g.Use(TlsHandler())
	dorm.Register(g)
	school.Register(g)
	user.Register(g)
	//g.Run(":80")
	g.RunTLS(":443", "static/fullchain.crt", "static/private.pem")
	return
}
