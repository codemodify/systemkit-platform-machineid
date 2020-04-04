# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) CPU Detection at Runtime
[![](https://img.shields.io/github/v/release/codemodify/systemkit-platform-machineid?style=flat-square)](https://github.com/codemodify/systemkit-platform-machineid/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-platform-machineid?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-platform-machineid?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-platform-machineid/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-platform-machineid?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-platform-machineid?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-platform-machineid)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-platform-machineid)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-platform-machineid?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-platform-machineid?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-platform-machineid?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-platform-machineid?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-platform-machineid?style=flat-square)

#### Detailed CPU Detection at Runtime
#### Supported: Linux, Raspberry Pi, FreeBSD, Mac OS, Windows, Solaris

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-platform-machineid
```
# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) API

&nbsp;																| &nbsp;
---     															| ---
MachineID() | MachineID returns the uuid specified at `/etc/hostid`
readFromFiles() | 
readKenv() | 
run()  | 
extractID() | 

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage
```go
package main

import mchineid "github.com/codemodify/systemkit-platform-machineid"

func main() {
	machineid "github.com/codemodify/systemkit-platform-machineid"
}
```

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Dox
- http://0pointer.de/blog/projects/ids.html
