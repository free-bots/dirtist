package info

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

// todo get ram cpu disk network usage and specs and also os infos like uptime version etc.

type RamUsage struct {
	TotalRam  uint64 `json:"total_ram"`
	UsedRam   uint64 `json:"used_ram"`
	TotalSwap uint64 `json:"total_swap"`
	UsedSwap  uint64 `json:"used_swap"`
}

func (usage RamUsage) String() string {
	return fmt.Sprintf("RAM: %d, USED: %d, SWAP: %d, USED: %d",
		usage.TotalRam, usage.UsedRam, usage.TotalSwap, usage.UsedSwap)
}

func GetRamUsage() (RamUsage, error) {
	ramUsage := RamUsage{}
	ram, err := mem.VirtualMemory()
	if err != nil {
		return RamUsage{}, err
	} else {
		ramUsage.TotalRam = ram.Total
		ramUsage.UsedRam = ram.Used
	}

	swap, err := mem.SwapMemory()
	if err != nil {
		return RamUsage{}, err
	} else {
		ramUsage.TotalSwap = swap.Total
		ramUsage.UsedSwap = swap.Used
	}

	return ramUsage, nil
}
