package main

import (
	"regexp"
	"testing"
)

func TestObtainSystemTime(t *testing.T) {
	availableFormats := map[string]*regexp.Regexp{
		"ansic": regexp.MustCompile(
			`\w{2,3} \w{3} [\s]?\d{1,2} \d{2}:\d{2}:\d{2} \d{4}`,
		),

		"unix_date": regexp.MustCompile(
			`\w{2,3} \w{3} [\s]?\d{1,2} \d{2}:\d{2}:\d{2} [\-|\+]?\d{2} \d{4}`,
		),

		"ruby_date": regexp.MustCompile(
			`\w{3} \w{3} \d{2} \d{2}:\d{2}:\d{2} [\-|\+]?\d{1,4} \d{4}`,
		),

		"rfc822": regexp.MustCompile(
			`\d{2} \w{3} \d{2} \d{2}:\d{2} [\-|\+]?\d{2}`,
		),

		"rfc822z": regexp.MustCompile(
			`\d{2} \w{3} \d{2} \d{2}:\d{2} [\-|\+]?\d{1,4}`,
		),

		"rfc850": regexp.MustCompile(
			`\w{6,9}, \d{2}-\w{3}-\d{2} \d{2}:\d{2}:\d{2} [\-|\+]?\d{1,4}`,
		),

		"rfc1123": regexp.MustCompile(
			`\w{3}, \d{2} \w{3} \d{4} \d{2}:\d{2}:\d{2} [\-|\+]?\d{2}`,
		),

		"rfc1123z": regexp.MustCompile(
			`\w{3}, \d{2} \w{3} \d{4} \d{2}:\d{2}:\d{2} [\-|\+]?\d{1,4}`,
		),

		"rfc3339": regexp.MustCompile(
			`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[\-|\+]?\d{2}:\d{2}`,
		),

		"rfc3339nano": regexp.MustCompile(
			`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{9}[\-|\+]?\d{2}:\d{2}`,
		),

		"empty": regexp.MustCompile(
			`\d{10}`,
		),
	}

	for format, matchPattern := range availableFormats {
		t.Logf("testing format %s\n", format)

		systemTime, err := obtainSystemTime(format)
		if err != nil {
			t.Errorf(
				"unexpected error: %s\n", err.Error(),
			)
		}
		matchedString := matchPattern.FindString(systemTime)

		if matchedString != systemTime {
			t.Errorf(
				"testing failed for format %s: got value: %s, expected: %s\n",
				format,
				matchedString,
				systemTime,
			)
		}

	}
}
