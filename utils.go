package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func GetAndExtract(url, key string) (string, error) {
	m := Get(url)
	if m[key] == nil {
		return "", errors.New(key + " not found")
	}
	return fmt.Sprint(m[key]), nil
}

func GetAsList(url string) []map[string]interface{} {
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

func Get(url string) map[string]interface{} {
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

func Contains(slice []string, value string) bool {
	for _, i := range slice {
		if i == value {
			return true
		}
	}
	return false
}
