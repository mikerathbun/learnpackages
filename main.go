package main

import (
	"fmt"

	"github.com/mikerathbun/learnpackages/packageutil"
)

func main() {
	a := "mike"

	packageutil.PrintMe(a)

	b := packageutil.Person{}

	b.Name = "Mark"
	b.Sex = "Male"
	fmt.Println(b.String())

}
