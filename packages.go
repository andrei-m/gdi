package gdi

import "go/build"

func GetDependencies(root, pkgName string) (map[string]*build.Package, error) {
	pkgs := map[string]*build.Package{}

	if err := doGetDependencies(pkgs, root, pkgName); err != nil {
		return nil, err
	}

	return pkgs, nil
}

func doGetDependencies(pkgs map[string]*build.Package, root, pkgName string) error {
	pkg, err := build.Default.Import(pkgName, root, 0)
	if err != nil {
		return err
	}
	// skip packages in the root of the standard library
	if pkg.Goroot {
		return nil
	}
	pkgs[pkg.ImportPath] = pkg

	for _, imp := range pkg.Imports {
		if _, ok := pkgs[imp]; !ok {
			if err := doGetDependencies(pkgs, root, imp); err != nil {
				return err
			}
		}
	}
	return nil
}
