package gdi

import (
	"testing"
)

func Test_MatchAndGetGithubURL(t *testing.T) {
	matched, _ := MatchAndGetGithubURL("foo")
	if matched {
		t.Error("expected no match")
	}

	matched, u := MatchAndGetGithubURL("github.com/foo/bar")
	if !matched {
		t.Error("expected match")
	}
	if u.Scheme != "https" {
		t.Errorf("unexpected URL scheme: %v", u.Scheme)
	}
	if u.Host != "github.com" {
		t.Errorf("unexpected URL host: %v", u.Host)
	}
	if u.Path != "foo/bar" {
		t.Errorf("unexpected path: %v", u.Path)
	}

	matched, u = MatchAndGetGithubURL("_vendor/github.com/foo/bar")
	if !matched {
		t.Error("expected match")
	}
	if u.Path != "foo/bar" {
		t.Errorf("unexpected path: %v", u.Path)
	}
}
