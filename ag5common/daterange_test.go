package ag5common_test

import (
	"github.com/ag5/gofunk/ag5common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateDateRange(t *testing.T) {

	start := ag5common.NewDate(2001, 01, 01)
	stop := ag5common.NewDate(2001, 12, 31)

	dr := ag5common.NewDateRange(start, stop)

	require.Equal(t, 365, dr.Len())
	require.Equal(t, "2001-01-01/2001-12-31", dr.String())
}
