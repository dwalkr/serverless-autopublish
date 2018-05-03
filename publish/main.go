package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type Response struct {
	Message string `json:"message"`
}

type PublishStatus struct {
	repo      string
	succeeded bool
}

func Handler() (Response, error) {
	repos := strings.Split(os.Getenv("repos"), ";")
	token := os.Getenv("github_token")
	authorName := os.Getenv("author_name")
	authorEmail := os.Getenv("author_email")
	return publishRepos(repos, token, authorName, authorEmail)
}

func publishRepos(repos []string, token string, authorName string, authorEmail string) (Response, error) {
	var err error
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	var succeeded []string
	var failed []string
	signature := &object.Signature{
		Name:  authorName,
		Email: authorEmail,
		When:  time.Now(),
	}

	wg.Add(len(repos))
	for _, repo := range repos {
		go func(repo string, token string) {
			defer wg.Done()
			err = publish(repo, token, signature)
			if err != nil {
				mutex.Lock()
				failed = append(failed, repo)
				mutex.Unlock()
			} else {
				mutex.Lock()
				succeeded = append(succeeded, repo)
				mutex.Unlock()
			}
		}(repo, token)
	}
	wg.Wait()
	message := ""
	if len(succeeded) > 0 {
		message = fmt.Sprintf("%d repos updated successfully. ", len(succeeded))
	}
	if len(failed) > 0 {
		message = fmt.Sprintf("%sThe following repos failed to update: %s", message, strings.Join(failed, ", "))
	}
	return Response{
		Message: message,
	}, nil
}

func publish(repo string, token string, signature *object.Signature) error {
	repoPath, err := ioutil.TempDir("", path.Base(repo))
	defer os.RemoveAll(repoPath)
	r, err := clone(repo, repoPath, token)
	if err != nil {
		return err
	}

	err = updatePublishDate(r, repoPath, signature)
	if err != nil {
		return err
	}
	return pushToOrigin(r, token)
}

func clone(repo string, repoPath string, token string) (*git.Repository, error) {
	r, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:               repo,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		return &git.Repository{}, err
	}
	return r, nil
}

func updatePublishDate(r *git.Repository, path string, signature *object.Signature) error {
	msg := fmt.Sprintf("Last published on %s", time.Now().Format(time.RFC3339))
	err := ioutil.WriteFile(filepath.Join(path, ".autopublish"), []byte(msg), 0644)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	_, err = w.Add(".autopublish")
	if err != nil {
		return err
	}
	_, err = w.Commit("Autopublish", &git.CommitOptions{
		Author: signature,
	})
	return err
}

func pushToOrigin(r *git.Repository, token string) error {
	return r.Push(&git.PushOptions{
		RemoteName: git.DefaultRemoteName,
		Auth: &http.BasicAuth{
			Username: token,
			Password: "",
		},
	})
}

func main() {
	lambda.Start(Handler)
}
