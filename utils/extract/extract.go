package extract

import (
	"net/url"
	"regexp"
	"strings"
)

func Target(CmdInput string) (string, string) {
	CmdInput = strings.TrimSpace(CmdInput)

	if regexp.MustCompile(`^https?`).MatchString(CmdInput) {
		read, _ := url.Parse(CmdInput)
		CmdInput = read.Host
	}

	if regexp.MustCompile(`^www\.`).MatchString(CmdInput) {
		CmdInput = regexp.MustCompile(`^www\.`).ReplaceAllString(CmdInput, "")
	}

	if regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}`).MatchString(CmdInput) {
		return regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}`).FindString(CmdInput), "ip"
	}

	if regexp.MustCompile(`([a-z0-9\-]+\.)+[a-z0-9\-]+`).FindString(CmdInput) != "" {
		return regexp.MustCompile(`([a-z0-9\-]+\.)+[a-z0-9\-]+`).FindString(CmdInput), "domain"
	} else {
		return "", "file"
	}
}
