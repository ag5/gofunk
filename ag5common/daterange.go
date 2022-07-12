package ag5common

import (
	"github.com/jackc/pgtype"
)

type DateRange struct {
	First Date
	Last  Date
}

func (r DateRange) Len() int {
	return int((r.Last - r.First) + 1)
}

func (r DateRange) String() interface{} {
	return r.First.String() + "/" + r.Last.String()
}

func NewDateRange(first Date, last Date) DateRange {
	return DateRange{
		First: first,
		Last:  last,
	}
}

func (dst *DateRange) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var pgval pgtype.Daterange
	err := pgval.Scan(src)
	if err != nil {
		return err
	}
	lTime := pgval.Lower.Time
	uTime := pgval.Upper.Time
	lVal := NewDateFromTime(lTime)
	uVal := NewDateFromTime(uTime)
	dr := NewDateRange(lVal, uVal)
	dst = &dr
	return nil
}
