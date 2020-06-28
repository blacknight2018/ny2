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
	if bytes, err := json.Marshal(dorm); err == nil {
		return true, string(bytes)
	}
	return false, utils.EmptyString
}
