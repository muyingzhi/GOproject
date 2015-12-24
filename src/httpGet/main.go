package main

import (
	"fmt"
	"runtime"
)

const (
	UA = "Golang Downloader from Ijibu.com"
)

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}
