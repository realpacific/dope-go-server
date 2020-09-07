package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/user/:user", func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		fmt.Println(ps.ByName("user"))
		repo := GetGitHubUserRepos(ps.ByName("user"))
		response, err := json.Marshal(repo)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Write(response)
	})

	http.ListenAndServe(":3000", router)
}
