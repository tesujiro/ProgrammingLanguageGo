package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(repo string) (IssuesSearchResult, error) {
	// GitHub API V3
	url := strings.Join([]string{APIURL, "repos", repo, "issues"}, "/")
	fmt.Println("url:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("response body: %s\n", body)
	var result IssuesSearchResult
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&result); err != nil {
		fmt.Println("Decode error")
		return nil, err
	}
	return result, nil
}
