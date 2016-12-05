// demo.go
package main

import (
	"github.com/Tom-Kail/log"
)

func A() {
	B()
}
func B() {
	C()
}

func C() {
	log.Debug("fdfdf")
}

func main() {
	log.Log.SetLogFuncCallDepth(3)
	log.Debug("")
	A()
}
