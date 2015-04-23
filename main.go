package main

import (
	"flag"
	"fmt"
	"github.com/ShionRyuu/ydict/dict"
	"os"
)

var (
	showVersion bool
)

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show version")
}

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "-help") {
		flag.PrintDefaults()
		return
	}
	if len(os.Args) < 2 {
		return
	}
	flag.Parse()
	if showVersion == true {
		fmt.Println(dict.Version)
		return
	}
	dict.Youdao(os.Args[1])
}
