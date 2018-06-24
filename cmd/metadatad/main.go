// Metadata is an application acting as Metadata server.
// Copyright (C) 2017 Kassisol inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/user"
	"github.com/kassisol/metadata/daemon"
)

func main() {
	u := user.New()

	if !u.IsRoot() {
		log.Fatal("You must be root to run that command")
	}

	cmd := daemon.NewCommand()

	cmd.SetHelpTemplate(helpTemplate)
	cmd.SetUsageTemplate(usageTemplate)

	if err := cmd.Execute(); err != nil {
		utils.Exit(err)
	}
}

var usageTemplate = `{{ .Short | trim }}

Usage:{{ if .Runnable }}
  {{ if .HasAvailableFlags }}{{ appendIfNotPresent .UseLine "[flags]" }}{{ else }}{{ .UseLine }}{{ end }}{{ end }}{{ if .HasAvailableSubCommands }}
  {{ .CommandPath }} [command]{{ end }}{{ if gt .Aliases 0 }}

Aliases:
  {{ .NameAndAliases }}{{ end }}{{ if .HasExample }}

Examples:
  {{ .Example }}{{ end }}{{ if .HasAvailableSubCommands }}

Available Commands:{{ range .Commands }}{{ if .IsAvailableCommand }}
  {{ rpad .Name .NamePadding }} {{ .Short }}{{ end }}{{ end }}{{ end }}{{ if .HasAvailableLocalFlags }}

Flags:
{{ .LocalFlags.FlagUsages | trimRightSpace }}{{ end }}{{ if .HasAvailableInheritedFlags }}

Global Flags:
  {{ .InheritedFlags.FlagUsages | trimRightSpace }}{{ end }}{{ if .HasHelpSubCommands }}

Additional help topics:{{ range .Commands }}{{ if .IsHelpCommand }}
  {{ rpad .CommandPath .CommandPathPadding }} {{ .Short }}{{ end }}{{ end }}{{ end }}{{ if .HasAvailableSubCommands }}

Use "{{ .CommandPath }} [command] --help" for more information about a command.{{ end }}
`

var helpTemplate = `
{{ if or .Runnable .HasSubCommands }}{{ .UsageString }}{{ end }}`
