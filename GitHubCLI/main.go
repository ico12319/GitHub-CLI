package main

import (
	commits2 "GitHubCLI/commits"
	"fmt"
	"net/http"
)

func main() {

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", "ico12319", "LeetCode-Tasks")
	resp, err2 := http.Get(url)
	if err2 != nil {
		panic(true)
	}
	commits, err := commits2.NewCommits(resp)
	if err != nil {
		panic(true)
	}
	commits.ShowCommits()
	//err := runner.Run()
	//if err != nil {
	//	panic(true)
	//}

	//colorReset := "\033[0m"
	//fmt.Println(colorReset)

}
