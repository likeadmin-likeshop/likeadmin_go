package util

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"likeadmin/core"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var ServerUtil = serverUtil{}

//serverUtil 服务器信息获取工具
type serverUtil struct{}

//GetFmtSize 按照正确的格式缩放字节
func (su serverUtil) GetFmtSize(data uint64) string {
	var factor float64 = 1024
	res := float64(data)
	for _, unit := range []string{"", "K", "M", "G", "T", "P"} {
		if res < factor {
			return fmt.Sprintf("%.2f%sB", res, unit)
		}
		res /= factor
	}
	return fmt.Sprintf("%.2f%sB", res, "P")
}

//GetFmtTime 格式化显示时间 (毫秒)
func (su serverUtil) GetFmtTime(ms int64) (res string) {
	rem := ms / 1000
	days, rem := rem/86400, rem%86400
	hours, rem := rem/3600, rem%3600
	minutes := rem / 60
	res = strconv.FormatInt(minutes, 10) + "分钟"
	if hours > 0 {
		res = strconv.FormatInt(hours, 10) + "小时" + res
	}
	if days > 0 {
		res = strconv.FormatInt(days, 10) + "天" + res
	}
	return res
}

//GetCpuInfo 获取CPU信息
func (su serverUtil) GetCpuInfo() (data map[string]interface{}) {
	cnt, err := cpu.Counts(true)
	if err != nil {
		core.Logger.Errorf("GetCpuInfo Counts err: err=[%+v]", err)
		return data
	}
	tss, err := cpu.Times(false)
	if err != nil {
		core.Logger.Errorf("GetCpuInfo Times err: err=[%+v]", err)
		return data
	}
	ts := tss[0]
	return map[string]interface{}{
		"cpu_num": cnt,
		"total":   ToolsUtil.Round(ts.Total(), 2),
		"sys":     ToolsUtil.Round(ts.System/ts.Total(), 2),
		"used":    ToolsUtil.Round(ts.User/ts.Total(), 2),
		"wait":    ToolsUtil.Round(ts.Iowait/ts.Total(), 2),
		"free":    ToolsUtil.Round(ts.Idle/ts.Total(), 2),
	}
}

//GetMemInfo 获取内存信息
func (su serverUtil) GetMemInfo() (data map[string]interface{}) {
	number := math.Pow(1024, 3)
	vm, err := mem.VirtualMemory()
	if err != nil {
		core.Logger.Errorf("GetMemInfo VirtualMemory err: err=[%+v]", err)
		return data
	}
	return map[string]interface{}{
		"total": ToolsUtil.Round(float64(vm.Total)/number, 2),
		"used":  ToolsUtil.Round(float64(vm.Used)/number, 2),
		"free":  ToolsUtil.Round(float64(vm.Available)/number, 2),
		"usage": ToolsUtil.Round(vm.UsedPercent, 2),
	}
}

//GetSysInfo 获取服务器信息
func (su serverUtil) GetSysInfo() (data map[string]interface{}) {
	infoStat, err := host.Info()
	if err != nil {
		core.Logger.Errorf("GetSysInfo Info err: err=[%+v]", err)
		return data
	}
	pwd, err := os.Getwd()
	if err != nil {
		core.Logger.Errorf("GetSysInfo Getwd err: err=[%+v]", err)
		return data
	}
	return map[string]interface{}{
		"computerName": infoStat.Hostname,
		"computerIp":   IpUtil.GetHostIp(),
		"userDir":      pwd,
		"osName":       infoStat.OS,
		"osArch":       infoStat.KernelArch,
	}
}

//GetDiskInfo 获取磁盘信息
func (su serverUtil) GetDiskInfo() (data []map[string]interface{}) {
	partStats, err := disk.Partitions(false)
	if err != nil {
		core.Logger.Errorf("GetDiskInfo Partitions err: err=[%+v]", err)
		return data
	}
	for i := 0; i < len(partStats); i++ {
		part := partStats[i]
		usage, uErr := disk.Usage(part.Mountpoint)
		if uErr != nil {
			core.Logger.Errorf("GetDiskInfo Usage err: err=[%+v]", err)
			continue
		}
		data = append(data, map[string]interface{}{
			"dirName":     part.Mountpoint,
			"sysTypeName": part.Fstype,
			"typeName":    part.Device,
			"total":       su.GetFmtSize(usage.Total),
			"free":        su.GetFmtSize(usage.Free),
			"used":        su.GetFmtSize(usage.Used),
			"usage":       ToolsUtil.Round(usage.UsedPercent, 2),
		})
	}
	return data
}

//GetGoInfo 获取Go环境及服务信息
func (su serverUtil) GetGoInfo() (data map[string]interface{}) {
	number := math.Pow(1024, 2)
	curProc, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		core.Logger.Errorf("GetGoInfo NewProcess err: err=[%+v]", err)
		return data
	}
	memInfo, err := curProc.MemoryInfo()
	if err != nil {
		core.Logger.Errorf("GetGoInfo MemoryInfo err: err=[%+v]", err)
		return data
	}
	startTime, err := curProc.CreateTime()
	if err != nil {
		core.Logger.Errorf("GetGoInfo CreateTime err: err=[%+v]", err)
		return data
	}
	return map[string]interface{}{
		"name":      "Go",
		"version":   runtime.Version(),
		"home":      os.Args[0],
		"inputArgs": fmt.Sprintf("[%s]", strings.Join(os.Args[1:], ", ")),
		"total":     ToolsUtil.Round(float64(memInfo.VMS)/number, 2),
		"max":       ToolsUtil.Round(float64(memInfo.VMS)/number, 2),
		"free":      ToolsUtil.Round((float64(memInfo.VMS-memInfo.RSS))/number, 2),
		"usage":     ToolsUtil.Round(float64(memInfo.RSS)/number, 2),
		"runTime":   su.GetFmtTime(time.Now().UnixMilli() - startTime),
		"startTime": time.UnixMilli(startTime).Format(core.TimeFormat),
	}
}
