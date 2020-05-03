package cleaner

import (
	"github.com/free-bots/dirtist/utils"
	"os/exec"
	"regexp"
	"strconv"
	"unicode/utf8"
)

const journalCtl = "journalctl"

func ClearJournal() {
	// todo run as root
	// todo vacuum journalctl
}

func GetJournalSize() (float64, error) {
	programPath, err := utils.ToolExists(journalCtl)
	if err != nil {
		return -1, err
	}
	command := exec.Command(programPath, "--disk-usage")
	output, err := command.Output()
	if err != nil {
		return -1, err
	}

	// get something like 32.0M
	rawExpr := regexp.MustCompile("\\d+.\\d+\\w+")
	found := rawExpr.Find(output)
	foundStr := string(found)

	var foundUnit byte
	length := utf8.RuneCountInString(foundStr)
	if length > 0 {
		foundUnit = foundStr[length-1]
	}

	sizeExpr := regexp.MustCompile("\\d+.\\d+")
	sizeBytes := sizeExpr.Find(found)
	size, err := strconv.ParseFloat(string(sizeBytes), 64)
	if err != nil {
		return -1, err
	}

	foundUnitStr := string(foundUnit)
	switch foundUnitStr {
	case "B":
		size *= 1
	case "M":
		size *= 1024
	case "G":
		size *= 1024 * 1024
	}

	return size, nil
}

// todo add other logs
