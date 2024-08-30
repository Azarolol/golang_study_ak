package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-github/v53/github"
)

type MockRepoLister struct {
	wantErr bool
}

type MockGistLister struct {
	wantErr bool
}

func (m *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	if m.wantErr {
		return nil, nil, errors.New("test error")
	}
	gists := make([]github.Gist, 4)
	output := make([]*github.Gist, 4)
	for i := 0; i < 4; i++ {
		gists[i].ID = github.String(fmt.Sprintf("%d", i+1))
		gists[i].Description = github.String(fmt.Sprintf("Description of gist number %d", i+1))
		gists[i].HTMLURL = github.String(fmt.Sprintf("URL of gist number %d", i+1))
		output[i] = &gists[i]
	}

	return output, &github.Response{Response: &http.Response{Status: "200 OK"}}, nil
}

func (m *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	if m.wantErr {
		return nil, nil, errors.New("test error")
	}
	repos := make([]github.Repository, 4)
	output := make([]*github.Repository, 4)
	for i := 0; i < 4; i++ {
		repos[i].Name = github.String(fmt.Sprintf("Name of repository number %d", i+1))
		repos[i].Description = github.String(fmt.Sprintf("Description of repository number %d", i+1))
		repos[i].HTMLURL = github.String(fmt.Sprintf("URL of repository number %d", i+1))
		output[i] = &repos[i]
	}
	return output, &github.Response{Response: &http.Response{Status: "200 OK"}}, nil
}

func TestGithubAdapter_GetGists(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		g       GithubAdapter
		args    args
		want    []Item
		wantErr bool
	}{
		{"Test github adapter mock get gists", GithubAdapter{GistList: &MockGistLister{false}}, args{context.Background(), "username"}, []Item{
			{"Gist ID:1", "Description of gist number 1", "URL of gist number 1"},
			{"Gist ID:2", "Description of gist number 2", "URL of gist number 2"},
			{"Gist ID:3", "Description of gist number 3", "URL of gist number 3"},
			{"Gist ID:4", "Description of gist number 4", "URL of gist number 4"},
		}, false},
		{"Test github adapter mock get gists error", GithubAdapter{GistList: &MockGistLister{true}}, args{context.Background(), "username"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.GetGists(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GithubAdapter.GetGists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GithubAdapter.GetGists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubAdapter_GetRepos(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		g       GithubAdapter
		args    args
		want    []Item
		wantErr bool
	}{
		{"Test github adapter mock get repos", GithubAdapter{RepoList: &MockRepoLister{false}}, args{context.Background(), "username"}, []Item{
			{"Name of repository number 1", "Description of repository number 1", "URL of repository number 1"},
			{"Name of repository number 2", "Description of repository number 2", "URL of repository number 2"},
			{"Name of repository number 3", "Description of repository number 3", "URL of repository number 3"},
			{"Name of repository number 4", "Description of repository number 4", "URL of repository number 4"},
		}, false},
		{"Test github adapter mock get repos error", GithubAdapter{RepoList: &MockRepoLister{true}}, args{context.Background(), "username"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.GetRepos(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GithubAdapter.GetRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GithubAdapter.GetRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}
