package my_os

import "os"

type OsWrapper struct {
}

func NewOsWrapper() *OsWrapper {
	wrapper := &OsWrapper{}
	return wrapper
}
func (w *OsWrapper) Create(name string) (*os.File, error) {
	return os.Create(name)
}

type FileWrapper struct {
}

func NewFileWrapper() *FileWrapper {
	wrapper := &FileWrapper{}
	return wrapper
}

func (w *FileWrapper) WriteString(f *os.File, s string) (n int, err error) {
	return f.WriteString(s)
}
