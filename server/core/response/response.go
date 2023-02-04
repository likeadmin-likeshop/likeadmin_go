package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"likeadmin/core"
	"net/http"
	"strconv"
)

//RespType 响应类型
type RespType struct {
	code int
	msg  string
	data interface{}
}

//Response 响应格式结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Success = RespType{code: 200, msg: "成功"}
	Failed  = RespType{code: 300, msg: "失败"}

	ParamsValidError    = RespType{code: 310, msg: "参数校验错误"}
	ParamsTypeError     = RespType{code: 311, msg: "参数类型错误"}
	RequestMethodError  = RespType{code: 312, msg: "请求方法错误"}
	AssertArgumentError = RespType{code: 313, msg: "断言参数错误"}

	LoginAccountError = RespType{code: 330, msg: "登录账号或密码错误"}
	LoginDisableError = RespType{code: 331, msg: "登录账号已被禁用了"}
	TokenEmpty        = RespType{code: 332, msg: "token参数为空"}
	TokenInvalid      = RespType{code: 333, msg: "token参数无效"}

	NoPermission    = RespType{code: 403, msg: "无相关权限"}
	Request404Error = RespType{code: 404, msg: "请求接口不存在"}
	Request405Error = RespType{code: 405, msg: "请求方法不允许"}

	SystemError = RespType{code: 500, msg: "系统错误"}
)

//Error 实现error方法
func (rt RespType) Error() string {
	return strconv.Itoa(rt.code) + ":" + rt.msg
}

//Make 以响应类型生成信息
func (rt RespType) Make(msg string) RespType {
	rt.msg = msg
	return rt
}

//MakeData 以响应类型生成数据
func (rt RespType) MakeData(data interface{}) RespType {
	rt.data = data
	return rt
}

//Code 获取code
func (rt RespType) Code() int {
	return rt.code
}

//Msg 获取msg
func (rt RespType) Msg() string {
	return rt.msg
}

//Data 获取data
func (rt RespType) Data() interface{} {
	return rt.data
}

//Result 统一响应
func Result(c *gin.Context, resp RespType, data interface{}) {
	if data == nil {
		data = resp.data
	}
	if resp != Success {
		c.Error(resp)
	}
	c.JSON(http.StatusOK, Response{
		Code: resp.code,
		Msg:  resp.msg,
		Data: data,
	})
}

//Copy 拷贝结构体
func Copy(toValue interface{}, fromValue interface{}) interface{} {
	if err := copier.Copy(toValue, fromValue); err != nil {
		core.Logger.Errorf("Copy err: err=[%+v]", err)
		panic(SystemError)
	}
	return toValue
}

//Ok 正常响应
func Ok(c *gin.Context) {
	Result(c, Success, []string{})
}

//OkWithMsg 正常响应附带msg
func OkWithMsg(c *gin.Context, msg string) {
	resp := Success
	resp.msg = msg
	Result(c, resp, []string{})
}

//OkWithData 正常响应附带data
func OkWithData(c *gin.Context, data interface{}) {
	Result(c, Success, data)
}

//respLogger 打印日志
func respLogger(resp RespType, template string, args ...interface{}) {
	loggerFunc := core.Logger.WithOptions(zap.AddCallerSkip(2)).Warnf
	if resp.code >= 500 {
		loggerFunc = core.Logger.WithOptions(zap.AddCallerSkip(1)).Errorf
	}
	loggerFunc(template, args...)
}

//Fail 错误响应
func Fail(c *gin.Context, resp RespType) {
	respLogger(resp, "Request Fail: url=[%s], resp=[%+v]", c.Request.URL.Path, resp)
	Result(c, resp, []string{})
}

//FailWithMsg 错误响应附带msg
func FailWithMsg(c *gin.Context, resp RespType, msg string) {
	resp.msg = msg
	respLogger(resp, "Request FailWithMsg: url=[%s], resp=[%+v]", c.Request.URL.Path, resp)
	Result(c, resp, []string{})
}

//FailWithData 错误响应附带data
func FailWithData(c *gin.Context, resp RespType, data interface{}) {
	respLogger(resp, "Request FailWithData: url=[%s], resp=[%+v], data=[%+v]", c.Request.URL.Path, resp, data)
	Result(c, resp, data)
}

//IsFailWithResp 判断是否出现错误，并追加错误返回信息
func IsFailWithResp(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	switch v := err.(type) {
	// 自定义类型
	case RespType:
		Fail(c, v)
	// 其他类型
	default:
		Fail(c, SystemError)
	}
	return true
}

//CheckAndResp 判断是否出现错误，并返回对应响应
func CheckAndResp(c *gin.Context, err error) {
	if IsFailWithResp(c, err) {
		return
	}
	Ok(c)
}

//CheckAndRespWithData 判断是否出现错误，并返回对应响应（带data数据）
func CheckAndRespWithData(c *gin.Context, data interface{}, err error) {
	if IsFailWithResp(c, err) {
		return
	}
	OkWithData(c, data)
}

//CheckErr 校验未知错误并抛出
func CheckErr(err error, template string, args ...interface{}) (e error) {
	prefix := ": "
	if len(args) > 0 {
		prefix = " ,"
	}
	args = append(args, err)
	if err != nil {
		core.Logger.WithOptions(zap.AddCallerSkip(1)).Errorf(template+prefix+"err=[%+v]", args...)
		return SystemError
	}
	return
}

//CheckErrDBNotRecord 校验数据库记录不存在的错误
func CheckErrDBNotRecord(err error, msg string) (e error) {
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		core.Logger.WithOptions(zap.AddCallerSkip(1)).Infof("CheckErrDBNotRecord err: err=[%+v]", err)
		return AssertArgumentError.Make(msg)
	}
	return
}
