package school

import (
	"encoding/json"
	"ny2/bs"
	"ny2/utils"
)

func getAllSchool() (bool, string) {
	r, data := bs.SelectAllSchool()
	if !r {
		return false, utils.EmptyString
	}
	if bytes, err := json.Marshal(data); err == nil {
		return true, string(bytes)
	}
	return false, utils.EmptyString
}

func getSchoolDorm(schoolId int64) (bool, string) {
	var s bs.School
	s.Id = schoolId
	ok, dorm := s.SelectSchoolAllDorm()
	if !ok {
		return false, utils.EmptyString
	}
	type t struct {
		bs.Dorm
		OrderSize int64 `json:"order_size"`
	}
	var tdata []t
	for _, v := range dorm {
		var tmp t
		tmp.Dorm = v
		_, orderSize := v.SelectOrderSize()
		tmp.OrderSize = orderSize
		tdata = append(tdata, tmp)
	}
	if bytes, err := json.Marshal(tdata); err == nil {
		return true, string(bytes)
	}
	return false, utils.EmptyString
}
