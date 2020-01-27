package main

import (
	my_os "github.com/happyprg/goplayground/gomock/os"
	. "github.com/happyprg/goplayground/gomock/util"
)

func main() {
	var FU = FileUtil{Ow: my_os.NewOsWrapper(), Fw: my_os.NewFileWrapper()}

	if err := FU.WriteMsgWithPath("./20200114.txt", "hello world"); err != nil {
		panic(err)
	}
}
