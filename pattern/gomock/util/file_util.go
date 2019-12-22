package util

import (
	my_os "goplayground/pattern/gomock/os"
	system_os "os"
)

type OsInterface interface {
	Create(name string) (*system_os.File, error)
}

type FileInterface interface {
	WriteString(f *system_os.File, s string) (n int, err error)
}

var FU = FileUtil{ow: my_os.NewOsWrapper(), fw: my_os.NewFileWrapper()}

type FileUtil struct {
	ow OsInterface
	fw FileInterface
}

func (f FileUtil) WriteMsgWithPath(path, msg string) (err error) {
	file, err := f.ow.Create(path)
	defer file.Close()
	if err != nil {
		return
	}
	_, err = f.fw.WriteString(file, msg)
	return
}
