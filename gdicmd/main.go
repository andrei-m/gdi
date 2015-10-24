package main

import "github.com/andrei-m/gdi"
import "log"
import "os"

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("expected one argument: the package name to process. Got %d", len(os.Args))
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pkgs, err := gdi.GetDependencies(cwd, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for path := range pkgs {
		if matched, url := gdi.MatchAndGetGithubURL(path); matched {
			log.Println(url.String())
		}
	}
}
