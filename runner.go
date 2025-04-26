package main

import (
	"fmt"
	"github.com/nik-git/nikgo/util"
	"github.com/nik-git/nikgo/util/pkg1"
)

func main() {
	res := util.Sub(10, 3)
	pkg1.Abc()
	pkg1.BBB()
	fmt.Println(res)
}
