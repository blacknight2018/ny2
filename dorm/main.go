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
		limit := context.Query("limit")
		offset := context.Query("offset")

		//显示未过期的 未取消的 未接送的
		flag := context.Query("flag")
		var valid bool
		if flag == "valid" {
			valid = true
		}

		limitInt, err2 := strconv.Atoi(limit)
		offsetInt, err3 := strconv.Atoi(offset)
		if err2 != nil || err3 != nil {
			limitInt = 5
			offsetInt = 0
		}

		if err == nil {
			if ok, data := getDormOrder(int64(dormIdInt), int64(limitInt), int64(offsetInt), valid); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
				return
			}
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
	g.GET("order/size", func(context *gin.Context) {
		dormId := context.Query("dorm_id")
		dormIdInt, err := strconv.Atoi(dormId)
		if err == nil {
			if ok, data := getDormOrderSize(int64(dormIdInt)); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
				return
			}
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
}
