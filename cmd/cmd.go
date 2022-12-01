package cmd

import (
	"bufio"
	"bytes"
	"github.com/indigo-sadland/covermy/utils/logcamp"
	"github.com/indigo-sadland/covermy/utils/normalizer"
	"fmt"
	"os/exec"
	"strings"
)

func Run(command []string) []string {
	var stderr bytes.Buffer
	var rawlines []string

	bin := command[0]
	args := command[1:]
	cmd := exec.Command(bin, args...)
	cmd.Stderr = &stderr

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		logcamp.ErrorLogger.Printf(err.Error())
	}
	// Read cmd output in real-time
	scanner := bufio.NewScanner(cmdReader)
	// Add the tool command to the list. Needed for the finale output.
	rawlines = append(rawlines, fmt.Sprintf("\n"+strings.Join(command, " ")+"\n"))
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			rawlines = append(rawlines, line)
		}
	}()

	if err := cmd.Run(); err != nil {
		fmt.Printf("\n")
		fmt.Println(&stderr)
	}

	output := normalizer.Start(bin, rawlines)
	return output
}
