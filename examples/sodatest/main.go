package main

import (
	"github.com/phimister/sodawater"
)

func main() {
	m := soda.NewModel(soda.WithTPS(30))
	m.Run()
}
