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
			errCode, data := getAllSchool()
			gerr.SetResponse(context, errCode, &data)
		})
		s.GET("canteen/list", func(context *gin.Context) {
			schoolId := context.Query("school_id")
			schoolIdInt, err := strconv.Atoi(schoolId)
			if err != nil {
				gerr.SetResponse(context, gerr.UnKnow, nil)
				return
			}
			errCode, data := getSchoolCanteen(int64(schoolIdInt))
			gerr.SetResponse(context, errCode, &data)
		})
		s.GET("shop/list", func(context *gin.Context) {
			schoolId := context.Query("school_id")
			schoolIdInt, err := strconv.Atoi(schoolId)
			if err != nil {
				gerr.SetResponse(context, gerr.UnKnow, nil)
				return
			}
			errCode, data := getSchoolShop(int64(schoolIdInt))
			gerr.SetResponse(context, errCode, &data)
		})
		s.GET("dorm", func(context *gin.Context) {
			schoolId := context.Query("school_id")
			schoolIdInt, err := strconv.Atoi(schoolId)
			if err != nil {
				gerr.SetResponse(context, gerr.UnKnow, nil)
				return
			}
			errCode, data := getSchoolDorm(int64(schoolIdInt))
			gerr.SetResponse(context, errCode, &data)

		})
	}
}
