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

func CreateIssue(owner, repo, number string) (*Issue, error) {
	issue := EditIssue{Title: "gopl exercise 4.11"}
	if err := issue.Edit(); err != nil {
		log.Fatal(err)
	}
	/*
		var jsonStr = []byte(`
			{
				"title":"Buy cheese and bread for breakfast.",
				"labels": ["gopl"]
			}`)
	*/
	jsonStr, err := json.Marshal(issue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("json: %s\n", jsonStr)
	buf := bytes.NewBuffer(jsonStr)
	/*
		encoder := json.NewEncoder(buf)
			err := encoder.Encode(fields)
			if err != nil {
				return nil, err
			}
	*/

	client := &http.Client{}
	//url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues"}, "/")
	fmt.Println("url=", url)
	req, err := http.NewRequest("POST", url, buf)
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
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
	var ret Issue
	if err = json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil

}
