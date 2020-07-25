// github.com/6r6
package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&Form1)
	vcl.Application.Run()
}
