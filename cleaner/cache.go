package cleaner

import (
	"fmt"
	"os"
)

// clear the .cache and other paths ... clear package cache from extra package -> packageManager

type CacheItem struct {
	IsDir    bool        `json:"is_dir"`
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	Children []CacheItem `json:"children"`
	Size     uint64      `json:"size"`
}

func GetCacheSize() int64 {
	// todo get from config or use default
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
