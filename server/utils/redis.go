package utils

import (
	"context"
	"likeadmin/config"
	"likeadmin/core"
	"time"
)

var RedisUtil = redisUtil{}

//redisUtil Redis操作工具类
type redisUtil struct{}

//Set 设置键值对
func (ru redisUtil) Set(key string, value interface{}, timeSec int) bool {
	err := core.Redis.Set(context.Background(),
		config.Config.RedisPrefix+key, value, time.Duration(timeSec)*time.Second).Err()
	if err != nil {
		core.Logger.Infof("redisUtil.Set err: err=[%+v]", err)
		return false
	}
	return true
}

//Get 获取key的值
func (ru redisUtil) Get(key string) string {
	res, err := core.Redis.Get(context.Background(), config.Config.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Infof("redisUtil.Get err: err=[%+v]", err)
		return ""
	}
	return res
}

//SSet 将数据放入set缓存
func (ru redisUtil) SSet(key string, values ...interface{}) bool {
	err := core.Redis.SAdd(context.Background(), config.Config.RedisPrefix+key, values...).Err()
	if err != nil {
		core.Logger.Infof("redisUtil.SSet err: err=[%+v]", err)
		return false
	}
	return true
}

//SGet 根据key获取Set中的所有值
func (ru redisUtil) SGet(key string) []string {
	res, err := core.Redis.SMembers(context.Background(), config.Config.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Infof("redisUtil.SGet err: err=[%+v]", err)
		return []string{}
	}
	return res
}

//HMSet 设置key, 通过字典的方式设置多个field, value对
func (ru redisUtil) HMSet(key string, mapping map[string]string, timeSec int) bool {
	err := core.Redis.HSet(context.Background(), config.Config.RedisPrefix+key, mapping).Err()
	if err != nil {
		core.Logger.Infof("redisUtil.HMSet err: err=[%+v]", err)
		return false
	}
	if timeSec > 0 {
		if !ru.Expire(key, timeSec) {
			return false
		}
	}
	return true
}

//HSet 向hash表中放入数据,如果不存在将创建
func (ru redisUtil) HSet(key string, field string, value string, timeSec int) bool {
	return ru.HMSet(key, map[string]string{field: value}, timeSec)
}

//HGet 获取key中field域的值
func (ru redisUtil) HGet(key string, field string) string {
	res, err := core.Redis.HGet(context.Background(), config.Config.RedisPrefix+key, field).Result()
	if err != nil {
		core.Logger.Infof("redisUtil.HGet err: err=[%+v]", err)
		return ""
	}
	return res
}

//HExists 判断key中有没有field域名
func (ru redisUtil) HExists(key string, field string) bool {
	res, err := core.Redis.HExists(context.Background(), config.Config.RedisPrefix+key, field).Result()
	if err != nil {
		core.Logger.Infof("redisUtil.HGet err: err=[%+v]", err)
		return false
	}
	return res
}

//HDel 删除hash表中的值
func (ru redisUtil) HDel(key string, fields ...string) bool {
	err := core.Redis.HDel(context.Background(), config.Config.RedisPrefix+key, fields...).Err()
	if err != nil {
		core.Logger.Infof("redisUtil.HGet err: err=[%+v]", err)
		return false
	}
	return true
}

//Exists 判断多项key是否存在
func (ru redisUtil) Exists(keys ...string) int64 {
	fullKeys := ru.toFullKeys(keys)
	cnt, err := core.Redis.Exists(context.Background(), fullKeys...).Result()
	if err != nil {
		core.Logger.Infof("redisUtil.Expire err: err=[%+v]", err)
		return 0
	}
	return cnt
}

//Expire 指定缓存失效时间
func (ru redisUtil) Expire(key string, timeSec int) bool {
	err := core.Redis.Expire(context.Background(), config.Config.RedisPrefix+key, time.Duration(timeSec)*time.Second).Err()
	if err != nil {
		core.Logger.Infof("redisUtil.Expire err: err=[%+v]", err)
		return false
	}
	return true
}

//TTL 根据key获取过期时间
func (ru redisUtil) TTL(key string) int {
	td, err := core.Redis.TTL(context.Background(), config.Config.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Infof("redisUtil.TTL err: err=[%+v]", err)
		return 0
	}
	return int(td / time.Second)
}

//Del 删除一个或多个键
func (ru redisUtil) Del(keys ...string) bool {
	fullKeys := ru.toFullKeys(keys)
	err := core.Redis.Del(context.Background(), fullKeys...).Err()
	if err != nil {
		core.Logger.Infof("redisUtil.Del err: err=[%+v]", err)
		return false
	}
	return true
}

//toFullKeys 为keys批量增加前缀
func (ru redisUtil) toFullKeys(keys []string) (fullKeys []string) {
	for _, k := range keys {
		fullKeys = append(fullKeys, config.Config.RedisPrefix+k)
	}
	return
}
