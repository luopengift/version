package main

import (
	"fmt"

	"github.com/luopengift/version"
)

func main() {
	fmt.Println(version.VERSION, version.TIME, version.APP, version.GIT)
}
