package monitor

import (
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/NimbleMarkets/ntcharts/linechart/streamlinechart"
	"github.com/shirou/gopsutil/process"
)

func MonitorProcess(args string) {
	processNameRegexp := regexp.MustCompile(args)
	if processNameRegexp == nil {
		panic("can not compile args")
	}

	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}

	for _, p := range processes {
		processName, err := p.Name()
		if err != nil {
			continue
		}
		if processNameRegexp.MatchString(processName) {
			wg.Add(1)
			go MonitorCPU(wg, processName, p.Pid)
		}
	}

	wg.Wait()
}

func MonitorCPU(wg *sync.WaitGroup, name string, pid int32) {
	defer wg.Done()
	proc, err := process.NewProcess(pid)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Second)
	slc := streamlinechart.New(64, 10)

	// 循环监控 CPU 使用率
	for range ticker.C {
		cpuPercent, err := proc.Percent(time.Second * 1)
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		memoryPercent, err := proc.MemoryPercent()
		if err != nil {
			panic(err)
		}
		fmt.Printf("name %v CPU占用: %.2f%%, 内存占用: %.2f%%\n", name, cpuPercent, memoryPercent)
		slc.Push(cpuPercent)
		slc.Draw()
		fmt.Println(slc.View())
	}
}
