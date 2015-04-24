package main

import (
	"flag"
	"fmt"
	"github.com/ShionRyuu/ydict/dict"
	"os"
	"strings"
)

var (
	showVersion bool
	engine      string
)

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.StringVar(&engine, "engine", "youdao", "Dict engine")
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
	var dic dict.Dict
	if strings.EqualFold(engine, "iciba") {
		dic = dict.NewYoudao()
	} else {
		dic = dict.NewIciba()
	}
	dic.Translate(os.Args[1])
}
