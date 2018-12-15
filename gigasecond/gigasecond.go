package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	gigasecondDuration, _ := time.ParseDuration("1000000000s")
	return t.Add(gigasecondDuration)
}
