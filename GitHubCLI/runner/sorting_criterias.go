package runner

import "GitHubCLI/gitHubRepos"

func sortByStars(r1 gitHubRepos.GitHubRepo, r2 gitHubRepos.GitHubRepo) bool {
	return r1.StarGazersCount > r2.StarGazersCount
}

func sortByName(r1 gitHubRepos.GitHubRepo, r2 gitHubRepos.GitHubRepo) bool {
	return r1.Name > r2.Name
}

func sortByWatchersCount(r1 gitHubRepos.GitHubRepo, r2 gitHubRepos.GitHubRepo) bool {
	return r1.WatchersCount > r2.WatchersCount
}

func sortByForksCount(r1 gitHubRepos.GitHubRepo, r2 gitHubRepos.GitHubRepo) bool {
	return r1.ForksCount > r2.ForksCount
}

func sortByLanguage(r1 gitHubRepos.GitHubRepo, r2 gitHubRepos.GitHubRepo) bool {
	return r1.Language > r2.Language
}
