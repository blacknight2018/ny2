package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"ny/stu"
	"ny2/bs/entity"
	"ny2/gerr"
	"ny2/utils"
	"ny2/wxapi"
	"strconv"
	"time"
)

func Register(engine *gin.Engine) {
	o := engine.Group("order")
	{
		o.POST("", func(context *gin.Context) {
			if ok, data := utils.GetRawData(context); ok {
				orderType := gjson.Get(data, "type").Int()

				if orderType < entity.Delivery || orderType > entity.Buy {
					gerr.SetResponse(context, gerr.ParamError, nil)
					return
				}

				stuId := gjson.Get(data, "stu_id").Int()

				if stu.GetStuExitsById(stuId) == false {
					gerr.SetResponse(context, gerr.UnKnowUser, nil)
					return
				}

				orderPrice := gjson.Get(data, "price").String()
				orderEndTime := gjson.Get(data, "end_time").Int()
				orderComment := gjson.Get(data, "comment").String()
				orderTemplateId := gjson.Get(data, "template_id").String()
				fmt.Println(orderType, orderPrice, orderEndTime, orderComment)

				if sendOrder(int(orderType), stuId, orderPrice, time.Unix(orderEndTime/1000, 0), orderComment, orderTemplateId) {
					gerr.SetResponse(context, gerr.Ok, nil)
					return
				}
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)
		})
	}
	u := engine.Group("user")
	{
		u.POST("code", func(context *gin.Context) {
			if ok, data := utils.GetRawData(context); ok {
				code := gjson.Get(data, "code").String()
				if ok, rust := wxapi.Code2Session(code); ok {
					gerr.SetResponse(context, gerr.Ok, &rust)
					return
				}
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)
		})
		u.POST("login", func(context *gin.Context) {
			if ok, data := utils.GetRawData(context); ok {
				openId := gjson.Get(data, "open_id").String()
				if isOpenIdExist(openId) {
					gerr.SetResponse(context, gerr.Ok, nil)
					return
				}
				nickName := gjson.Get(data, "nick_name").String()
				avatarUrl := gjson.Get(data, "avatar_url").String()

				if login(avatarUrl, nickName, openId) {
					gerr.SetResponse(context, gerr.Ok, nil)
					return
				}
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)
		})
		u.GET("info", func(context *gin.Context) {
			openId := context.Query("open_id")
			stuId := context.Query("stu_id")
			stuIdInt, err := strconv.Atoi(stuId)
			if err == nil {
				if ok, data := getSimpleInfoByStuId(int64(stuIdInt)); ok {
					gerr.SetResponse(context, gerr.Ok, &data)
					return
				}

			}
			if ok, data := getSimpleInfoByOpenId(openId); ok {
				gerr.SetResponse(context, gerr.Ok, &data)
				return
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)
		})
		u.PUT("", func(context *gin.Context) {
			if ok, data := utils.GetRawData(context); ok {
				openId := gjson.Get(data, "open_id").String()
				dormId := gjson.Get(data, "dorm_id").Int()
				stuNumber := gjson.Get(data, "stu_number").String()
				stuMobile := gjson.Get(data, "mobile").String()
				room := gjson.Get(data, "room").String()
				if putInfoByOpenId(openId, dormId, stuNumber, stuMobile, room) {
					gerr.SetResponse(context, gerr.Ok, &data)
					return
				}
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)
		})
		u.POST("msg", func(context *gin.Context) {
			if ok, data := utils.GetRawData(context); ok {
				senderStuId := gjson.Get(data, "sender_stuid").Int()
				recipientStuId := gjson.Get(data, "recipient_stuid").Int()
				content := gjson.Get(data, "content").String()
				if len(content) == 0 {
					gerr.SetResponse(context, gerr.SendNotEmpty, nil)
					return
				}
				if ok := sendTxtMessage(senderStuId, recipientStuId, content); ok {
					gerr.SetResponse(context, gerr.Ok, nil)
					return
				}
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)
		})
		u.GET("msg", func(context *gin.Context) {
			stuIdA := context.Query("stua_id")
			stuIdB := context.Query("stub_id")
			stuIdAInt, err := strconv.Atoi(stuIdA)
			stuIdBInt, err2 := strconv.Atoi(stuIdB)
			if err == nil && err2 == nil {
				if ok, data := getStuMsg(int64(stuIdAInt), int64(stuIdBInt), 10); ok {
					gerr.SetResponse(context, gerr.Ok, &data)
				}
				return
			}
			gerr.SetResponse(context, gerr.UnKnow, nil)
		})
	}
}