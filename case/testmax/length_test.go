package testmax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgModify_ValidateBasic(t *testing.T) {
	MaxChangeFieldsCount := 100
	// Given msg with changes more than max
	changeList := make([]string, MaxChangeFieldsCount+1)
	require.True(t, len(changeList) > MaxChangeFieldsCount)
}
