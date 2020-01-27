package util

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	system_os "os"
	"testing"
)

const (
	checkMark = "\u2713"
)

func TestFileWrite(t *testing.T) {
	path := "path"
	msg := "jeehong lee"

	t.Log("Create Failed", checkMark)
	{
		controller := gomock.NewController(t)
		defer controller.Finish()

		ow := NewMockOsInterface(controller)
		fw := NewMockFileInterface(controller)
		fu := FileUtil{ow, fw}
		var f system_os.File
		expectedErr := errors.New("file creation failed because of runtime exception")
		ow.EXPECT().Create(path).Return(&f, expectedErr).Times(1)
		fw.EXPECT().WriteString(gomock.Any(), gomock.Any()).Times(0)

		actualErr := fu.WriteMsgWithPath(path, msg)
		require.Equal(t, expectedErr, actualErr)
	}

	t.Log("Create,Write Succeed", checkMark)
	{
		controller := gomock.NewController(t)
		defer controller.Finish()

		ow := NewMockOsInterface(controller)
		fw := NewMockFileInterface(controller)
		fu := FileUtil{ow, fw}
		var f system_os.File
		var expectedErr error = nil
		ow.EXPECT().Create(path).Return(&f, expectedErr).Times(1)
		fw.EXPECT().WriteString(&f, msg).Times(1)

		actualErr := fu.WriteMsgWithPath(path, msg)
		require.Equal(t, expectedErr, actualErr)
	}
}
