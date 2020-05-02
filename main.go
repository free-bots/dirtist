package main

import (
	"fmt"
	"github.com/free-bots/dirtist/cleaner"
)

func main() {
	fmt.Println("hi")
	fmt.Println(cleaner.GetCacheSize())
	fmt.Println(cleaner.GetJournalSize())
}
