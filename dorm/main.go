package dorm

import (
	"fmt"
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
		stuId := context.Query("stu_id")
		//显示未过期的 未取消的 未接送的
		flag := context.Query("flag")
		var valid bool
		if flag == "valid" {
			valid = true
		}
		stuIdInt, err := strconv.Atoi(stuId)
		limitInt, err2 := strconv.Atoi(limit)
		offsetInt, err3 := strconv.Atoi(offset)
		if err2 != nil || err3 != nil {
			limitInt = 5
			offsetInt = 0
		}
		fmt.Println(stuId)
		if err == nil {
			errCode, data := getDormOrder(int64(stuIdInt), int64(dormIdInt), int64(limitInt), int64(offsetInt), valid)
			gerr.SetResponse(context, errCode, &data)
			return
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
	g.GET("order/size", func(context *gin.Context) {
		dormId := context.Query("dorm_id")
		stuId := context.Query("stu_id")
		stuIdInt, _ := strconv.Atoi(stuId)
		dormIdInt, err2 := strconv.Atoi(dormId)
		fmt.Println(stuId)
		if err2 == nil {
			errCode, data := getDormValidOrderSize(int64(stuIdInt), int64(dormIdInt))
			gerr.SetResponse(context, errCode, &data)
			return
		}
		gerr.SetResponse(context, gerr.UnKnow, nil)
	})
}
