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

	qt "github.com/mappu/miqt/qt6"
	"github.com/mark-summerfield/cardindex/model"
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
	distro := ""
	if out, err := exec.Command("lsb_release",
		"-ds").Output(); err == nil {
		distro = strings.TrimSpace(string(out))
	}
	var utsname syscall.Utsname
	_ = syscall.Uname(&utsname)
	sysname := int8ToStr(utsname.Sysname[:])
	cpu := int8ToStr(utsname.Machine[:])
	release := int8ToStr(utsname.Release[:])
	sqlite_version, _ := model.SqliteVersion()
	// TODO use qVersion() for Qt
	return fmt.Sprintf(
		`<h3 align=center><font color=navy>%s v%s</font></h3>
<p align=center><font color=navy>A software implementation<br>of a
card index system.</font></p>
<p align=center>
<a
href="https://github.com/mark-summerfield/cardindex">https://github.com/mark-summerfield/cardindex</a>
</p>
<p align=center>
<font color=green>Copyright © %s Mark Summerfield.<br>
All rights reserved.<br>
License: GPLv3.
</p>
<p align=center><font color=#222>%s (%s/%s)</font><br>
<font color=#222>Qt 6 • SQLite %s</font><br>
<font color=#222>%s</font><br>
<font color=#222>%s-%s/%s</font><br>
</p>`,
		APPNAME, Version, year, runtime.Version(), runtime.GOOS,
		runtime.GOARCH, sqlite_version, distro, sysname, release, cpu)
}
