package dorm

import (
	"encoding/json"
	"ny2/bs"
	"ny2/bs/entity"
	"ny2/gerr"
	"ny2/utils"
)

// 获取宿舍楼里的有效的订单总数量
func getDormValidOrderSize(stuId int64, dormId int64) (int, string) {

	var d bs.Dorm
	d.Id = dormId
	var ok bool
	var data int64

	if stuId != 0 {
		//排除掉被拉黑的订单
		ok, data = d.SelectValidOrderWithBlockSize(stuId)
	} else {
		ok, data = d.SelectValidOrderSize()
	}
	if !ok {
		return gerr.UnKnow, utils.EmptyString
	}
	if bytes, err := json.Marshal(data); err == nil {
		return gerr.Ok, string(bytes)
	}
	return gerr.UnKnow, utils.EmptyString
}

func getDormOrder(stuId int64, dormId int64, limit int64, offset int64, selectValid bool) (int, string) {
	var d bs.Dorm
	d.Id = dormId
	var ok bool
	var data []entity.Order

	//是否返回未被接送的、未过期的有效的订单
	if selectValid {

		if stuId != 0 {
			ok, data = d.SelectValidOrderWithBlock(stuId, limit, offset)
		} else {
			ok, data = d.SelectValidOrder(limit, offset)
		}
	} else {

		ok, data = d.SelectAllOrder(limit, offset)
	}
	if !ok {
		return gerr.UnKnow, utils.EmptyString
	}

	type t struct {
		entity.Order
		DormName string `json:"dorm_name"`
		Place    string `json:"place"`
	}
	var tdata []t
	for _, v := range data {
		var tmp t
		tmp.Order = v

		var u bs.Stu
		u.StuId = v.StuId
		u.SelectByStuId()
		tmp.DormName = u.DormRoom

		tmp.Place = v.SelectPlaceName()

		////去除黑名单里的订单
		//if u.IsBeDstBlock(stuId){
		//	continue
		//}

		tdata = append(tdata, tmp)
	}
	if bytes, err := json.Marshal(tdata); err == nil {
		return gerr.Ok, string(bytes)
	}
	return gerr.UnKnow, utils.EmptyString

}
