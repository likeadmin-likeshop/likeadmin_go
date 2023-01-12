package util

import (
	"likeadmin/core"
	"net"
)

var IpUtil = ipUtil{}

//serverUtil IP工具类
type ipUtil struct{}

//GetHostIp 获取本地主机名
func (su ipUtil) GetHostIp() (ip string) {
	conn, err := net.Dial("udp", "114.114.114.114:80")
	if err != nil {
		core.Logger.Errorf("GetHostIp Dial err: err=[%+v]", err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
