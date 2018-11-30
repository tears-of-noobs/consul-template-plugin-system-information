package main

import (
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/reconquest/hierr-go"
)

var version = "[manual]"

func main() {
	usage := `consul-template-plugin-system-information

Usage:
    consul-template-plugin-system-information hostname
    consul-template-plugin-system-information time [<format>]
    consul-template-plugin-system-information ipv4 [<netdev>]
    consul-template-plugin-system-information ipv6 [<netdev>]

Options:
    <format>  Format of return time value. Available formats
               ansic, unix_date, ruby_date, rfc822, rfc822z,
               rfc850, rfc1123, rfc1123z, rfc3339, rfc3339nano.
               If format not defined plugin will return time in 
               UNIX timestamp format.

	<netdev>  Return first IPv4 or IPv6 address from network
               device [default: eth0].
`

	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		panic(err)
	}

	var (
		statusCode  = 0
		emptyFormat = "empty"
		netdev      = "eth0"
	)

	defer func() {
		os.Exit(statusCode)
	}()

	switch true {
	case args["hostname"].(bool):
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println(
				hierr.Errorf(
					err,
					"can't obtain hostname",
				).Error(),
			)

			statusCode = 1
			return
		}

		fmt.Printf("%s", hostname)
		return

	case args["time"].(bool):
		timeFormat := emptyFormat

		if args["<format>"] != nil {
			timeFormat = args["<format>"].(string)
		}

		systemTime, err := obtainSystemTime(timeFormat)
		if err != nil {
			fmt.Println(
				hierr.Errorf(
					err,
					"can't obtain system time",
				).Error(),
			)

			statusCode = 1
			return
		}

		fmt.Printf("%s", systemTime)
		return

	case args["ipv4"].(bool):

		var IPv4Flag = true

		if args["<netdev>"] != nil {
			netdev = args["<netdev>"].(string)
		}

		IPv4, err := getIP(netdev, IPv4Flag)
		if err != nil {
			fmt.Println(
				hierr.Errorf(
					err,
					"can't get IPv4 address for interface %s",
					netdev,
				).Error(),
			)

			statusCode = 1
			return
		}

		fmt.Printf("%s", IPv4)
		return

	case args["ipv6"].(bool):

		var IPv4Flag = false

		if args["<netdev>"] != nil {
			netdev = args["<netdev>"].(string)
		}

		IPv6, err := getIP(netdev, IPv4Flag)
		if err != nil {
			fmt.Println(
				hierr.Errorf(
					err,
					"can't get IPv6 address for interface %s",
					netdev,
				).Error(),
			)

			statusCode = 1
			return
		}

		fmt.Printf("%s", IPv6)
		return

	default:
		statusCode = 1

		fmt.Println("unexpected behavior")
	}
}
