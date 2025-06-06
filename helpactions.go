// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/mappu/miqt/qt"
	"github.com/mark-summerfield/cardindex/database"
)

var (
	distro      string
	lsb_release bool
)

func (me *App) helpHelp() {
	fmt.Println("helpHelp") // TODO
}

func (me *App) helpAbout() {
	qt.QMessageBox_About(me.window.QWidget, "About — "+APPNAME, aboutHtml())
}

func aboutHtml() string {
	var year string
	y := time.Now().Year()
	if y == 2025 {
		year = fmt.Sprintf("%d", y)
	} else {
		year = fmt.Sprintf("2025-%d", y-2000)
	}
	if !lsb_release {
		if raw, err := exec.Command("lsb_release",
			"-ds").Output(); err == nil {
			distro = fmt.Sprintf("<font color=#222>%s</font><br>",
				strings.TrimSpace(string(raw)))
		}
		lsb_release = true
	}
	var utsname syscall.Utsname
	_ = syscall.Uname(&utsname)
	sysname := int8ToStr(utsname.Sysname[:])
	cpu := int8ToStr(utsname.Machine[:])
	release := int8ToStr(utsname.Release[:])
	sqlite_version, _ := database.SqliteVersion()
	qt_version := qt.QLibraryInfo_Version().ToString()
	return fmt.Sprintf(
		`<h3 align=center><font color=navy>%s v%s</font></h3>
<p align=center><font color=navy>A software implementation<br>of a
card index system.</font></p>
<p align=center>
<a
href="https://github.com/mark-summerfield/cardindex">https://github.com/mark-summerfield/cardindex</a>
</p>
<p align=center>
<font color=green>Copyright © %s<br>Mark Summerfield.<br>
All rights reserved.<br>
License: GPLv3.
</p>
<p align=center><font color=#222>%s (%s/%s)</font><br>
<font color=#222>miQt %s • Qt %s</font><br>
<font color=#222>SQLite %s</font><br>
%s
<font color=#222>%s-%s/%s</font><br>
</p>`,
		APPNAME, Version, year, runtime.Version(), runtime.GOOS,
		runtime.GOARCH, miqtVersion(), qt_version, sqlite_version, distro,
		sysname, release, cpu)
}
