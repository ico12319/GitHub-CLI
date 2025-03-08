package runner

import (
	"GitHubCLI/gitHubRepos"
	"GitHubCLI/gitHubUser"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const SEARCH_FOR_USER_OPTION = "1"
const SHOW_REPOSITORIES = "2"
const TERMINATE_GITHUB = "3"
const SORT_BY_STARS = "1"
const SORT_BY_NAME = "2"
const SORT_BY_WATCHERS_COUNT = "3"
const SORT_BY_FORKS_COUNT = "4"
const SORT_BY_LANGUAGE = "5"
const SHOW_MOST_FAMOUS_REPO = "1"
const SHOW_STARS_GATHERED = "2"
const SHOW_SORTED_BY_CRITERIA = "3"
const SHOW_FILTERED_BY_LANGUAGE = "4"

func readInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}
	return input, nil
}

func handleReposRequest(response *http.Response) (*gitHubRepos.GitHubRepos, error) {
	repos, err := gitHubRepos.NewGitHubReposDatabase(response)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func searchForUser() (*gitHubUser.GitHubUser, error) {
	fmt.Print("Search user: ")
	username, err := readInput()
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(USER_PATH + username)
	if err != nil {
		return nil, err
	}
	foundUser, err := gitHubUser.NewGitHubUser(resp)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func getUserRepos(userName string) (*gitHubRepos.GitHubRepos, error) {
	resp, err := http.Get(USER_PATH + userName + "/repos")
	if err != nil {
		return nil, err
	}
	repos, err := gitHubRepos.NewGitHubReposDatabase(resp)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func handleSortingCriteria(criteria string, repos *gitHubRepos.GitHubRepos) error {
	if criteria == SORT_BY_STARS {
		repos.SortReposByCriteria(sortByStars)
	} else if criteria == SORT_BY_NAME {
		repos.SortReposByCriteria(sortByName)
	} else if criteria == SORT_BY_WATCHERS_COUNT {
		repos.SortReposByCriteria(sortByWatchersCount)
	} else if criteria == SORT_BY_FORKS_COUNT {
		repos.SortReposByCriteria(sortByForksCount)
	} else if criteria == SORT_BY_LANGUAGE {
		repos.SortReposByCriteria(sortByLanguage)
	} else {
		return fmt.Errorf("invalid criteria")
	}
	return nil
}

func handleTotalStarsRequest(repos *gitHubRepos.GitHubRepos, user *gitHubUser.GitHubUser) {
	stars := repos.GetTotalStarsEarned()
	fmt.Printf("%v managed to gather %d stars in total\n", user.Login, stars)
}

func handleMostStarredRepoRequest(repos *gitHubRepos.GitHubRepos) {
	mostFamousRepo := repos.GetMostStarredRepo()
	mostFamousRepo.ShowRepoInfo()
}

func printCriteriaInfo() {
	fmt.Println("Select a criteria: ")
	fmt.Println("1. Sort by stars gathered")
	fmt.Println("2. Sort by names")
	fmt.Println("3. Sort by watchers count")
	fmt.Println("4. Sort by forks count")
	fmt.Println("5. Sort by language")
}

func printActionsForReposInfo(user *gitHubUser.GitHubUser) {
	fmt.Println("Actions for repos: ")
	fmt.Printf("1.Show %v most famous repo\n", user.Login)
	fmt.Printf("2.Show %v stars gathered count\n", user.Login)
	fmt.Printf("3.Show %v repos sorted by criteria\n", user.Login)
	fmt.Printf("4.Show %v repos filtered by language\n", user.Login)
}

func printReposAfterSorting(repos *gitHubRepos.GitHubRepos, criteria string) error {
	err := handleSortingCriteria(criteria, repos)
	if err != nil {
		return err
	}
	repos.ShowReposInfo()
	return nil
}

func printReposAfterFilteredByLanguage(repos *gitHubRepos.GitHubRepos, language string) {
	reposByLanguage := repos.FilterByLanguage(language)
	reposByLanguage.ShowReposInfo()
}

func printGitHubFunctionalityInfo() {
	fmt.Println("Select option:")
	fmt.Println("1.Search for another user")
	fmt.Println("2.Show this user repositories")
	fmt.Println("3.Leave GitHub")
}

func handleReposAction(repos *gitHubRepos.GitHubRepos, user *gitHubUser.GitHubUser, reposAction string, reader *bufio.Reader) error {
	if reposAction == SHOW_MOST_FAMOUS_REPO {
		handleMostStarredRepoRequest(repos)
	} else if reposAction == SHOW_STARS_GATHERED {
		handleTotalStarsRequest(repos, user)
	} else if reposAction == SHOW_SORTED_BY_CRITERIA {
		printCriteriaInfo()
		criteria, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		criteria = strings.TrimSpace(criteria)
		err = printReposAfterSorting(repos, criteria)
		if err != nil {
			return err
		}
	} else if reposAction == SHOW_FILTERED_BY_LANGUAGE {
		fmt.Print("Enter a language ")
		language, err5 := reader.ReadString('\n')
		if err5 != nil {
			return err5
		}
		language = strings.TrimSpace(language)
		printReposAfterFilteredByLanguage(repos, language)
	}
	return nil
}

func handleUserSearch() error {
	user, err := searchForUser()
	if err != nil {
		return err
	}
	user.ShowUserInfo()
	return nil
}

func Run() error {
	user, err := searchForUser()
	if err != nil {
		return err
	}
	user.ShowUserInfo()
	for {
		printGitHubFunctionalityInfo()
		reader := bufio.NewReader(os.Stdin)
		choice, err2 := reader.ReadString('\n')
		if err2 != nil {
			return err2
		}
		choice = strings.TrimSpace(choice)
		if choice == SHOW_REPOSITORIES {
			repos, err3 := getUserRepos(user.Login)
			if err3 != nil {
				return err3
			}
			repos.ShowReposInfo()
			printActionsForReposInfo(user)
			reposAction, err3 := reader.ReadString('\n')
			if err3 != nil {
				return err3
			}
			reposAction = strings.TrimSpace(reposAction)
			err4 := handleReposAction(repos, user, reposAction, reader)
			if err4 != nil {
				return err4
			}
		} else if choice == SEARCH_FOR_USER_OPTION {
			user, err = searchForUser()
			if err != nil {
				return err
			}
			user.ShowUserInfo()
		} else if choice == TERMINATE_GITHUB {
			break
		}
	}
	return nil
}
