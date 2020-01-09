package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/CCDirectLink/ccms/cmd/help"
	"github.com/CCDirectLink/ccms/internal/utils"
)

// New test
func New(wd string, pkg *utils.Package) string {
	if len(os.Args) < 3 {
		help.New()
		return ""
	}

	name := strings.Join(os.Args[2:], " ")

	name = utils.FormatPackageName(name)

	pkg.Name = name // test

	newWd := path.Join(wd, name)
	err := os.Mkdir(newWd, 0666)

	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("mod %s already exists", name)
			return ""
		}
		panic(err)
	}

	err = os.Mkdir(path.Join(newWd, "assets"), 0666)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	fmt.Printf("Successfully created mod %s", name)
	return newWd
}
