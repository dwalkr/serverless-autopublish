package main

import "testing"

func TestClone(t *testing.T) {
	repos := []string{"url-to-your-repo", "url-to-another-repo"}
	token := "your-github-oauth-token"
	publishRepos(repos, token, "your-author-name", "your-email")
}
