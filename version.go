package main

import (
	"fmt"
	"runtime"
)

var Version string //版本号由Makefile指定

func printVersion() {
	fmt.Printf("===============================\nv2ray_simple %v (%v), %v %v %v\n", Version, desc, runtime.Version(), runtime.GOOS, runtime.GOARCH)

}
