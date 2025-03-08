package gitHubRepos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type GitHubRepos struct {
	repos []GitHubRepo
}

func parseGitHubRepos(response *http.Response) ([]GitHubRepo, error) {
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not okay status code: %v", response.StatusCode)
	}
	var repositories []GitHubRepo
	err := json.NewDecoder(response.Body).Decode(&repositories)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}

func NewGitHubReposDatabase(response *http.Response) (*GitHubRepos, error) {
	parsedRepositories, err := parseGitHubRepos(response)
	if err != nil {
		return nil, err
	}
	return &GitHubRepos{parsedRepositories}, nil
}

func (reposData *GitHubRepos) ShowReposInfo() {
	for i, repo := range reposData.repos {
		fmt.Printf("%v. ", i+1)
		repo.ShowRepoInfo()
	}
}

func (reposData *GitHubRepos) GetTotalStarsEarned() int {
	totalStars := 0
	for _, repo := range reposData.repos {
		totalStars += repo.StarGazersCount
	}
	return totalStars
}

func (reposData *GitHubRepos) SortReposByCriteria(criteria func(r1, r2 GitHubRepo) bool) {
	sort.Slice(reposData.repos, func(i, j int) bool {
		return criteria(reposData.repos[i], reposData.repos[j])
	})
}

func (reposData *GitHubRepos) GetMostStarredRepo() *GitHubRepo {
	mostStarsGathered := 0
	var mostStarredRepo *GitHubRepo = nil
	for _, repo := range reposData.repos {
		if repo.StarGazersCount > mostStarsGathered {
			mostStarsGathered = repo.StarGazersCount
			mostStarredRepo = &repo
		}
	}
	return mostStarredRepo
}

func (reposData *GitHubRepos) FilterByLanguage(language string) *GitHubRepos {
	languageSpecificRepos := make([]GitHubRepo, 0, 8)
	for _, repo := range reposData.repos {
		if repo.Language == language {
			languageSpecificRepos = append(languageSpecificRepos, repo)
		}
	}
	return &GitHubRepos{languageSpecificRepos}
}

func (repos *GitHubRepos) GetRepos() []GitHubRepo {
	return repos.repos
}
