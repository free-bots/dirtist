package cleaner

import (
	"fmt"
	"os"
)

// clear the .cache and other paths ... clear package cache from extra package -> packageManager

func GetCacheSize() int64 {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(homeDir)

	fi, err := os.Stat(homeDir + "/.cache")
	if err != nil {
		fmt.Println(err)
	}

	return fi.Size()
}
