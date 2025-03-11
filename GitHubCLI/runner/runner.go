package runner

import (
	commits2 "GitHubCLI/commits"
	"GitHubCLI/constants"
	"GitHubCLI/gitHubRepos"
	"GitHubCLI/gitHubUser"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func getUsersCommitsForRepo(userLogin string, repoName string) (*commits2.Commits, error) {
	resp, err := http.Get(COMMITS_PATH + userLogin + "/" + repoName + "/commits")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not okay status code %v", resp.StatusCode)
	}
	commits, err := commits2.NewCommits(resp)
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func printGitHubHeader() {
	fmt.Println(constants.Bold + constants.ColorCyan + "===================================" + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorCyan + "       GitHub CLI Application      " + constants.ColorReset)
	fmt.Println(constants.Bold + constants.ColorCyan + "===================================" + constants.ColorReset)
}

func searchForUser(reader *bufio.Reader) (*gitHubUser.GitHubUser, error) {
	fmt.Print("Search user: ")
	username, err := readInput(reader)
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
	fmt.Printf(constants.Bold+constants.ColorYellow+"5. Show %v's commits in a selected repo\n"+constants.ColorReset, user.Login)
	fmt.Printf(constants.Bold + constants.ColorYellow + "6. Back\n" + constants.ColorReset)
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

func handleInvalidRepoNameWhenSearchingCommits(user *gitHubUser.GitHubUser, repoName string, reader *bufio.Reader) {
	for {
		commits, err := getUsersCommitsForRepo(user.Login, repoName)
		if err != nil {
			fmt.Println("Invalid repo name!")
			fmt.Print("Enter repo's name: ")
			repoName, err = reader.ReadString('\n')
			repoName = strings.TrimSpace(repoName)
			if err != nil {
				return
			}
			continue
		}
		commits.ShowCommits()
		break
	}
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
		language, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		language = strings.TrimSpace(language)
		printReposAfterFilteredByLanguage(repos, language)
	} else if reposAction == constants.SHOW_COMMITS {
		fmt.Print(constants.Bold + constants.ColorBlue + "Enter repo name: " + constants.ColorReset)
		repoName, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		repoName = strings.TrimSpace(repoName)
		handleInvalidRepoNameWhenSearchingCommits(user, repoName, reader)
	}
	return nil
}

func handleUserSearch(reader *bufio.Reader) *gitHubUser.GitHubUser {
	for {
		us, err := searchForUser(reader)
		if err != nil {
			continue
		}
		us.ShowUserInfo()
		return us
	}
}

func Run() error {
	printGitHubHeader()
	reader := bufio.NewReader(os.Stdin)
	user := handleUserSearch(reader)
	for {
		printGitHubFunctionalityInfo()
		choice, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		choice = strings.TrimSpace(choice)
		if choice == constants.SHOW_REPOSITORIES {
			repos, err := getUserRepos(user.Login)
			if err != nil {
				return err
			}
			repos.ShowReposInfo()
			for {
				printActionsForReposInfo(user)
				reposAction, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				reposAction = strings.TrimSpace(reposAction)
				if reposAction == constants.GO_BACK {
					break
				}
				err = handleReposAction(repos, user, reposAction, reader)
				if err != nil {
					return err
				}
			}
		} else if choice == constants.SEARCH_FOR_USER_OPTION {
			user, err = searchForUser(reader)
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
