package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(owner, repo string) (IssuesSearchResult, error) {
	api := new(GitHubAPI)
	api.setUrlPath("repos", owner, repo, "issues")

	resp, err := api.get()
	if err != nil {
		log.Print("API call error")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("response body: %s\n", body)
	var result IssuesSearchResult
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&result); err != nil {
		log.Print("Decode error")
		return nil, err
	}
	return result, nil
}

func ReadIssue(owner, repo string, number int) (*Issue, error) {
	api := new(GitHubAPI)
	api.setUrlPath("repos", owner, repo, "issues", strconv.Itoa(number))

	resp, err := api.get()
	if err != nil {
		log.Print("API call error")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("response body: %s\n", body)
	var issue Issue
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&issue); err != nil {
		log.Print("Decode error")
		return nil, err
	}
	return &issue, nil
}

func CreateIssue(owner, repo string) (*Issue, error) {
	api := new(GitHubAPI)
	api.setUrlPath("repos", owner, repo, "issues")

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

	resp, err := api.post(buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ret Issue
	if err = json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
