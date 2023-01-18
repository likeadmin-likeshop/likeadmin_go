package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"likeadmin/core/response"
	"mime/multipart"
)

var VerifyUtil = verifyUtil{}

//verifyUtil 参数验证工具类
type verifyUtil struct{}

func (vu verifyUtil) VerifyJSON(c *gin.Context, obj any) {
	if err := c.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
}

func (vu verifyUtil) VerifyJSONArray(c *gin.Context, obj any) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
}

func (vu verifyUtil) VerifyBody(c *gin.Context, obj any) {
	if err := c.ShouldBind(obj); err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
}

func (vu verifyUtil) VerifyHeader(c *gin.Context, obj any) {
	if err := c.ShouldBindHeader(obj); err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
}

func (vu verifyUtil) VerifyQuery(c *gin.Context, obj any) {
	if err := c.ShouldBindQuery(obj); err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
}

func (vu verifyUtil) VerifyFile(c *gin.Context, name string) (file *multipart.FileHeader) {
	file, err := c.FormFile(name)
	if err != nil {
		panic(response.ParamsValidError.MakeData(err.Error()))
		return
	}
	return file
}
