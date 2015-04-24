package dict

import (
	"encoding/xml"
	"testing"
)

func TestXml(t *testing.T) {
	rest := IcibaTran{
		Key: "love",
		Ps: []string{
			"lʌv",
			"lʌv",
		},
		Pron: []string{
			"http://res.iciba.com/resource/amp3/1/0/b5/c0/b5c0b187fe309af0f4d35982fd961d7e.mp3",
			"http://res.iciba.com/resource/amp3/1/0/b5/c0/b5c0b187fe309af0f4d35982fd961d7e.mp3",
		},
		Pos: []string{
			"vt.&amp; vi.",
			"vt.",
			"n.",
		},
		Acceptation: []string{
			"爱，热爱；爱戴；喜欢；赞美，称赞；",
			"喜爱；喜好；喜欢；爱慕；",
			"爱情，爱意；疼爱；热爱；爱人，所爱之物；",
		},
		Sent: []Sent{
			{
				Orig:  "They happily reflect the desire for a fusional love that inspired the legendary LOVE bracelet Cartier.",
				Trans: "快乐地反映出为富有传奇色彩的卡地亚LOVE手镯所赋予的水乳交融之爱恋情愫.",
			},
			{
				Orig:  "Love is the radical of lovely , loveliness , and loving.",
				Trans: "Love是lovely, loveliness 及loving的词根.",
			},
			{
				Orig:  `She rhymes " love " with " dove ".`,
				Trans: `她将 " love " 与 " dove " 两字押韵.`,
			},
			{
				Orig:  "In sports, love means nil.",
				Trans: "体育中, love的意思是零.",
			},
			{
				Orig:  "Ludde Omholt with his son, Love, in S ? derma a bohemian and culturally rich district in Stockholm.",
				Trans: "LuddeOmholt和他的儿子Love在南城 —— 斯德哥尔摩市 的一个充满波西米亚风情的文化富饶区散步.",
			},
		},
	}
	data, err := xml.Marshal(rest)
	if err != nil {
		t.Error("error: ", err.Error())
		return
	}
	xml.Unmarshal(data, &rest)
}
