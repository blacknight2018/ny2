package dorm

import (
	"github.com/gin-gonic/gin"
	"ny2/gerr"
	"strconv"
)

func Register(engine *gin.Engine) {
	g := engine.Group("dorm")
	g.GET("order", func(context *gin.Context) {
		dormId := context.Query("dorm_id")
		dormIdInt, err := strconv.Atoi(dormId)
		if err == nil {
			if ok, data := getDormAllOrder(int64(dormIdInt)); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
				return
			}
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
}
