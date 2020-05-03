package main

import (
	"fmt"
	"github.com/free-bots/dirtist/cleaner"
	"github.com/free-bots/dirtist/info"
)

func main() {

	//fmt.Println(cpu.Info())

	info, _ := info.GetRamUsage()
	fmt.Println(info.String())

	return
	fmt.Println(cleaner.ClearJournal())
	fmt.Println(cleaner.GetCacheSize())
	fmt.Println(cleaner.GetJournalSize())
}
