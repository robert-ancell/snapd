// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2022 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package builtin

const dbusControlSummary = `allows access to all of DBus`

const dbusControlBaseDeclarationSlots = `
  dbus-control:
    allow-installation:
      slot-snap-type:
        - core
    deny-auto-connection: true
`

const dbusControlConnectedPlugAppArmor = `
dbus (receive, send),
`

type dbusControlInterface struct {
	commonInterface
}

func init() {
	registerIface(&dbusControlInterface{commonInterface: commonInterface{
		name:                  "dbus-control",
		summary:               dbusControlSummary,
		implicitOnCore:        true,
		implicitOnClassic:     true,
		baseDeclarationSlots:  dbusControlBaseDeclarationSlots,
		connectedPlugAppArmor: dbusControlConnectedPlugAppArmor,
	}})
}
