package yoonotify

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthenticity(t *testing.T) {
	a, err := InitAuthenticity([]string{"172.17.0.0/16"})
	require.NoError(t, err)

	allowed, err := a.Allowed("172.17.0.2")
	require.NoError(t, err)
	require.True(t, allowed)

	allowed, err = a.Allowed("92.39.211.82")
	require.NoError(t, err)
	require.False(t, allowed)
}
