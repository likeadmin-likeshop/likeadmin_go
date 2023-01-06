package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"likeadmin/config"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	ToolsUtil    = toolsUtil{}
	allRandomStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

//toolsUtil 常用工具集合类
type toolsUtil struct{}

//RandomString 返回随机字符串
func (tu toolsUtil) RandomString(length int) string {
	byteList := make([]byte, length)
	for i := 0; i < length; i++ {
		byteList[i] = allRandomStr[rand.Intn(62)]
	}
	return string(byteList)
}

//MakeUuid 制作UUID
func (tu toolsUtil) MakeUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

//MakeMd5 制作MD5
func (tu toolsUtil) MakeMd5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

//MakeToken 生成唯一Token
func (tu toolsUtil) MakeToken() string {
	ms := time.Now().UnixMilli()
	token := tu.MakeMd5(tu.MakeUuid() + strconv.FormatInt(ms, 10) + tu.RandomString(8))
	tokenSecret := token + config.Config.Secret
	return tu.MakeMd5(tokenSecret) + tu.RandomString(6)
}

//Contains 判断src是否包含elem元素
func (tu toolsUtil) Contains(src interface{}, elem interface{}) bool {
	srcArr := reflect.ValueOf(src)
	if srcArr.Kind() == reflect.Slice {
		for i := 0; i < srcArr.Len(); i++ {
			if srcArr.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}
