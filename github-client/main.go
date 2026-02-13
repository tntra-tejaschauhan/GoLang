package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Struct to map JSON response
type GitHubUser struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	PublicRepos int    `json:"public_repos"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <github-username>")
		return
	}

	username := os.Args[1]

	url := "https://api.github.com/users/" + username

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: user not found")
		return
	}

	var user GitHubUser
	// parses json response
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Username:", user.Login)
	fmt.Println("User ID:", user.ID)
	fmt.Println("Public Repos:", user.PublicRepos)
}
