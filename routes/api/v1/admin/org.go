// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	api "github.com/farshmak/go-gogs-client"

	"github.com/farshmak/gogs/pkg/context"
	"github.com/farshmak/gogs/routes/api/v1/user"
	"github.com/farshmak/gogs/routes/api/v1/org"
)

// https://github.com/farshmak/go-gogs-client/wiki/Administration-Organizations#create-a-new-organization
func CreateOrg(c *context.APIContext, form api.CreateOrgOption) {
	org.CreateOrgForUser(c, form, user.GetUserByParams(c))
}
