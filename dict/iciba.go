package dict

import (
	"encoding/xml"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	Iciba_KEY = "D191EBD014295E913574E1EAF8E06666"
	Iciba_URL = "http://dict-co.iciba.com/api/dictionary.php?key=%s&w=%s"
)

type Sent struct {
	Orig  string `xml:"orig"`
	Trans string `xml:"trans"`
}

type IcibaTran struct {
	Key         string   `xml:"key"`
	Ps          []string `xml:"ps"`
	Pron        []string `xml:"pron"`
	Pos         []string `xml:"pos"`
	Acceptation []string `xml:"acceptation"`
	Sent        []Sent   `xml:"sent"`
}

type Iciba struct {
}

func NewIciba() *Iciba {
	return &Iciba{}
}

func (this *Iciba) Translate(word string) {
	url := fmt.Sprintf(Iciba_URL, Iciba_KEY, word)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("error: ", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	var result IcibaTran
	if err := xml.Unmarshal(body, &result); err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	this.prettyPrint(result)
}

func (this *Iciba) prettyPrint(result IcibaTran) {
	fmt.Printf("  %s\t%s\n", color.WhiteString(result.Key), color.GreenString("~ iciba"))
	color.Cyan("\n  发音:")
	if len(result.Ps) == 2 {
		fmt.Printf("    英 /%s/     美 /%s/\n", color.YellowString(result.Ps[0]), color.YellowString(result.Ps[1]))
	} else if len(result.Ps) == 1 {
		fmt.Printf("    /%s/", color.YellowString(result.Ps[0]))
	}
	color.Cyan("\n  翻译:")
	for i := 0; i < len(result.Pos); i++ {
		if len(result.Pos[i]) > 0 {
			fmt.Printf("    * %s %s\n", color.BlueString(strings.Trim(result.Pos[i], "\r\n")), color.BlueString(strings.Trim(result.Acceptation[i], "\r\n")))
		} else {
			fmt.Printf("    * %s\n", color.BlueString(strings.Trim(result.Acceptation[i], "\r\n")))
		}
	}
	color.Cyan("\n  双语例句:")
	for _, sent := range result.Sent {
		fmt.Printf("    * %s\n", color.YellowString(strings.Trim(sent.Orig, "\r\n")))
		fmt.Printf("      %s\n", color.MagentaString(strings.Trim(sent.Trans, "\r\n")))
	}
}
