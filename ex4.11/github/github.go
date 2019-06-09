// Package github provides a Go API for the GitHub issue tracker.
package github

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// GitHub API V3 URL
const APIURL = "https://api.github.com"

type GitHubAPI struct {
	url string
}

func (api *GitHubAPI) setUrlPath(list ...string) {
	urlpath := append([]string{APIURL}, list...)
	api.url = strings.Join(urlpath, "/")
}

func (api *GitHubAPI) get() (*http.Response, error) {
	resp, err := http.Get(api.url)

	if resp.StatusCode != http.StatusOK {
		log.Print("API response status error")
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	return resp, err
}

func (api *GitHubAPI) post(body io.Reader) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", api.url, body)

	github_user := os.Getenv("GITHUB_USER")
	github_pass := os.Getenv("GITHUB_PASS")
	if github_user == "" || github_pass == "" {
		log.Fatal("env not set: GITHUB_USER, GITHUB_PASS")
	}
	req.SetBasicAuth(github_user, github_pass)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		log.Print("API response status error")
		return nil, fmt.Errorf("failed to edit issue: %s", resp.Status)
	}
	return resp, err
}

type IssuesSearchResult []Issue

type Issue struct {
	Number    int    `json:"number"`
	HTMLURL   string `json:"url"`
	User      map[string]json.RawMessage
	CreatedAt time.Time `json:"created_at"`
	EditableIssue
}

type EditableIssue struct {
	Title string `json:"title"`
	State string `json:"state"`
	Body  string `json:"body"` // in Markdown format
	//Labels []string `json:"labels"`  // request labels ([]string) and response labels (map[string]RawMessage) are different
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
	return nil
}
