package hardware

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// 获取内存信息
func GetMem() {
	vir, _ := mem.VirtualMemory()
	// vir.Total       // 总内存大小
	// vir.Available   // 闲置可用内存
	// vir.Used        // 已使用内存
	// vir.UsedPercent // 已使用百分比
	fmt.Println(vir)
}

// 获取Cpu基础信息
func GetCpu() {
	res, err := cpu.Times(false) // false 是展示全部总和 true 是分布展示
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// 采集cpu相关信息
func GetCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

// 获取Cpu负载信息
func GetCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}

// 获取host信息
func GetHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// 获取磁盘disk信息
func GetDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

// 获取NET IO信息
func GetNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}
