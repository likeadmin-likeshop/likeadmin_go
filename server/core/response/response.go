package response

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"likeadmin/core"
	"net/http"
)

//RespType 响应类型
type RespType struct {
	code int
	msg  string
}

//Response 响应格式结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Success = RespType{200, "成功"}
	Failed  = RespType{300, "失败"}

	ParamsValidError    = RespType{310, "参数校验错误"}
	ParamsTypeError     = RespType{311, "参数类型错误"}
	RequestMethodError  = RespType{312, "请求方法错误"}
	AssertArgumentError = RespType{313, "断言参数错误"}

	LoginAccountError = RespType{330, "登录账号或密码错误"}
	LoginDisableError = RespType{331, "登录账号已被禁用了"}
	TokenEmpty        = RespType{332, "token参数为空"}
	TokenInvalid      = RespType{333, "token参数无效"}

	NoPermission    = RespType{403, "无相关权限"}
	Request404Error = RespType{404, "请求接口不存在"}
	Request405Error = RespType{405, "请求方法不允许"}

	SystemError = RespType{500, "系统错误"}
)

//Make 以响应类型生成信息
func (rt RespType) Make(msg string) RespType {
	rt.msg = msg
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

//Result 统一响应
func Result(c *gin.Context, resp RespType, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: resp.code,
		Msg:  resp.msg,
		Data: data,
	})
}

//Copy 拷贝结构体
func Copy(c *gin.Context, toValue interface{}, fromValue interface{}) interface{} {
	if err := copier.Copy(toValue, fromValue); err != nil {
		Fail(c, SystemError)
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
