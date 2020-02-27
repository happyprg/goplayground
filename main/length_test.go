package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	AttributeKeyName := "a"
	keys := map[string]bool{AttributeKeyName: true}

	searchKey := "a"
	require.True(t, keys[searchKey])
}
