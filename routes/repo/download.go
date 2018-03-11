// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"fmt"
	"io"
	"path"

	"github.com/gogits/git-module"

	"github.com/gogits/gogs/pkg/context"
	"github.com/gogits/gogs/pkg/setting"
	"github.com/gogits/gogs/pkg/tool"
)

func ServeData(c *context.Context, name string, reader io.Reader, isAttachable bool) error {
	buf := make([]byte, 1024)
	n, _ := reader.Read(buf)
	if n >= 0 {
		buf = buf[:n]
	}
	if (!tool.IsTextFile(buf) && !tool.IsImageFile(buf)) || isAttachable {
		c.Resp.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, name))
		c.Resp.Header().Set("Content-Transfer-Encoding", "binary")
	} else if !setting.Repository.EnableRawFileRenderMode || !c.QueryBool("render") {
		c.Resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
	}
	c.Resp.Write(buf)
	_, err := io.Copy(c.Resp, reader)
	return err
}

func ServeBlob(c *context.Context, blob *git.Blob, isAttachable bool) error {
	dataRc, err := blob.Data()
	if err != nil {
		return err
	}

	return ServeData(c, path.Base(c.Repo.TreePath), dataRc, isAttachable)
}

func SingleDownload(c *context.Context) {
	blob, err := c.Repo.Commit.GetBlobByPath(c.Repo.TreePath)
	if err != nil {
		if git.IsErrNotExist(err) {
			c.Handle(404, "GetBlobByPath", nil)
		} else {
			c.Handle(500, "GetBlobByPath", err)
		}
		return
	}

	if err = ServeBlob(c, blob, false); err != nil {
		c.Handle(500, "ServeBlob", err)
	}
}

func SingleFileDownload(c *context.Context) {
	blob, err := c.Repo.Commit.GetBlobByPath(c.Repo.TreePath)
	if err != nil {
		if git.IsErrNotExist(err) {
			c.Handle(404, "GetBlobByPath", nil)
		} else {
			c.Handle(500, "GetBlobByPath", err)
		}
		return
	}
	if err = ServeBlob(c, blob, true); err != nil {
		c.Handle(500, "ServeBlob", err)
	}
}
