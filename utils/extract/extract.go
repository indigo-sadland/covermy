package extract

import (
	"regexp"
	"strings"
)

func Target(CmdInput string) (string, string) {
	CmdInput = strings.TrimSpace(CmdInput)

	// If input contains "http(s)://", strip it only to URL domain
	if regexp.MustCompile(`\bhttps?://`).MatchString(CmdInput) {
		CmdInput = regexp.MustCompile(`\bhttps?://(.*?)/?`).FindStringSubmatch(CmdInput)[1]
	}

	// Strip "www." from domain
	if regexp.MustCompile(`\bwww\.`).MatchString(CmdInput) {
		CmdInput = regexp.MustCompile(`\bwww\.`).ReplaceAllString(CmdInput, "")
	}

	// Check if target is an IP address
	if regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}`).MatchString(CmdInput) {
		return regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}`).FindString(CmdInput), "ip"
	}

	// Extract domain name
	if regexp.MustCompile(`([a-z0-9\-]+\.)+[a-z0-9\-]+`).FindString(CmdInput) != "" {
		return regexp.MustCompile(`([a-z0-9\-]+\.)+[a-z0-9\-]+`).FindString(CmdInput), "domain"
	} else {
		return "", "file"
	}
}
