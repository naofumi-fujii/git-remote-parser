package main

import "testing"

func Test_formatURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test", args{"test"}, "test"},
		{"test", args{"git@github.com:naofumi-fujii/git-browse-web"}, "ssh://git@github.com/naofumi-fujii/git-browse-web"},
		{"test", args{"ssh://git@gitlab.com/naofumi-fujii/gitlab.git"}, "ssh://git@gitlab.com/naofumi-fujii/gitlab"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFormattedGitRemoteURL(tt.args.url); got != tt.want {
				t.Errorf("getFormattedGitRemoteURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
