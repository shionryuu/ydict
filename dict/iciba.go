package dict

import (
	"encoding/xml"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
<?xml version="1.0"?>
<dict num="219" id="219" name="219">
    <key>love</key>
    <ps>lʌv</ps>
    <pron>http://res.iciba.com/resource/amp3/oxford/0/4f/5b/4f5bbc0f19c33e5f1a0b6b974b4eacce.mp3</pron>
    <ps>lʌv</ps>
    <pron>http://res.iciba.com/resource/amp3/1/0/b5/c0/b5c0b187fe309af0f4d35982fd961d7e.mp3</pron>
    <pos>vt.&amp; vi.</pos>
    <acceptation>爱，热爱；爱戴；喜欢；赞美，称赞；</acceptation>
    <pos>vt.</pos>
    <acceptation>喜爱；喜好；喜欢；爱慕；</acceptation>
    <pos>n.</pos>
    <acceptation>爱情，爱意；疼爱；热爱；爱人，所爱之物；</acceptation>
    <sent>
        <orig>They happily reflect the desire for a fusional love that inspired the legendary LOVE bracelet Cartier.</orig>
        <trans>快乐地反映出为富有传奇色彩的卡地亚LOVE手镯所赋予的水乳交融之爱恋情愫.</trans>
    </sent>
    <sent>
        <orig>Love is the radical of lovely , loveliness , and loving.</orig>
        <trans>Love是lovely, loveliness 及loving的词根.</trans>
    </sent>
    <sent>
        <orig>She rhymes " love " with " dove ".</orig>
        <trans>她将 " love " 与 " dove " 两字押韵.</trans>
    </sent>
    <sent>
        <orig>In sports, love means nil.</orig>
        <trans>体育中, love的意思是零.</trans>
    </sent>
    <sent>
        <orig>Ludde Omholt with his son, Love, in S ? derma a bohemian and culturally rich district in Stockholm.</orig>
        <trans>LuddeOmholt和他的儿子Love在南城 —— 斯德哥尔摩市 的一个充满波西米亚风情的文化富饶区散步.</trans>
    </sent>
</dict>
*/

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
