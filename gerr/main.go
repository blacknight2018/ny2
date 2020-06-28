package gerr

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Ok           = iota
	UnKnow       = iota
	ParamError   = iota
	UnKnowUser   = iota
	SendNotEmpty = iota
)

var Msg = map[int]string{
	Ok:           "成功",
	UnKnow:       "未知错误",
	ParamError:   "参数错误",
	UnKnowUser:   "未知用户",
	SendNotEmpty: "发送内容不能为空",
}

func SetResponse(ctx *gin.Context, ErrCode int, RespJson *string) {
	type tmp struct {
		ErrCode int         `json:"err_code"`
		ErrMsg  string      `json:"err_msg"`
		Data    interface{} `json:"data"`
	}
	var t tmp
	t.ErrCode = ErrCode
	t.ErrMsg = Msg[ErrCode]
	var obj interface{}
	if RespJson != nil {
		json.Unmarshal([]byte(*RespJson), &obj)
	}
	t.Data = obj
	ctx.JSON(http.StatusOK, t)

}
