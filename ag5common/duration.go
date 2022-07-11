package ag5common

import (
	"fmt"
)

type Duration struct {
	Years   int32
	Months  int32
	Days    int32
	Hours   int32
	Minutes int32
	Seconds int32
}

func (d Duration) String() string {
	return fmt.Sprintf("P%dY%dM%dDT%dH%dM%dS",
		d.Years,
		d.Months,
		d.Days,
		d.Hours,
		d.Minutes,
		d.Seconds,
	)
}

func (d Duration) HasSubDayParts() bool {
	return d.Hours != 0 || d.Minutes != 0 || d.Seconds != 0
}

func (d Duration) AddDuration(other Duration) Duration {
	return Duration{
		Years:   d.Years + other.Years,
		Months:  d.Months + other.Months,
		Days:    d.Days + other.Days,
		Hours:   d.Hours + other.Hours,
		Minutes: d.Minutes + other.Minutes,
		Seconds: d.Seconds + other.Seconds,
	}
}

func Years(years int) Duration {
	return Duration{
		Years: int32(years),
	}
}

func Months(months int) Duration {
	return Duration{
		Months: int32(months),
	}
}

func Days(days int) Duration {
	return Duration{
		Days: int32(days),
	}
}
