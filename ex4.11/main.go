package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tesujiro/ProgrammingLanguageGo/ex4.11/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("argument error")
		os.Exit(1)
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" && len(args) > 0 {
		result, err := github.SearchIssues(args)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
		os.Exit(0)
	}
	if len(args) < 3 {
		fmt.Println("argument error 2")
		os.Exit(1)
	}

	switch {
	case cmd == "create" && len(args) == 2:
		owner, repo := args[0], args[1]
		result, err := github.CreateIssue(owner, repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("issue: %v\n", *result)

	default:
		fmt.Println("argument error 3")
		os.Exit(1)
	}

}
