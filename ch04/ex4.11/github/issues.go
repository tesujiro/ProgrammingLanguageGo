package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(owner, repo string) (IssuesSearchResult, error) {
	api := new(GitHubAPI)
	api.setUrlPath("repos", owner, repo, "issues")

	resp, err := api.get()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("response body: %s\n", body)
	var result IssuesSearchResult
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateIssue(owner, repo string) (*Issue, error) {
	api := new(GitHubAPI)
	api.setUrlPath("repos", owner, repo, "issues")

	issue := EditableIssue{Title: "gopl exercise 4.11", State: "open"}
	if err := issue.Edit(); err != nil {
		return nil, fmt.Errorf("Edit issue error: %v", err)
	}
	jsonStr, err := json.Marshal(issue)
	if err != nil {
		return nil, fmt.Errorf("Marshal issue error: %v", err)
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
		return nil, fmt.Errorf("Unmarshal http response error: %v", err)
	}
	return &ret, nil
}

func ReadIssue(owner, repo string, number int) (*Issue, error) {
	api := new(GitHubAPI)
	api.setUrlPath("repos", owner, repo, "issues", strconv.Itoa(number))

	resp, err := api.get()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("response body: %s\n", body)
	var issue Issue
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&issue); err != nil {
		return nil, fmt.Errorf("Unmarshal http response error: %v", err)
	}
	return &issue, nil
}

func UpdateIssue(owner, repo string, number int) (*Issue, error) {
	issue, err := ReadIssue(owner, repo, number)
	if err != nil {
		return nil, fmt.Errorf("Read issue error: %v", err)
	}

	api := new(GitHubAPI)
	api.setUrlPath("repos", owner, repo, "issues", strconv.Itoa(number))

	if err := issue.EditableIssue.Edit(); err != nil {
		return nil, fmt.Errorf("Edit issue error: %v", err)
	}
	jsonStr, err := json.Marshal(issue.EditableIssue)
	if err != nil {
		return nil, fmt.Errorf("Marshal issue error: %v", err)
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
		return nil, fmt.Errorf("Unmarshal http response error: %v", err)
	}
	return &ret, nil
}

func CloseIssue(owner, repo string, number int) (*Issue, error) {
	issue, err := ReadIssue(owner, repo, number)
	if err != nil {
		return nil, fmt.Errorf("Read issue error: %v", err)
	}

	if issue.State == "closed" {
		return nil, fmt.Errorf("issue #%v already closed", number)
	}

	api := new(GitHubAPI)
	api.setUrlPath("repos", owner, repo, "issues", strconv.Itoa(number))

	issue.State = "close"
	jsonStr, err := json.Marshal(issue.EditableIssue)
	if err != nil {
		return nil, fmt.Errorf("Marshal issue error: %v", err)
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
		return nil, fmt.Errorf("Unmarshal http response error: %v", err)
	}
	return &ret, nil
}
