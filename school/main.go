package school

import (
	"github.com/gin-gonic/gin"
	"ny2/gerr"
	"strconv"
)

func Register(engine *gin.Engine) {
	s := engine.Group("school")
	{
		s.GET("list", func(context *gin.Context) {
			if ok, data := getAllSchool(); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
				return
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)
		})
		s.GET("dorm", func(context *gin.Context) {
			schoolId := context.Query("school_id")
			schoolIdInt, err := strconv.Atoi(schoolId)
			if err != nil {
				gerr.SetResponse(context, gerr.UnKnow, nil)
				return
			}
			if ok, data := getSchoolDorm(int64(schoolIdInt)); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
				return
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)

		})
	}
}
