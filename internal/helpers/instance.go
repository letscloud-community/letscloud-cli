package helpers

import (
	"regexp"
	"strings"
)

var (
	ipRegex, _ = regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
)

func IsIpv4Regex(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")
	return ipRegex.MatchString(ipAddress)
}

func GetInstanceStatus(Booted bool, Locked bool, Built bool, Suspended bool) string {
	status := "off"
	if Booted {
		status = "running"
	}
	if Locked {
		status = "action in progess"
	}
	if !Built {
		status = "building"
	}
	if Suspended {
		status = "suspended"
	}

	return status
}
