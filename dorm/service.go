package dorm

import (
	"encoding/json"
	"ny2/bs"
	"ny2/bs/entity"
	"ny2/utils"
)

func getDormAllOrder(dormId int64) (bool, string) {
	var d bs.Dorm
	d.Id = &dormId
	ok, data := d.SelectAllOrder()
	if !ok {
		return false, utils.EmptyString
	}

	type t struct {
		entity.Order
		DormName string `json:"dorm_name"`
	}
	var tdata []t
	for _, v := range data {
		var tmp t
		tmp.AvatarUrl = v.AvatarUrl
		tmp.Type = v.Type
		tmp.TemplateId = v.TemplateId
		tmp.Comment = v.Comment
		tmp.Price = v.Price
		tmp.FinishTime = v.FinishTime
		tmp.SchoolId = v.SchoolId
		tmp.StuId = v.StuId
		tmp.RecvStu = v.RecvStu
		tmp.DormId = v.DormId
		var u bs.Stu
		u.StuId = &v.StuId
		u.SelectByStuId()
		tmp.DormName = u.DormRoom
		tdata = append(tdata, tmp)
	}
	if bytes, err := json.Marshal(tdata); err == nil {
		return true, string(bytes)
	}
	return false, utils.EmptyString

}
