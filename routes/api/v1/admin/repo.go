// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	api "github.com/farshmak/go-gogs-client"

	"github.com/farshmak/gogs/pkg/context"
	"github.com/farshmak/gogs/routes/api/v1/repo"
	"github.com/farshmak/gogs/routes/api/v1/user"
)

// https://github.com/farshmak/go-gogs-client/wiki/Administration-Repositories#create-a-new-repository
func CreateRepo(c *context.APIContext, form api.CreateRepoOption) {
	owner := user.GetUserByParams(c)
	if c.Written() {
		return
	}

	repo.CreateUserRepo(c, owner, form)
}
