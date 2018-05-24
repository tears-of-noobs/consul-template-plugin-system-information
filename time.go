package main

import (
	"strconv"
	"time"
)

func obtainSystemTime(format string) (string, error) {
	currentTime := time.Now()

	switch format {
	case "ansic":
		return currentTime.Format(time.ANSIC), nil

	case "unix_date":
		return currentTime.Format(time.UnixDate), nil

	case "ruby_date":
		return currentTime.Format(time.RubyDate), nil

	case "rfc822":
		return currentTime.Format(time.RFC822), nil

	case "rfc822z":
		return currentTime.Format(time.RFC822Z), nil

	case "rfc850":
		return currentTime.Format(time.RFC850), nil

	case "rfc1123":
		return currentTime.Format(time.RFC1123), nil

	case "rfc1123z":
		return currentTime.Format(time.RFC1123Z), nil

	case "rfc3339":
		return currentTime.Format(time.RFC3339), nil

	case "rfc3339nano":
		return currentTime.Format(time.RFC3339Nano), nil

	default:
		return strconv.FormatInt(currentTime.Unix(), 10), nil
	}
}
