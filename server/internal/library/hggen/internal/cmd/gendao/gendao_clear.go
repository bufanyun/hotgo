package gendao

import (
	"context"

	"github.com/gogf/gf/v2/os/gfile"

	"hotgo/internal/library/hggen/internal/utility/mlog"
	"hotgo/internal/library/hggen/internal/utility/utils"
)

func doClear(ctx context.Context, dirPath string) {
	files, err := gfile.ScanDirFile(dirPath, "*.go", true)
	if err != nil {
		mlog.Fatal(err)
	}
	for _, file := range files {
		if utils.IsFileDoNotEdit(file) {
			if err = gfile.Remove(file); err != nil {
				mlog.Print(err)
			}
		}
	}
}
