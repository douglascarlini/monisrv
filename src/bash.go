package main

import (
	"fmt"
	"os/exec"
)

type Bash struct {
	HDD_USAGE string
	MEM_USAGE string
	CPU_USAGE string
	NET_USAGE string
}

func NewBash() *Bash {
	return &Bash{
		HDD_USAGE: "df -h | grep %s | awk '{print $5}'",
		MEM_USAGE: "free | grep Mem | awk '{print $3/$2*100}'",
		CPU_USAGE: "top -bn1 | grep \"Cpu(s)\" | awk '{print 100 - $8}'",
		NET_USAGE: "iftop -t -s 2 -n -P | grep \"Total send and receive\" | awk '{print $8}'",
	}
}

func (sh *Bash) Run(cmd string) string {
	res := exec.Command("bash", "-c", cmd)
	out, e := res.Output()

	if e != nil {
		panic(e)
	}

	return string(out)
}

func (sh *Bash) DiskUsage(disk string) string {
	return sh.Run(fmt.Sprintf(sh.HDD_USAGE, disk))
}

func (sh *Bash) MemUsage() string {
	return sh.Run(sh.MEM_USAGE)
}

func (sh *Bash) CpuUsage() string {
	return sh.Run(sh.CPU_USAGE)
}

func (sh *Bash) NetUsage() string {
	return sh.Run(sh.NET_USAGE)
}
