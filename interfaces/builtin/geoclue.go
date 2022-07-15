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

const geoclueSummary = `allows access to the GeoClue geolocation service`

const geoclueBaseDeclarationSlots = `
  geoclue:
    allow-installation:
      slot-snap-type:
        - core
    deny-auto-connection: true
`

const geoclueConnectedPlugAppArmor = `
# Description: Allow access to the GeoClue geolocation service.
#include <abstractions/dbus-strict>

# Allow access to GeoClue methods and signals
dbus (send, receive)
    bus=system
    interface=org.freedesktop.GeoClue2.*
    path=/org/freedesktop/GeoClue2/**
    peer=(label=unconfined),

# Allow access to GeoClue properties
dbus (send)
    bus=system
    path=/org/freedesktop/GeoClue2/**
    interface=org.freedesktop.DBus.Properties
    member="{Get,GetAll,Set}"
    peer=(label=unconfined),
dbus (receive)
    bus=system
    path=/org/freedesktop/GeoClue2/**
    interface=org.freedesktop.DBus.Properties
    member="PropertiesChanged"
    peer=(label=unconfined),
`

func init() {
	registerIface(&commonInterface{
		name:                  "geoclue",
		summary:               geoclueSummary,
		implicitOnCore:        true,
		implicitOnClassic:     true,
		baseDeclarationSlots:  geoclueBaseDeclarationSlots,
		connectedPlugAppArmor: geoclueConnectedPlugAppArmor,
	})
}
