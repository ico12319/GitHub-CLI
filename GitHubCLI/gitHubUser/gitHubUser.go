package gitHubUser

import (
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
	fmt.Printf("username👽: %v\n", user.Login)
	fmt.Printf("full name🧐: %v\n", user.Name)
	fmt.Printf("location👁: %v\n", user.Location)
	fmt.Printf("public repositories count✍️: %v\n", user.PublicRepos)
	fmt.Printf("followers count🤳: %v\n", user.Followers)
	fmt.Printf("following🕵️‍♂️: %v\n", user.Following)
}
