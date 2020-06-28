package utils

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const EmptyString = ""

func SendGet(url string, param map[string]string) (bool, string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, EmptyString
	}
	q := req.URL.Query()
	for v := range param {
		q.Add(v, param[v])
	}
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return false, EmptyString
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, EmptyString
	}
	return true, string(data)
}
func GetRawData(ctx *gin.Context) (bool, string) {
	bytes, err := ctx.GetRawData()
	if err != nil {
		return false, EmptyString
	}
	return true, string(bytes)
}
