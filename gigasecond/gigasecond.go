package gigasecond

import "time"

const gigasecond = 1e9 * time.Second

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
