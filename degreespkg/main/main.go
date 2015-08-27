package main

import (
	"degreespkg"
	//"degreespkg/structs"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS((numCPU * 3) / 4)
	degreespkg.StartServer()
}
