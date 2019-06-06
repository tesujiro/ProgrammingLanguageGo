package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func CreateIssue(owner, repo string) (*Issue, error) {
	buf := &bytes.Buffer{}
	/*
		encoder := json.NewEncoder(buf)
			err := encoder.Encode(fields)
			if err != nil {
				return nil, err
			}
	*/

	client := &http.Client{}
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	req, err := http.NewRequest("PATCH", url, buf)
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to edit issue: %s", resp.Status)
	}
	var issue Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil

}
