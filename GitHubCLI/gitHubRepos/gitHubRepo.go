package gitHubRepos

import (
	"GitHubCLI/gitHubUser"
	"fmt"
)

type GitHubRepo struct {
	Name            string                `json:"name"`
	Owner           gitHubUser.GitHubUser `json:"owner"`
	Description     string                `json:"description"`
	StarGazersCount int                   `json:"stargazers_count"`
	WatchersCount   int                   `json:"watchers_count"`
	Language        string                `json:"language"`
	ForksCount      int                   `json:"forks_count"`
}

func (repo *GitHubRepo) ShowRepoInfo() {
	fmt.Printf("repository name😎: %v\n", repo.Name)
	fmt.Printf("owner🤠: %v\n", repo.Owner.Login)
	fmt.Printf("description👨‍💻: %v\n", repo.Description)
	fmt.Printf("stars🤩: %v\n", repo.StarGazersCount)
	fmt.Printf("watchers👁: %v\n", repo.WatchersCount)
	fmt.Printf("language🧠: %v\n", repo.Language)
	fmt.Printf("forks Count🍴: %v\n", repo.ForksCount)
}
