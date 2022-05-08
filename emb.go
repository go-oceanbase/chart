package main

import (
	"github.com/zserge/lorca"
)

func embed() {
	lorca.Embed("main", "asset.go", "tpl")
}
