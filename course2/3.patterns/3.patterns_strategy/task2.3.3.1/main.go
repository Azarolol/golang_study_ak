package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

type Item struct {
	Title       string
	Description string
	Link        string
}

type GithubGist struct {
	gists *github.GistsService
}
type GithubRepo struct {
	repos *github.RepositoriesService
}
type GeneralGithub struct {
	client *github.Client
}

func (g *GithubGist) GetItems(ctx context.Context, username string) ([]Item, error) {
	opt := &github.GistListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	gists, resp, err := g.gists.List(ctx, username, opt)
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

func (g *GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error) {
	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	gists, resp, err := g.repos.List(ctx, username, opt)
	if err != nil {
		return nil, err
	}

	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}

	items := make([]Item, len(gists))
	for i := 0; i < len(gists); i++ {
		items[i].Title = gists[i].GetName()
		items[i].Description = gists[i].GetDescription()
		items[i].Link = gists[i].GetHTMLURL()
	}

	return items, nil
}

func NewGithubGist(client *github.Client) *GithubGist {
	return &GithubGist{gists: client.Gists}
}

func NewGithubRepo(client *github.Client) *GithubRepo {
	return &GithubRepo{repos: client.Repositories}
}

type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

func NewGeneralGithub(client *github.Client) *GeneralGithub {
	return &GeneralGithub{client: client}
}

func (g *GeneralGithub) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error) {
	return strategy.GetItems(ctx, username)
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "access_token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)

	gg := NewGeneralGithub(client)

	data, err := gg.GetItems(context.Background(), "Azarolol", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)

	data, err = gg.GetItems(context.Background(), "Azarolol", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
