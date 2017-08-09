package tryInternal

import (
	"fmt"
	"practice/mass/tryInternal/internal"
	"practice/mass/tryInternal/internal/submodule"
)

func Print() {
	fmt.Println("this is tryInternal module")
}

func InternalPrint() {
	fmt.Println("call internal.Print() at parent module")
	internal.Print()
}

func CallSubmodule() {
	fmt.Println("call internal/submodule.Print() at internal parent module")
	submodule.Print()
}
