package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type MockRepoLister struct {
	wantErr bool
}

type MockGistLister struct {
	wantErr bool
}

func (m *MockGistLister) GetItems(ctx context.Context, username string) ([]Item, error) {
	if m.wantErr {
		return nil, errors.New("test error")
	}
	gists := make([]Item, 4)
	for i := 0; i < 4; i++ {
		gists[i].Title = fmt.Sprintf("%d", i+1)
		gists[i].Description = fmt.Sprintf("Description of gist number %d", i+1)
		gists[i].Link = fmt.Sprintf("URL of gist number %d", i+1)
	}

	return gists, nil
}

func (m *MockRepoLister) GetItems(ctx context.Context, username string) ([]Item, error) {
	if m.wantErr {
		return nil, errors.New("test error")
	}
	repos := make([]Item, 4)
	for i := 0; i < 4; i++ {
		repos[i].Title = fmt.Sprintf("%d", i+1)
		repos[i].Description = fmt.Sprintf("Description of repository number %d", i+1)
		repos[i].Link = fmt.Sprintf("URL of repository number %d", i+1)
	}

	return repos, nil
}

func TestGeneralGithub_GetItems(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		strategy GithubLister
	}
	tests := []struct {
		name    string
		g       *GeneralGithub
		args    args
		want    []Item
		wantErr bool
	}{
		{"Test get items gists mock", &GeneralGithub{}, args{nil, "", &MockGistLister{false}}, []Item{
			{"1", "Description of gist number 1", "URL of gist number 1"},
			{"2", "Description of gist number 2", "URL of gist number 2"},
			{"3", "Description of gist number 3", "URL of gist number 3"},
			{"4", "Description of gist number 4", "URL of gist number 4"},
		}, false},
		{"Test get items gists mock err", &GeneralGithub{}, args{nil, "", &MockGistLister{true}}, nil, true},
		{"Test get items repos mock", &GeneralGithub{}, args{nil, "", &MockRepoLister{false}}, []Item{
			{"1", "Description of repository number 1", "URL of repository number 1"},
			{"2", "Description of repository number 2", "URL of repository number 2"},
			{"3", "Description of repository number 3", "URL of repository number 3"},
			{"4", "Description of repository number 4", "URL of repository number 4"},
		}, false},
		{"Test get items repos mock err", &GeneralGithub{}, args{nil, "", &MockRepoLister{true}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.GetItems(tt.args.ctx, tt.args.username, tt.args.strategy)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneralGithub.GetItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneralGithub.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
