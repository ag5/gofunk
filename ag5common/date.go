package ag5common

import (
	"fmt"
	"time"
)

// A Date is internally stored as a Julian Date int
type Date int32

const (
	InfiniteFuture Date = 0x7FFFFFFF
	InfinitePast   Date = -0x80000000
)

func NewDate(y int32, m int32, d int32) Date {
	// https://pdc.ro.nu/jd-code.html
	y += 8000
	if m < 3 {
		y--
		m += 12
	}
	dateInt := (y * 365) + (y / 4) - (y / 100) + (y / 400) - 1200820 + (m*153+3)/5 - 92 + d - 1
	return Date(dateInt)
}

func (date Date) AsYMD() (y int32, m int32, d int32) {
	jd := int32(date)
	l := jd + 68569
	n := 4 * l / 146097
	l = l - ((146097*n + 3) / 4)
	i := 4000 * (l + 1) / 1461001
	l = l - ((1461 * i) / 4) + 31
	j := 80 * l / 2447
	d = l - ((2447 * j) / 80)
	l = j / 11
	m = j + 2 - (12 * l)
	y = 100*(n-49) + i + l
	return y, m, d
}

func (date Date) Time() time.Time {
	y, m, d := date.AsYMD()
	return time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.UTC)
}

func (date Date) String() string {
	if date == InfiniteFuture {
		return "∞"
	}
	if date == InfinitePast {
		return "-∞"
	}
	y, m, d := date.AsYMD()
	return fmt.Sprintf("%04d-%02d-%02d", y, m, d)
}

func (date Date) Julian() int {
	return int(date)
}

func (date Date) AddDuration(dur Duration) Date {
	if dur.HasSubDayParts() {
		panic(fmt.Errorf("duration has sub-day parts"))
	}
	return date.AddYears(dur.Years).AddWholeMonths(dur.Months).AddDays(dur.Days)
}

func (date Date) AddYears(years int32) Date {
	y, m, d := date.AsYMD()
	return NewDate(y+years, m, d)
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

var daysInMonth = [12]int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func leapYear(year int32) int32 {
	if (year%4 != 0) || ((year%100 == 0) && (year%400 != 0)) {
		return 0
	}
	return 1
}
func DaysInMonth(year int32, month int32) int32 {
	dim := daysInMonth[month-1]
	if month == 2 {
		return dim + leapYear(year)
	}
	return dim
}

func (date Date) AddWholeMonths(months int32) Date {
	y, m, d := date.AsYMD()
	m = m + months
	y = y + (m / 12)
	m = m % 12

	dayIndex := min(d, DaysInMonth(y, m))

	return NewDate(y, m, dayIndex)
}

func (date Date) AddDays(days int32) Date {
	y, m, d := date.AsYMD()
	return NewDate(y, m, d+days)

}

func (date Date) Year() int32 {
	y, _, _ := date.AsYMD()
	return y
}

func (date Date) Month() int32 {
	_, m, _ := date.AsYMD()
	return m
}

func (date Date) DayOfMonth() int32 {
	_, _, d := date.AsYMD()
	return d
}
