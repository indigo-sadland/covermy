// Package normalizer provides functions that removes artefacts from tools output
package normalizer

import (
	"covermy/utils/logcamp"
	"strings"
)

const ffufhead = "\r"
const ffufmidle = "[2K"
const ffuftail = "[0m"
const ffuflongspace = "                    "
const ffufesc = "\u001B"

func Start(tool string, rawlines []string) []string {
	var normilized []string

	switch tool {
	case "ffuf":
		normilized = ffuf(rawlines)
	default:
		logcamp.InfoLogger.Printf("No normalization rules for the %s found.\n "+
			"Feel free to leave a request!", tool)
		return rawlines
	}

	return normilized
}

func ffuf(rawlines []string) []string {
	var lines []string

	for _, l := range rawlines {
		l = strings.ReplaceAll(l, ffufhead, "")
		l = strings.ReplaceAll(l, ffufesc, "")
		l = strings.ReplaceAll(l, ffufmidle, "")
		l = strings.ReplaceAll(l, ffuftail, "")
		l = strings.ReplaceAll(l, ffuflongspace, " ")
		lines = append(lines, l)
	}

	return lines
}
