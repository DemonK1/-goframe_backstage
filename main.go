package main

import (
	"01-practice/internal/cmd"
	_ "01-practice/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.MainRun()
}
