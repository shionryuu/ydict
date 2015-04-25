package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ShionRyuu/ydict/dict"
	"os"
	"strings"
)

var (
	showVersion bool
	preferEng   string
	engine      *dict.Engine
)

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.StringVar(&preferEng, "engine", "youdao", "Dict engine")
}

func interactive() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("进人交互模式，按ctrl+c退出，:h帮助\n")
	for {
		fmt.Print("> ")
		bytes, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("error: ", err.Error())
			return
		}
		data := string(bytes)

		var command, rest string
		if strings.HasPrefix(data, ":") {
			arr := strings.SplitN(data, " ", 2)
			if len(arr) == 2 {
				command = arr[0]
				rest = arr[1]
			} else {
				command = arr[0]
			}
		} else {
			command = ":re"
			rest = data
		}

		switch command {
		case ":h":
			fmt.Println(`
:h 帮助
:e 修改词典引擎
ctrl+c 退出
			`)
			break
		case ":e":
			engine.ReNew(rest)
			break
		default:
			engine.Translate(rest)
		}
	}
}

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "-help") {
		flag.PrintDefaults()
		return
	}
	flag.Parse()
	if showVersion == true {
		fmt.Println(dict.Version)
		return
	}
	engine = dict.NewEngine(preferEng)
	if len(flag.Args()) > 0 {
		engine.Translate(flag.Args()[0])
	} else {
		interactive()
	}
}
