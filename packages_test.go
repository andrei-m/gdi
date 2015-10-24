package gdi

import (
	"os"
	"testing"
)

func Test_GetDependencies(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	pkgs, err := GetDependencies(cwd, "github.com/andrei-m/gdi")
	if err != nil {
		t.Error(err)
	}
	if len(pkgs) != 1 {
		t.Fail()
	}
	_, ok := pkgs["github.com/andrei-m/gdi"]
	if !ok {
		t.Fail()
	}

	pkgs, err = GetDependencies(cwd, "github.com/andrei-m/gdi/gdicmd")
	if err != nil {
		t.Error(err)
	}
	if len(pkgs) != 2 {
		t.Fail()
	}
	_, ok = pkgs["github.com/andrei-m/gdi"]
	if !ok {
		t.Fail()
	}
	_, ok = pkgs["github.com/andrei-m/gdi/gdicmd"]
	if !ok {
		t.Fail()
	}
}
