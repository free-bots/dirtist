package utils

import (
	"log"
	"os/exec"
)

// check for eternal programs used by this one

func ToolExists(name string) (string, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		log.Fatalf("install %s to conatinue \n", name)
		return "", err
	}
	return path, nil
}
