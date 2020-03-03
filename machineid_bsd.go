// +build freebsd netbsd openbsd dragonfly solaris

package machineid

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// machineID returns the uuid specified at `/etc/hostid`.
// If the returned value is empty, the uuid from a call to `kenv -q smbios.system.uuid` is returned.
// If there is an error an empty string is returned.

// MachineID -
func MachineID() string {
	id, err := readFromFiles()
	if err == nil {
		return id
	}

	id, err = readKenv()
	if err == nil {
		return id
	}

	return ""
}

func readFromFiles() (string, error) {
	knownFiles := []string{
		"/etc/hostid",
	}

	for _, knownFile := range knownFiles {
		data, err := ioutil.ReadFile(filename)
		if err == nil {
			return strings.TrimSpace(string(data)), nil
		}
	}

	return "", fmt.Errorf("")
}

func readKenv() (string, error) {
	buf := &bytes.Buffer{}
	err := run(buf, os.Stderr, "kenv", "-q", "smbios.system.uuid")
	if err != nil {
		return "", err
	}
	return trim(buf.String()), nil
}

func run(stdout, stderr io.Writer, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = stdout
	c.Stderr = stderr
	return c.Run()
}
