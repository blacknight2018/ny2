package dorm

import (
	"encoding/json"
	"ny2/bs"
	"ny2/utils"
)

func getDormAllOrder(dormId int64) (bool, string) {
	var d bs.Dorm
	d.Id = &dormId
	ok, data := d.SelectAllOrder()
	if !ok {
		return false, utils.EmptyString
	}
	if bytes, err := json.Marshal(data); err == nil {
		return true, string(bytes)
	}
	return false, utils.EmptyString

}
