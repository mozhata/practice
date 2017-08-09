package sibsub

import (
	"fmt"
	"practice/mass/tryInternal/internal"
	"practice/mass/tryInternal/internal/submodule"
)

func CallInternalFunc() {
	fmt.Println("call internal.Print() at internal's sibling's submodule")
	internal.Print()
}

func CallInternalSubmodul() {
	fmt.Println("call internal/submodule at internal's sibling's submodule")
	submodule.Print()
}
