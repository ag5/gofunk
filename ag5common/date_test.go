package ag5common_test

import (
	"github.com/ag5/gofunk/ag5common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateDate(t *testing.T) {

	date := ag5common.NewDate(2001, 01, 01)

	require.Equal(t, 2451911, date.Julian())
	require.Equal(t, "2001-01-01", date.String())
}

func TestInfiniteFuture(t *testing.T) {

	date := ag5common.InfiniteFuture

	require.Equal(t, 0x7FFFFFFF, date.Julian())
	require.Equal(t, "∞", date.String())
}

func TestInfinitePast(t *testing.T) {

	date := ag5common.InfinitePast

	require.Equal(t, -0x80000000, date.Julian())
	require.Equal(t, "-∞", date.String())
}
