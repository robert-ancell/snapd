// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2023 Canonical Ltd
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

package daemon

import (
	"net/http"

	"github.com/snapcore/snapd/overlord/auth"
)

var (
	accessoriesChangeCmd = &Command{
		Path: "/v2/refresh-status",
		GET:  getRefreshStatus,
		ReadAccess: refreshStatusOpenAccess{},
	}
)

func getRefreshStatus(c *Command, r *http.Request, user *auth.UserState) Response {
        // FIXME: Per session info.
	m := map[string]interface{}{
		"system-restart-required":      false,
		"session-restart-required":     false,
		"application-restart-required": [],
		}

	return SyncResponse(m)
}
