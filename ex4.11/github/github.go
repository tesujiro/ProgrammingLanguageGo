// Package github provides a Go API for the GitHub issue tracker.
package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
const APIURL = "https://api.github.com"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	User      *User
	CreatedAt time.Time `json:"created_at"`
	EditableIssue
}

type EditableIssue struct {
	Title string `json:"title"`
	State string `json:"state"`
	Body  string `json:"body"` // in Markdown format
	//Labels []string `json:"labels"`  // request labels and response labels are different
}

func (edit *EditableIssue) Edit() error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	tempfile, err := ioutil.TempFile("", "issue_crud")
	if err != nil {
		log.Fatal(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	encoder := json.NewEncoder(tempfile)
	/*
		err = encoder.Encode(map[string]string{
			"title": edit.Title,
			"state": edit.State,
			"body":  edit.Body,
		})
	*/
	err = encoder.Encode(edit)
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Reopen the file
	tempfile.Close()
	tempfile, err = os.Open(tempfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	_, err = tempfile.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewDecoder(tempfile).Decode(&edit); err != nil {
		fmt.Println("Decode error")
		log.Fatal(err)
	}
	/*
		fields := make(map[string]string)
		if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
			fmt.Println("Decode error")
			log.Fatal(err)
		}
		edit.title=fields["title"]
		edit.state=fields["state"]
		edit.body=fields["body"]
	*/
	//fmt.Printf("edit to post: %v\n", edit)
	return nil
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
