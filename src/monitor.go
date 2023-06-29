package main

import "fmt"

type Monitor struct {
	NUL_CHAR string
	BAR_CHAR string
}

func NewMonitor() *Monitor {
	return &Monitor{
		NUL_CHAR: "\u2591",
		BAR_CHAR: "\u2588",
	}
}

func (mon *Monitor) Disk() {
	sh := NewBash()

	mem := sh.MemUsage()
	fmt.Printf("MEM: %s\n", mem)

	cpu := sh.CpuUsage()
	fmt.Printf("CPU: %s\n", cpu)

	net := sh.NetUsage()
	fmt.Printf("NET: %s\n", net)

	disk := sh.DiskUsage("disk3s5")
	fmt.Printf("DISK: %s\n", disk)
}
