package wxapi

import "ny2/utils"

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
