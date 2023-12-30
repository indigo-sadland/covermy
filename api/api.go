package api

import (
	"bytes"
	"covermy/utils/errors"
	"covermy/utils/logcamp"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	coverMyHome      = "/.local/share/covermy/api"
	notesEndpoint    = "http://localhost:41184/notes?token="
	notebookEndpoint = "http://localhost:41184/folders?token="
	searchEndpoint   = "http://localhost:41184/search?query="
)

var key string

type Notebooks struct {
	Items []struct {
		ID       string `json:"id"`
		ParentID string `json:"parent_id"`
		Title    string `json:"title"`
	} `json:"items"`
	HasMore bool `json:"has_more"`
}

type Notebook struct {
	Title           string `json:"title"`
	ID              string `json:"id"`
	UpdatedTime     int64  `json:"updated_time"`
	CreatedTime     int64  `json:"created_time"`
	UserUpdatedTime int64  `json:"user_updated_time"`
	UserCreatedTime int64  `json:"user_created_time"`
	Type            int    `json:"type_"`
}

// GetKey extracts Joplin API key.
func GetKey() {
	userHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to detect user home directory")
		os.Exit(-1)
	}

	path := userHome + coverMyHome
	data, err := os.ReadFile(path)
	if err != nil {
		e := fmt.Errorf("unable to get API key: %s", err.Error())
		fmt.Println(e)
		os.Exit(-1)
	}

	key = strings.Trim(string(data), "\n")
}

func ListNotebooks() (Notebooks, error) {
	var notebooks Notebooks

	resp, err := http.Get(notebookEndpoint + key)
	if err != nil {
		fmt.Println(errors.LISTNOTEBOOKFAIL)
		logcamp.ErrorLogger.Println(err.Error())
		return Notebooks{}, err
	}
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &notebooks); err != nil {
		fmt.Println(errors.UNMARSHALFAIL)
		logcamp.ErrorLogger.Println(err.Error())
		return Notebooks{}, err
	}
	return notebooks, nil
}

// CheckNotebook gets information about a notebook if exists.
func CheckNotebook(notebookName string) (string, error) {
	var notebooks Notebooks
	var notebookID string

	resp, err := http.Get(searchEndpoint + notebookName + "&type=folder" + "&token=" + key)
	if err != nil {
		fmt.Println(errors.GETNOTEBOOKFAIL)
		logcamp.ErrorLogger.Println(err.Error())
		return "", err
	}
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &notebooks); err != nil {
		fmt.Println(errors.UNMARSHALFAIL)
		logcamp.ErrorLogger.Println(err.Error())
		return "", err
	}
	// TODO: what if a user has multiple notebooks with the same name?
	if len(notebooks.Items) != 0 {
		for _, v := range notebooks.Items {
			notebookID = v.ID
		}
	} else {
		notebookID = createNewNotebook(notebookName)
	}

	return notebookID, nil
}

func createNewNotebook(notebookName string) string {
	var notebook Notebook

	title := notebookName
	content := map[string]string{"title": title}
	jsonData, err := json.Marshal(content)
	if err != nil {
		logcamp.ErrorLogger.Println(err.Error())
	}

	resp, err := http.Post(notebookEndpoint+key, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(errors.CREATENOTEBOOKFAIL)
		logcamp.ErrorLogger.Println(err.Error())
		os.Exit(-1)
	}
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &notebook); err != nil {
		fmt.Println(errors.UNMARSHALFAIL)
		logcamp.ErrorLogger.Println(err.Error())
		os.Exit(-1)
	}

	return notebook.ID
}

func CreateNewNote(noteName, notebookId string, input []string) error {
	title := noteName
	body := strings.Join(input, "\n")
	// Format body as code block
	body = "```" + body
	content := map[string]string{"title": title, "parent_id": notebookId, "body": body}
	jsonData, err := json.Marshal(content)
	if err != nil {
		fmt.Println(errors.MARSHALFAIL)
		logcamp.ErrorLogger.Println(err.Error())
		return err
	}

	resp, err := http.Post(notesEndpoint+key, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(errors.CREATENEWNOTEFAIL)
		logcamp.ErrorLogger.Println(err.Error())
		return err
	}
	// log response from Joplin.
	respBody, _ := io.ReadAll(resp.Body)
	logcamp.InfoLogger.Writer().Write(respBody)

	return nil
}
