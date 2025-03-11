package main

import (
	"GitHubCLI/runner"
	"fmt"
)

func main() {

	//url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", "ico12319", "FilterKit")

	err := runner.Run()
	if err != nil {
		fmt.Println("Fatal error!")
		return
	}

	//err := runner.Run()
	//if err != nil {
	//	panic(true)
	//}

	//colorReset := "\033[0m"
	//fmt.Println(colorReset)

}
