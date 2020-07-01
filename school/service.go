package school

import (
	"encoding/json"
	"ny2/bs"
	"ny2/gerr"
	"ny2/utils"
)

func getAllSchool() (int, string) {
	r, data := bs.SelectAllSchool()
	if !r {
		return gerr.UnKnow, utils.EmptyString
	}
	if bytes, err := json.Marshal(data); err == nil {
		return gerr.Ok, string(bytes)
	}
	return gerr.UnKnow, utils.EmptyString
}

func getSchoolDorm(schoolId int64) (int, string) {
	var s bs.School
	s.Id = schoolId
	ok, dorm := s.SelectSchoolAllDorm()
	if !ok {
		return gerr.UnKnow, utils.EmptyString
	}
	type t struct {
		bs.Dorm
		OrderSize int64 `json:"order_size"`
	}
	var tdata []t
	for _, v := range dorm {
		var tmp t
		tmp.Dorm = v
		_, orderSize := v.SelectValidOrderSize()
		tmp.OrderSize = orderSize
		tdata = append(tdata, tmp)
	}
	if bytes, err := json.Marshal(tdata); err == nil {
		return gerr.Ok, string(bytes)
	}
	return gerr.UnKnow, utils.EmptyString
}
