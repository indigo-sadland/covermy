// Package normalizer provides functions that removes artefacts from tools output
package normalizer

import (
	"fmt"
	"regexp"
	"strings"
)

const ffufhead = "\r"
const ffufmidle = "[2K"
const ffuftail = "[0m"
const ffuflongspace = "                    "
const ffufesc = "\u001B"

const gobusteresc = "\u001B"
const gobustermidle = `^\r\[2K`

func Start(tool string, rawlines []string) []string {
	var normilized []string

	switch tool {
	case "ffuf":
		normilized = ffuf(rawlines)
	case "gobuster":
		normilized = gobuster(rawlines)
	case "rustscan":
		normilized = rustscan(rawlines)
	case "nuclei":
		normilized = nuclei(rawlines)
	default:
		fmt.Printf("No normalization rules for the %s found.\n "+
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

func gobuster(rawlines []string) []string {
	var lines []string

	for _, l := range rawlines {
		l = strings.ReplaceAll(l, gobusteresc, "")
		l = regexp.MustCompile(gobustermidle).ReplaceAllString(l, "")
		lines = append(lines, l)
	}

	return lines
}

func rustscan(rawlines []string) []string {
	var lines []string

	for _, l := range rawlines {
		l = strings.ReplaceAll(l, "\u001B", "")
		l = regexp.MustCompile(`\[38;2;\d{1,3};\d{1,3};\d{1,3}m`).ReplaceAllString(l, "")
		l = strings.ReplaceAll(l, "[1;34m", "")
		l = strings.ReplaceAll(l, "[1;38;2;0;255;9m", "")

		x := regexp.MustCompile(`(^\[.])(\[0m)`)
		l = x.ReplaceAllString(l, "$1")

		l = regexp.MustCompile(`^\[1;?.*?m`).ReplaceAllString(l, "")
		l = regexp.MustCompile(`(\[0m)*$`).ReplaceAllString(l, "")
		lines = append(lines, l)

	}

	return lines
}

func nuclei(rawlines []string) []string {
	var lines []string

	for _, l := range rawlines {
		l = strings.ReplaceAll(l, "\u001B", "")
		l = regexp.MustCompile(`\[1;?.*?m`).ReplaceAllString(l, "")
		l = regexp.MustCompile(`\[\d{1,3}m`).ReplaceAllString(l, "")
		lines = append(lines, l)
	}

	return lines
}
