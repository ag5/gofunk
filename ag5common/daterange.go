package ag5common

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
