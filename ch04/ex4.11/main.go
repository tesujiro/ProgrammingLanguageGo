package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tesujiro/TheGoProgrammingLanguage/ch04/ex4.11/github"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Print(`	issue search [owner] [repo]
	issue create [owner] [repo]
	issue read   [owner] [repo] [issue #]
	issue update [owner] [repo] [issue #]
	issue close  [owner] [repo] [issue #]
Environ variables "GITHUB_USER", "GITHUB_PASS", "EDITOR" are used.
`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	switch {
	case cmd == "search" && len(args) == 2:
		owner, repo := args[0], args[1]
		result, err := github.SearchIssues(owner, repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d issues:\n", len(result))

		for _, issue := range result {
			fmt.Printf("#%-5d %v %10.10s %.55s\n",
				issue.Number, issue.CreatedAt.In(time.Local), string(issue.User["login"]), issue.Title)
		}
	case cmd == "create" && len(args) == 2:
		owner, repo := args[0], args[1]
		result, err := github.CreateIssue(owner, repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("new issue: #%v\n", (*result).Number)

	case cmd == "read" && len(args) == 3:
		owner, repo, numberStr := args[0], args[1], args[2]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatalf("issue number format error: %v", numberStr)
		}
		issue, err := github.ReadIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("#%-5d %v %10.10s %.55s\n",
			issue.Number, issue.CreatedAt.In(time.Local), string(issue.User["login"]), issue.Title)

	case cmd == "update" && len(args) == 3:
		owner, repo, numberStr := args[0], args[1], args[2]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatalf("issue number format error: %v", numberStr)
		}
		issue, err := github.UpdateIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("#%-5d %v %10.10s %.55s\n",
			issue.Number, issue.CreatedAt.In(time.Local), string(issue.User["login"]), issue.Title)

	case cmd == "close" && len(args) == 3:
		owner, repo, numberStr := args[0], args[1], args[2]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatalf("issue number format error: %v", numberStr)
		}
		issue, err := github.CloseIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("#%-5d closed.\n", issue.Number)

	default:
		usage()
		os.Exit(1)
	}

}
