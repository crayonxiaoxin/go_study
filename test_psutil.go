package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	// cpuInfo()
	memoryInfo()
	// diskInfo()
}

func diskInfo() {
	ps, err := disk.Partitions(false)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("------------- Disk -------------")
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{
		"Path",
		"Fstype",
		"Total",
		"Free",
		"Used",
		"UsedPercent(%)",
	}
	table.SetHeader(header)

	for _, partitionStat := range ps {
		us, err := disk.Usage(partitionStat.Mountpoint)
		if err != nil {
			continue
		}
		data := []string{
			us.Path,
			us.Fstype,
			sizeUnit(us.Total),
			sizeUnit(us.Free),
			sizeUnit(us.Used),
			strconv.FormatFloat(us.UsedPercent, 'f', 4, 64),
		}
		table.Append(data)
	}
	table.Render()
	fmt.Println("------------- Disk -------------")
}

func cpuInfo() {
	infoStats, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("------------- CPU -------------")
	count := len(infoStats)
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{}
	if count > 1 {
		header = []string{
			"Index",
			"Cores",
			"ModelName",
			"Mhz",
		}
	} else {
		header = []string{
			"Cores",
			"ModelName",
			"Mhz",
		}
	}
	table.SetHeader(header)
	for i := 0; i < count; i++ {
		cpu := infoStats[i]
		data := []string{}
		if count > 1 {
			data = []string{
				strconv.FormatInt(int64(cpu.CPU), 10),
				strconv.FormatInt(int64(cpu.Cores), 10),
				cpu.ModelName,
				strconv.FormatFloat(cpu.Mhz, 'f', 4, 64),
			}
		} else {
			data = []string{
				strconv.FormatInt(int64(cpu.Cores), 10),
				cpu.ModelName,
				strconv.FormatFloat(cpu.Mhz, 'f', 4, 64),
			}
		}
		table.Append(data)
	}
	table.Render()
	fmt.Println("------------- CPU -------------")
}

func memoryInfo() {
	vms, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("------------- Memory -------------")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgRedColor, tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgGreenColor, tablewriter.Bold},
	)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.FgGreenColor},
	)
	table.Append([]string{"Total", sizeUnit(vms.Total)})
	table.Append([]string{"Free", sizeUnit(vms.Free)})
	table.Append([]string{"Available", sizeUnit(vms.Available)})
	table.Append([]string{"Used", sizeUnit(vms.Used)})
	table.Append([]string{"UsedPercent", strconv.FormatFloat(vms.UsedPercent, 'f', 4, 64) + " %"})
	table.Append([]string{"Active", sizeUnit(vms.Active)})
	table.Append([]string{"Inactive", sizeUnit(vms.Inactive)})
	table.Append([]string{"Wired", sizeUnit(vms.Wired)})
	table.Append([]string{"Buffers", sizeUnit(vms.Buffers)})
	table.Append([]string{"Cached", sizeUnit(vms.Cached)})
	table.Append([]string{"Shared", sizeUnit(vms.Shared)})
	table.Append([]string{"SwapTotal", sizeUnit(vms.SwapTotal)})
	table.Append([]string{"SwapCached", sizeUnit(vms.SwapCached)})
	table.Append([]string{"SwapFree", sizeUnit(vms.SwapFree)})
	table.Render()
	fmt.Println("------------- Memory -------------")
}

func sizeUnit(size uint64) string {
	bytes := float64(size)
	kb := float64(1024)
	mb := 1024 * kb
	gb := 1024 * mb
	if bytes >= gb {
		return strconv.FormatFloat(bytes/gb, 'f', 2, 64) + " GB"
	} else if bytes >= mb {
		return strconv.FormatFloat(bytes/mb, 'f', 2, 64) + " MB"
	} else if bytes >= kb {
		return strconv.FormatFloat(bytes/kb, 'f', 2, 64) + " KB"
	} else if bytes > 0 {
		return strconv.FormatFloat(bytes, 'f', 2, 64) + " Bytes"
	} else {
		return "0"
	}
}
