package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "access_token"})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	g := NewGithub(client)

	fmt.Println(g.GetGists(context.Background(), "Azarolol"))
	fmt.Println(g.GetRepos(context.Background(), "Azarolol"))
}

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error) // opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
}

type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

type GithubProxy struct {
	github Githuber
	cache  map[string][]Item
}

func NewGithub(githubClient *github.Client) *GithubProxy {
	return &GithubProxy{
		github: NewGithubAdapter(githubClient),
		cache:  make(map[string][]Item),
	}
}

func NewGithubAdapter(githubClient *github.Client) *GithubAdapter {
	g := &GithubAdapter{
		RepoList: githubClient.Repositories,
		GistList: githubClient.Gists,
	}

	return g
}

func (g GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	opt := &github.GistListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	gists, resp, err := g.GistList.List(ctx, username, opt)
	if err != nil {
		return nil, err
	}

	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}

	items := make([]Item, len(gists))
	for i := 0; i < len(gists); i++ {
		items[i].Title = "Gist ID:" + gists[i].GetID()
		items[i].Description = gists[i].GetDescription()
		items[i].Link = gists[i].GetHTMLURL()
	}

	return items, nil
}

func (g GithubProxy) GetGists(ctx context.Context, username string) ([]Item, error) {
	key := "gists " + username
	if value, ok := g.cache[key]; ok {
		return value, nil
	}

	value, err := g.github.GetGists(ctx, username)
	if err != nil {
		return nil, err
	}

	g.cache[key] = value

	return value, nil
}

func (g GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	repos, resp, err := g.RepoList.List(ctx, username, opt)
	if err != nil {
		return nil, err
	}

	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}

	items := make([]Item, len(repos))
	for i := 0; i < len(repos); i++ {
		items[i].Title = repos[i].GetName()
		items[i].Description = repos[i].GetDescription()
		items[i].Link = repos[i].GetHTMLURL()
	}

	return items, nil
}

func (g GithubProxy) GetRepos(ctx context.Context, username string) ([]Item, error) {
	key := "repos " + username
	if value, ok := g.cache[key]; ok {
		return value, nil
	}

	value, err := g.github.GetRepos(ctx, username)
	if err != nil {
		return nil, err
	}

	g.cache[key] = value

	return value, nil
}

type Item struct {
	Title       string
	Description string
	Link        string
}
