package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func CreateIssue(owner, repo string) (*Issue, error) {
	issue := EditableIssue{Title: "gopl exercise 4.11", State: "open"}
	if err := issue.Edit(); err != nil {
		log.Fatal(err)
	}
	jsonStr, err := json.Marshal(issue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("json: %s\n", jsonStr)
	buf := bytes.NewBuffer(jsonStr)

	client := &http.Client{}
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	//fmt.Println("url=", url)
	req, err := http.NewRequest("POST", url, buf)
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
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to edit issue: %s", resp.Status)
	}
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("Response Body: %s\n", body)
	var ret Issue
	if err = json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
