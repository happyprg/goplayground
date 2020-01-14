package util

import (
	system_os "os"
)

type OsInterface interface {
	Create(name string) (*system_os.File, error)
}

type FileInterface interface {
	WriteString(f *system_os.File, s string) (n int, err error)
}

type FileUtil struct {
	Ow OsInterface
	Fw FileInterface
}

func (f FileUtil) WriteMsgWithPath(path, msg string) (err error) {
	file, err := f.Ow.Create(path)
	defer file.Close()
	if err != nil {
		return
	}
	_, err = f.Fw.WriteString(file, msg)
	return
}
