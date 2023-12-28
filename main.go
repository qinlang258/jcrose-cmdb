package main

import (
	_ "jcrose-cmdb/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"jcrose-cmdb/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
