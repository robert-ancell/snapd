// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017 Canonical Ltd
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

package mount

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry describes an /etc/fstab-like mount entry.
//
// Fields are named after names in struct returned by getmntent(3).
//
// struct mntent {
//     char *mnt_fsname;   /* name of mounted filesystem */
//     char *mnt_dir;      /* filesystem path prefix */
//     char *mnt_type;     /* mount type (see Mntent.h) */
//     char *mnt_opts;     /* mount options (see Mntent.h) */
//     int   mnt_freq;     /* dump frequency in days */
//     int   mnt_passno;   /* pass number on parallel fsck */
// };
type Entry struct {
	Name    string
	Dir     string
	Type    string
	Options []string

	DumpFrequency   int
	CheckPassNumber int
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// EqualEntries checks if one entry is equal to another
func (a *Entry) Equal(b *Entry) bool {
	return (a.Name == b.Name && a.Dir == b.Dir && a.Type == b.Type &&
		equalStrings(a.Options, b.Options) && a.DumpFrequency == b.DumpFrequency &&
		a.CheckPassNumber == b.CheckPassNumber)
}

// escape replaces whitespace characters so that getmntent can parse it correctly.
//
// According to the manual page, the following characters need to be escaped.
//  space     => (\040)
//  tab       => (\011)
//  newline   => (\012)
//  backslash => (\134)
func escape(s string) string {
	return whitespaceEscape.Replace(s)
}

// unescape replaces escape sequences used by setmnt with whitespace characters.
//
// According to the manual page, the following characters need to be unescaped.
//  space     <= (\040)
//  tab       <= (\011)
//  newline   <= (\012)
//  backslash <= (\134)
func unescape(s string) string {
	return whitespaceUnescape.Replace(s)
}

var (
	whitespaceEscape = strings.NewReplacer(
		" ", `\040`, "\t", `\011`, "\n", `\012`, "\\", `\134`)
	whitespaceUnescape = strings.NewReplacer(
		`\040`, " ", `\011`, "\t", `\012`, "\n", `\134`, "\\")
)

func (e Entry) String() string {
	// Name represents name of the device in a mount entry.
	name := "none"
	if e.Name != "" {
		name = escape(e.Name)
	}
	// Dir represents mount directory in a mount entry.
	dir := "none"
	if e.Dir != "" {
		dir = escape(e.Dir)
	}
	// Type represents file system type in a mount entry.
	fsType := "none"
	if e.Type != "" {
		fsType = escape(e.Type)
	}
	// Options represents mount options in a mount entry.
	options := "defaults"
	if len(e.Options) != 0 {
		options = escape(strings.Join(e.Options, ","))
	}
	return fmt.Sprintf("%s %s %s %s %d %d",
		name, dir, fsType, options, e.DumpFrequency, e.CheckPassNumber)
}

// ParseEntry parses a fstab-like entry.
func ParseEntry(s string) (Entry, error) {
	var e Entry
	fields := strings.Fields(s)
	// do all error checks before any assignments to `e'
	if len(fields) != 6 {
		return e, fmt.Errorf("expected exactly six fields, found %d", len(fields))
	}
	df, err := strconv.Atoi(fields[4])
	if err != nil {
		return e, fmt.Errorf("cannot parse dump frequency: %s", err)
	}
	cpn, err := strconv.Atoi(fields[5])
	if err != nil {
		return e, fmt.Errorf("cannot parse check pass number: %s", err)
	}
	e.Name = unescape(fields[0])
	e.Dir = unescape(fields[1])
	e.Type = unescape(fields[2])
	e.Options = strings.Split(unescape(fields[3]), ",")
	e.DumpFrequency = df
	e.CheckPassNumber = cpn
	return e, nil
}
