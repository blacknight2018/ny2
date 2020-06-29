package wxapi

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"ny2/utils"
	"strconv"
	"strings"
)

const AppId = "wx53e43951ac45b9a4"
const AppSecret = "3cd2d8c08ebf9f10ec4abeb34e89b828"

func Code2Session(code string) (bool, string) {
	var params = make(map[string]string)
	params["appid"] = AppId
	params["secret"] = AppSecret
	params["js_code"] = code
	params["grant_type"] = code
	return utils.SendGet("https://api.weixin.qq.com/sns/jscode2session", params)
}
func getAccessToken() (bool, string) {
	var param = make(map[string]string)
	param["grant_type"] = "client_credential"
	param["appid"] = AppId
	param["secret"] = AppSecret
	r, data := utils.SendGet("https://api.weixin.qq.com/cgi-bin/token", param)
	if r {
		return true, gjson.Get(data, "access_token").String()
	}
	return false, utils.EmptyString
}
func SendOrderNotify(openId string, templateId string, dormName string, orderId int64, comment string) bool {
	var param = make(map[string]interface{})
	param["thing2"] = dormName
	param["character_string3"] = strconv.Itoa(int(orderId))
	param["phrase6"] = "已被配送"
	param["thing7"] = comment

	return sendNotify(openId, templateId, param)
}
func sendNotify(openId string, templateId string, data map[string]interface{}) bool {
	//UserLogger.InfoLog("sendNotify():" + openId)
	ok, accessToken := getAccessToken()
	if !ok {
		return false
	}
	reqUrl := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + accessToken

	var param = make(map[string]interface{})
	param["access_token"] = accessToken
	param["touser"] = openId
	param["template_id"] = templateId //"mtm0AlM9wiWsYdE-ihi8lpiioqFi6EtQeusYaY7UrRk"

	//消息参数
	type dataValue struct {
		Value interface{} `json:"value"`
	}
	for k, v := range data {
		data[k] = dataValue{v}
	}
	param["data"] = data

	reqJsonBytes, _ := json.Marshal(param)
	reqJsonString := string(reqJsonBytes)

	req, err1 := http.NewRequest("POST", reqUrl, strings.NewReader(reqJsonString))
	resp, err2 := http.DefaultClient.Do(req)
	if err1 != nil || err2 != nil {
		return false
	}
	defer resp.Body.Close()
	respData, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return false
	}
	respDataString := string(respData)
	errCode := int(gjson.Get(respDataString, "errcode").Num)
	if errCode != 0 {
		return false
	}
	return true
}
