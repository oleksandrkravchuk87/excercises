package gigasecond

import (
	"time"
	"math"
)

func AddGigasecond(in time.Time) time.Time {
	return in.Add(time.Second*time.Duration(math.Pow(10, 9)))
}