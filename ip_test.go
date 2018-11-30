package main

import (
	"regexp"
	"testing"
)

func TestGetIPAddress(t *testing.T) {

	availableVersions := map[string]*regexp.Regexp{
		"ipv4": regexp.MustCompile(
			`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)` +
				`(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`,
		),
		"ipv6": regexp.MustCompile(
			`(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:)` +
				`{1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|` +
				`([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|` +
				`([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|` +
				`([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|` +
				`([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|` +
				`[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:` +
				`[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4})` +
				`{0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}` +
				`((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}` +
				`(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]` +
				`{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])` +
				`\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`,
		),
	}

	var netdev = "lo"

	for testVersion, matchPattern := range availableVersions {
		t.Logf("testing version %s\n", testVersion)

		var IPv4Flag = false
		if testVersion == "ipv4" {
			IPv4Flag = true
		}

		ipaddress, err := getIP(netdev, IPv4Flag)
		if err != nil {
			t.Errorf(
				"unexpected error: %s\n", err.Error(),
			)
		}

		matchedString := matchPattern.FindString(ipaddress)

		if matchedString != ipaddress {
			t.Errorf(
				"testing failed for %s address: got value: %s, expected: %s\n",
				testVersion,
				matchedString,
				ipaddress,
			)
		}

	}
}
