package hardware

import "testing"

// 内存获取测试
func TestGetMem(t *testing.T) {
	GetMem()
	t.Log("successful!")
}

func TestGetCpu(t *testing.T) {
	GetCpu()
	t.Log("successful!")
}

func TestGetCpuInfo(t *testing.T) {
	GetCpuInfo()
	t.Log("successful!")
}

func TestGetCpuLoad(t *testing.T) {
	GetCpuLoad()
	t.Log("successful!")
}

func TestGetHostInfo(t *testing.T) {
	GetHostInfo()
	t.Log("successful!")
}

func TestDiskInfo(t *testing.T) {
	GetDiskInfo()
	t.Log("successful!")
}

func TestNetInfo(t *testing.T) {
	GetMem()
	t.Log("successful!")
}
