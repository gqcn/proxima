package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"proxima/app/word/internal/cmd"
	_ "proxima/app/word/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
