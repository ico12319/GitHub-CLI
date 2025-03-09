package gitHubUser

import (
	"GitHubCLI/constants"
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
}

func parseGitHubUser(response *http.Response) (*GitHubUser, error) {
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-okay response code %v", response.StatusCode)
	}
	var user GitHubUser
	err := json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewGitHubUser(response *http.Response) (*GitHubUser, error) {
	parsedUser, err := parseGitHubUser(response)
	if err != nil {
		return nil, err
	}
	return &GitHubUser{parsedUser.Login, parsedUser.Name, parsedUser.Location, parsedUser.PublicRepos, parsedUser.Followers, parsedUser.Following}, nil
}

func (user *GitHubUser) ShowUserInfo() {
	fmt.Printf("%sUsername👽:%s %v\n", constants.ColorMagenta, constants.ColorReset, user.Login)
	fmt.Printf("%sFull Name🧐:%s %v\n", constants.ColorBlue, constants.ColorReset, user.Name)
	fmt.Printf("%sLocation👁:%s %v\n", constants.ColorGreen, constants.ColorReset, user.Location)
	fmt.Printf("%sPublic Repositories Count✍️:%s %v\n", constants.ColorYellow, constants.ColorReset, user.PublicRepos)
	fmt.Printf("%sFollowers Count🤳:%s %v\n", constants.ColorCyan, constants.ColorReset, user.Followers)
	fmt.Printf("%sFollowing🕵️‍♂️:%s %v\n", constants.ColorRed, constants.ColorReset, user.Following)
}
