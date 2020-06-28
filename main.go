package main

import (
	"github.com/gin-gonic/gin"
	"ny2/dorm"
	"ny2/school"
	"ny2/user"
)

func main() {
	g := gin.Default()
	dorm.Register(g)
	school.Register(g)
	user.Register(g)
	g.Run(":80")
	return
}
