package main

import (
	my_os "goplayground/pattern/gomock/os"
)

func main() {
	var FU = FileUtil{Ow: my_os.NewOsWrapper(), Fw: my_os.NewFileWrapper()}

	if err := FU.WriteMsgWithPath("./20200114.txt", "hello world"); err != nil {
		panic(err)
	}
}
