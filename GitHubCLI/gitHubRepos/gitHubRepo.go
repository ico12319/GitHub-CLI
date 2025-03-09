package gitHubRepos

import (
	"GitHubCLI/constants"
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
	fmt.Printf("%sRepository Name😎:%s %v\n", constants.ColorMagenta, constants.ColorReset, repo.Name)
	fmt.Printf("%sOwner🤠:%s %v\n", constants.ColorBlue, constants.ColorReset, repo.Owner.Login)
	fmt.Printf("%sDescription👨‍💻:%s %v\n", constants.ColorGreen, constants.ColorReset, repo.Description)
	fmt.Printf("%sStars🤩:%s %v\n", constants.ColorYellow, constants.ColorReset, repo.StarGazersCount)
	fmt.Printf("%sWatchers👁:%s %v\n", constants.ColorRed, constants.ColorReset, repo.WatchersCount)
	fmt.Printf("%sLanguage🧠:%s %v\n", constants.ColorCyan, constants.ColorReset, repo.Language)
	fmt.Printf("%sForks Count🍴:%s %v\n", constants.ColorWhite, constants.ColorReset, repo.ForksCount)
}
