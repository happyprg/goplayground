package main

import (
	. "goplayground/pattern/gomock/util"
)

func main() {
	if err := FU.WriteMsgWithPath("./test_3.txt", "happyprg"); err != nil {
		panic(err)
	}
}
