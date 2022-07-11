package ag5common_test

import (
	"github.com/ag5/gofunk/ag5common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateDuration(t *testing.T) {

	dur := ag5common.Years(1)

	require.Equal(t, "P1Y0M0DT0H0M0S", dur.String())
}

func TestAddMonths(t *testing.T) {

	d := ag5common.NewDate(2001, 01, 31)
	dur := ag5common.Months(1)

	oneMonthLater := d.AddDuration(dur)

	require.Equal(t, ag5common.NewDate(2001, 02, 28), oneMonthLater)
}

func TestAddDurations(t *testing.T) {

	dur := ag5common.Months(1).AddDuration(ag5common.Years(1))
	d := ag5common.NewDate(2001, 01, 31)

	oneYearAndMonthLater := d.AddDuration(dur)

	require.Equal(t, ag5common.NewDate(2002, 02, 28), oneYearAndMonthLater)
}
