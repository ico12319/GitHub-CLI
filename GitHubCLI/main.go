package main

import "GitHubCLI/runner"

func main() {

	err := runner.Run()
	if err != nil {
		panic(true)
	}

}
