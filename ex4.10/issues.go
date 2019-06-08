// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tesujiro/ProgrammingLanguageGo/ex4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	oneMonthBefore := now.AddDate(0, -1, 0)
	oneYearBefore := now.AddDate(-1, 0, 0)
	ageCategory := func(t time.Time) string {
		switch {
		case t.After(oneMonthBefore):
			return "less than a month old"
		case t.After(oneYearBefore):
			return "iess than a year old"
		default:
			return "more than a year old"
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %v %9.9s %.55s\n",
			item.Number, ageCategory(item.CreatedAt), item.User.Login, item.Title)
	}
}
