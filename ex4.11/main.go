package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tesujiro/ProgrammingLanguageGo/ex4.11/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("argument error")
		os.Exit(1)
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	switch {
	case cmd == "search" && len(args) == 1:
		repo := args[0]
		result, err := github.SearchIssues(repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d issues:\n", len(result))

		for _, item := range result {
			fmt.Printf("#%-5d %v %10.10s %.55s\n",
				item.Number, item.CreatedAt.In(time.Local), string(item.User["login"]), item.Title)
		}
	case cmd == "create" && len(args) == 2:
		owner, repo := args[0], args[1]
		result, err := github.CreateIssue(owner, repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("new issue: #%v\n", (*result).Number)

	default:
		fmt.Println("argument error")
		os.Exit(1)
	}

}
