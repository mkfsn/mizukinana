package concerts

import (
	"time"
)

var japan = time.FixedZone("Asia/Tokyo", 9*60*60)

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, japan)
}

func clock(hour, minute time.Duration) time.Duration {
	return time.Hour*hour + time.Minute*minute
}
