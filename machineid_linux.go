// +build linux
// +build !android

package machineid

import (
	"io/ioutil"
	"strings"
)

// MachineID -
func MachineID() string {
	knownFiles := []string{
		"/var/lib/dbus/machine-id",
		"/etc/machine-id",
	}

	for _, knownFile := range knownFiles {
		data, err := ioutil.ReadFile(knownFile)
		if err == nil {
			return strings.TrimSpace(string(data))
		}
	}

	return ""
}
