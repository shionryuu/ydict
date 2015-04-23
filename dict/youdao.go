package dict

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	API_FROM = "YouDaoCV"
	API_KEY  = "659600698"
	API_URL  = "http://fanyi.youdao.com/openapi.do?keyfrom=%s&key=%s&type=data&doctype=json&version=1.1&q=%s"
)

type WebTrans struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

type BasicTrans struct {
	Us_phonetic string   `json:"us-phonetic"`
	Phonetic    string   `json:"phonetic"`
	Uk_phonetic string   `json:"uk-phonetic"`
	Explains    []string `json:"explains"`
}

type Result struct {
	Translation []string   `json:"translation"`
	Basic       BasicTrans `json:"basic"`
	Query       string     `json:"query"`
	ErrorCode   int        `json:"errorCode"`
	Web         []WebTrans `json:"web"`
}

func Youdao(word string) {
	url := fmt.Sprintf(API_URL, API_FROM, API_KEY, word)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	if result.ErrorCode != 0 {
		printError(result)
		return
	}

	prettyPrint(result)
}

func printError(result Result) {
	color.Red("\n  查找不到指定的单词: %s\n", result.Query)
}

func prettyPrint(result Result) {
	fmt.Printf("  %s\n", color.WhiteString(result.Query))
	color.Cyan("\n  发音:")
	if result.Basic.Uk_phonetic != "" && result.Basic.Us_phonetic != "" {
		fmt.Printf("    英 /%s/     美 /%s/\n", color.YellowString(result.Basic.Uk_phonetic), color.YellowString(result.Basic.Us_phonetic))
	} else if result.Basic.Phonetic != "" {
		fmt.Printf("    /%s/", color.YellowString(result.Basic.Phonetic))
	} else {
	}
	color.Cyan("\n  翻译:")
	for _, trans := range result.Basic.Explains {
		fmt.Printf("    * %s\n", color.BlueString(trans))
	}
	color.Cyan("\n  网络释义:")
	for _, trans := range result.Web {
		fmt.Printf("    * %s\n", color.YellowString(trans.Key))
		fmt.Printf("      %s\n", color.MagentaString(strings.Join(trans.Value, "；")))
	}
}
