// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	api "github.com/gogits/go-gogs-client"

	"github.com/gogits/gogs/pkg/context"
	"github.com/gogits/gogs/routers/api/v1/repo"
	"github.com/gogits/gogs/routers/api/v1/user"
)

// https://github.com/gogits/go-gogs-client/wiki/Administration-Repositories#create-a-new-repository
func CreateRepo(c *context.APIContext, form api.CreateRepoOption) {
	owner := user.GetUserByParams(c)
	if c.Written() {
		return
	}

	repo.CreateUserRepo(c, owner, form)
}
