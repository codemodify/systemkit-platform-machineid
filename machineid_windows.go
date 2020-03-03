// +build windows

package machineid

import (
	"golang.org/x/sys/windows/registry"
)

// MachineID -
func MachineID() string {
	regKey, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		"SOFTWARE\\Microsoft\\Cryptography",
		registry.QUERY_VALUE|registry.WOW64_64KEY,
	)
	if err != nil {
		return ""
	}
	defer regKey.Close()

	machineGuid, _, err := regKey.GetStringValue("MachineGuid")
	if err != nil {
		return ""
	}

	return machineGuid
}
