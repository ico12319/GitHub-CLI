package commits

import (
	"GitHubCLI/gitHubUser"
	"encoding/json"
	"fmt"
	"net/http"
)

type Commit struct {
	Detail    CommitDetail           `json:"commit"`
	Author    *gitHubUser.GitHubUser `json:"author"`
	Committer *gitHubUser.GitHubUser `json:"committer"`
}

func parseCommitFromResponse(response *http.Response) (*Commit, error) {
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok status code %v", response.StatusCode)
	}
	var com Commit
	err := json.NewDecoder(response.Body).Decode(&com)
	if err != nil {
		return nil, err
	}
	return &com, nil
}

func NewCommit(response *http.Response) (*Commit, error) {
	commit, err := parseCommitFromResponse(response)
	if err != nil {
		return nil, err
	}
	return &Commit{commit.Detail, commit.Author, commit.Committer}, nil
}

func (c *Commit) ShowCommitInfo() {
	fmt.Println(c.Detail.Message)
	if c.Author != nil {
		fmt.Printf("Made by %s\n", c.Author.Login)
	}
}
