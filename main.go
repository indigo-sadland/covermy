package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/indigo-sadland/covermy/api"
	"github.com/indigo-sadland/covermy/cmd"
	"github.com/indigo-sadland/covermy/utils/extract"
	"github.com/indigo-sadland/covermy/utils/logcamp"
	"github.com/jpillora/go-tld"
	"os"
	"strings"
)

const Usage = `Usage:
covermy [DESIRED TOOL COMMAND]
`

func main() {
	// Show usage text.
	if len(os.Args) == 1 {
		fmt.Println(Usage, os.Args[0])
		os.Exit(-1)
	}

	// Initialize logger.
	logcamp.Init()

	// Obtain API key.
	api.GetKey()

	// Get target name specified in a tool command.
	command := os.Args[1:]
	target, targetType := extract.Target(strings.Join(command, " "))

	// Get top level domain to use it as a root notebook.
	var rootNotebook string
	var notebookId string
	if targetType == "domain" {
		root, _ := tld.Parse("http://" + target)
		rootNotebook = root.Domain
	} else if targetType == "ip" {
		rootNotebook = target
	} else {
		if targetType == "file" {
			fmt.Printf("It seems like you use a file as target input. You need to specify target name manualy" +
				" (it will be used as a name for new note).\n")
			fmt.Printf("Target name:")

			// Parse user input
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				target = scanner.Text()
			}

		}
		fmt.Printf("Unable to detect root notebook. Where do you want to store the command result?\n")
		notebooks, err := api.ListNotebooks()
		if err != nil {
			os.Exit(-1)
		}
		i := len(notebooks.Items)
		if i != 0 {
			c := color.New(color.FgGreen)
			fmt.Printf("Available Notebooks:\n")
			for k, v := range notebooks.Items {
				c.Printf("[%d] %s\n", k, v.Title)
			}
			c.Printf("[%d] Create New Notebook\n", i)

			var index int
			fmt.Printf("\nEnter the index number:")
			fmt.Scanln(&index)
			if index != i {
				item := notebooks.Items[index]
				rootNotebook = item.Title
			} else {
				fmt.Printf("Enter name (without spaces) for new notebook:")
				fmt.Scanln(&rootNotebook)
			}

		}
	}
	// Get ID of the root notebook.
	notebookId, err := api.CheckNotebook(rootNotebook)
	if err != nil {
		os.Exit(-1)
	}

	// Execute the tool command.
	output := cmd.Run(command)

	// Note name is a target name plus a tool name
	noteName := fmt.Sprintf(target+" [%s]", command[0])
	api.CreateNewNote(noteName, notebookId, output)

}
