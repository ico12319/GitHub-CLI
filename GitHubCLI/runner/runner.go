package runner

import (
	"GitHubCLI/constants"
	"GitHubCLI/gitHubRepos"
	"GitHubCLI/gitHubUser"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func readInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}
	return input, nil
}

func printGitHubHeader() {
	fmt.Println(constants.Bold + constants.ColorCyan + "===================================" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorCyan + "       GitHub CLI Application      " + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorCyan + "===================================" + constants.ColorReset)
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
	if criteria == constants.SORT_BY_STARS {
		repos.SortReposByCriteria(sortByStars)
	} else if criteria == constants.SORT_BY_NAME {
		repos.SortReposByCriteria(sortByName)
	} else if criteria == constants.SORT_BY_WATCHERS_COUNT {
		repos.SortReposByCriteria(sortByWatchersCount)
	} else if criteria == constants.SORT_BY_FORKS_COUNT {
		repos.SortReposByCriteria(sortByForksCount)
	} else if criteria == constants.SORT_BY_LANGUAGE {
		repos.SortReposByCriteria(sortByLanguage)
	} else {
		return fmt.Errorf("invalid criteria")
	}
	return nil
}

func handleTotalStarsRequest(repos *gitHubRepos.GitHubRepos, user *gitHubUser.GitHubUser) {
	stars := repos.GetTotalStarsEarned()
	fmt.Printf("%v managed to gather %v%d stars in total\n", user.Login, constants.Bold+constants.ColorYellow, stars)
}

func handleMostStarredRepoRequest(repos *gitHubRepos.GitHubRepos) {
	mostFamousRepo := repos.GetMostStarredRepo()
	mostFamousRepo.ShowRepoInfo()
}

func printCriteriaInfo() {
	fmt.Println(constants.Bold + constants.ColorGreen + "Select a criteria:" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorBlue + "1. Sort by stars gathered" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorBlue + "2. Sort by names" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorBlue + "3. Sort by watchers count" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorBlue + "4. Sort by forks count" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorBlue + "5. Sort by language" + constants.ColorReset)
}

func printActionsForReposInfo(user *gitHubUser.GitHubUser) {
	fmt.Println(constants.Bold + constants.ColorMagenta + "Actions for " + user.Login + "'s repositories:" + constants.ColorReset)
	fmt.Printf(constants.Bold+constants.ColorYellow+"1. Show %v's most famous repo\n"+constants.ColorReset, user.Login)
	fmt.Printf(constants.Bold+constants.ColorYellow+"2. Show %v's stars gathered count\n"+constants.ColorReset, user.Login)
	fmt.Printf(constants.Bold+constants.ColorYellow+"3. Show %v's repos sorted by criteria\n"+constants.ColorReset, user.Login)
	fmt.Printf(constants.Bold+constants.ColorYellow+"4. Show %v's repos filtered by language\n"+constants.ColorReset, user.Login)
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
	fmt.Println(constants.Bold + constants.ColorGreen + "Select an option:" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorBlue + "1. Search for another user" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorBlue + "2. Show this user's repositories" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorBlue + "3. Leave GitHub" + constants.ColorReset)
}

func handleReposAction(repos *gitHubRepos.GitHubRepos, user *gitHubUser.GitHubUser, reposAction string, reader *bufio.Reader) error {
	if reposAction == constants.SHOW_MOST_FAMOUS_REPO {
		handleMostStarredRepoRequest(repos)
	} else if reposAction == constants.SHOW_STARS_GATHERED {
		handleTotalStarsRequest(repos, user)
	} else if reposAction == constants.SHOW_SORTED_BY_CRITERIA {
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
	} else if reposAction == constants.SHOW_FILTERED_BY_LANGUAGE {
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

func Run() error {
	printGitHubHeader()
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
		if choice == constants.SHOW_REPOSITORIES {
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
		} else if choice == constants.SEARCH_FOR_USER_OPTION {
			user, err = searchForUser()
			if err != nil {
				return err
			}
			user.ShowUserInfo()
		} else if choice == constants.TERMINATE_GITHUB {
			break
		}
	}
	return nil
}
