/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fire833/k8ssm/cmd/k8ssm/app"
)

var (
	Version   string = "unknown"
	Commit    string = "unknown"
	BuildTime string = "unknown"
	Go        string = runtime.Version()
	Os        string = runtime.GOOS
	Arch      string = runtime.GOARCH

	VersionStr string = fmt.Sprintf("K8sSM - Version: %s\nBuilt on: %s\nCommit: %s\nGo version: %s\nOS: %s\nArch: %s\n\n",
		Version, BuildTime, Commit, Go, Os, Arch)
)

func main() {
	if e := app.NewK8sSMRunCommand().Execute(); e != nil {
		fmt.Printf("error with running server: %s", e.Error())
		os.Exit(1)
	}
}
