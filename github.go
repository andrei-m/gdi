package gdi

import (
	"net/url"
	"regexp"
)

var githubPkgPath = regexp.MustCompile(`.*?github\.com/([A-Za-z0-9_-]+/[A-Za-z0-9_-]+).*`)

func MatchAndGetGithubURL(pkgName string) (bool, *url.URL) {
	if matches := githubPkgPath.FindStringSubmatch(pkgName); matches != nil {
		u := url.URL{
			Scheme: "https",
			Host:   "github.com",
			Path:   matches[1],
		}
		return true, &u
	}
	return false, nil
}
