package cleaner

import (
	"fmt"
	"github.com/free-bots/dirtist/utils"
	"os/exec"
	"regexp"
	"unicode/utf8"
)

const journalCtl = "journalctl"

func ClearJournal() {
	// todo run as root
	// todo vacuum journalctl
}

func GetJournalSize() (int, error) {
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
	rawExpr := regexp.MustCompile("\\d+.\\d\\w+")
	found := rawExpr.Find(output)
	foundStr := fmt.Sprintf("%s", found)

	var foundUnit byte
	length := utf8.RuneCountInString(foundStr)
	if length > 0 {
		foundUnit = foundStr[length-1]
	}

	fmt.Printf("%s \n", foundStr)
	fmt.Printf("%c \n", foundUnit)

	foundUnitStr := fmt.Sprintf("%c", foundUnit)
	switch foundUnitStr {
	case "B":
	case "m":
	case "G":
	}
	// todo calculate the size

	return 0, nil
}

// todo add other logs
