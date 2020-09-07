package main

import (
	"fmt"
	"log"
	"strconv"
)

type Project struct {
	Name      string
	Language  string
	StarCount int
}

func GetGitHubUserRepos(username string) []Project {
	reposUrl, err := GetAndExtract(fmt.Sprintf("https://api.github.com/users/%s", username), "repos_url")
	if err != nil {
		log.Fatal(err)
	}

	repositoryResults := GetAsList(reposUrl)

	result := make([]Project, len(repositoryResults))

	for index, repos := range repositoryResults {
		starCount, _ := strconv.Atoi(fmt.Sprint(repos["stargazers_count"]))
		user := Project{
			Name:      fmt.Sprint(repos["full_name"]),
			Language:  fmt.Sprint(repos["Language"]),
			StarCount: starCount,
		}
		result[index] = user
	}
	return result
}
