package windows

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// V1809 is the Windows version that is 1809 and also known as Windows 2019
	// ltsc.
	V1809 = "1809"
	// V2004 is the Windows version that is 2004 sac.
	V2004 = "2004"
	// V20H2 is the Windows version that is 2009 sac.
	V20H2 = "2009"
	// V21H1 is the Windows version that is 21H1 sac.
	V21H1 = "21H1"
	// V21H2 is the Windows version that is 21H2 also known as Windows 2022 LTSC.
	V21H2 = "21H2"
)

var ErrUnsupportedWindowsVersion = errors.New("unsupported windows version")

var supportedWindowsBuilds = map[string]string{
	// Windows server versions: https://en.wikipedia.org/wiki/List_of_Microsoft_Windows_versions#Server_versions
	// Compatibility: https://learn.microsoft.com/en-us/virtualization/windowscontainers/deploy-containers/version-compatibility#windows-server-host-os-compatibility
	"10.0.17763": V1809,
	"10.0.19041": V2004,
	"10.0.19042": V20H2,
	"10.0.20348": V21H2,

	// Windows client versions: https://en.wikipedia.org/wiki/List_of_Microsoft_Windows_versions#Personal_computer_versions
	// Compatibility: https://learn.microsoft.com/en-us/virtualization/windowscontainers/deploy-containers/version-compatibility#windows-client-host-os-compatibility
	"10.0.19043": V1809,
	"10.0.19044": V1809,
	"10.0.19045": V1809,
	"10.0.22000": V21H2,
	"10.0.22621": V21H2,
}

// Version checks the specified kernel version to see if it's one of the
// supported Windows version. If true, it returns the os version.
// UnsupportedWindowsVersionError is returned when no supported Windows version
// is found in the string.
func Version(version string) (string, error) {
	semver := strings.FieldsFunc(version, func(r rune) bool {
		return r == '.' || r == ' '
	})
	if len(semver) < 3 {
		return "", fmt.Errorf("%w: %v", ErrUnsupportedWindowsVersion, version)
	}

	build := strings.Join(semver[:3], ".")
	windowsVersion, ok := supportedWindowsBuilds[build]
	if ok {
		return windowsVersion, nil
	}

	return "", fmt.Errorf("%w: %v", ErrUnsupportedWindowsVersion, version)
}
