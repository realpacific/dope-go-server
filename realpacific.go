package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var KEYS = []string{"full_name", "language", "stargazers_count"}

func main() {
	repos, err := getAndExtract("https://api.github.com/users/realpacific", "repos_url")
	if err != nil {
		log.Fatal(err)
	}
	repositoryResults := getAsList(repos)

	for _, result := range repositoryResults {
		fmt.Println("------------")
		for k, v := range result {
			if contains(KEYS, k) {
				fmt.Println(strings.Title(k)+":", v)
			}
		}
	}
}

func getAndExtract(url, key string) (string, error) {
	m := get(url)
	fmt.Println(m)
	if m[key] == nil {
		return "", errors.New(key + " not found")
	}
	return fmt.Sprint(m[key]), nil
}

func getAsList(url string) []map[string]interface{} {
	fmt.Println("Getting", url, "...")
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal(err)
	}

	var result []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func get(url string) map[string]interface{} {
	fmt.Println("Getting", url, "...")
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		if resp != nil {
			fmt.Println(resp.StatusCode)
		}
		log.Fatal(err)
	}

	var result map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&result)
	return result
}

func contains(slice []string, value string) bool {
	for _, i := range slice {
		if i == value {
			return true
		}
	}
	return false
}
