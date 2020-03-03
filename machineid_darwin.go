// +build darwin

package machineid

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func MachineID() string {
	buf := &bytes.Buffer{}

	err := run(buf, os.Stderr, "ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	if err != nil {
		return ""
	}

	id, err := extractID(buf.String())
	if err != nil {
		return ""
	}

	return strings.TrimSpace(id)
}

func run(stdout, stderr io.Writer, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = stdout
	c.Stderr = stderr
	return c.Run()
}

func extractID(lines string) (string, error) {
	for _, line := range strings.Split(lines, "\n") {
		if strings.Contains(line, "IOPlatformUUID") {
			parts := strings.SplitAfter(line, `" = "`)
			if len(parts) == 2 {
				return strings.TrimRight(parts[1], `"`), nil
			}
		}
	}
	return "", fmt.Errorf("Failed to extract 'IOPlatformUUID' value from `ioreg` output.\n%s", lines)
}
