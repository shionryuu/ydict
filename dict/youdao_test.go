package dict

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	rest := Result{
		Translation: []string{
			"爱",
		},
		Basic: BasicTrans{
			Us_phonetic: "lʌv",
			Phonetic:    "lʌv",
			Uk_phonetic: "lʌv",
			Explains: []string{
				"n. 恋爱；亲爱的；酷爱；喜爱的事物；爱情，爱意；疼爱；热爱；爱人，所爱之物",
				"v. 爱，热爱；爱戴；赞美，称赞；喜爱；喜好；喜欢；爱慕",
				"n. （英）洛夫（人名）",
			},
		},
		Query:     "love",
		ErrorCode: 0,
		Web: []WebTrans{
			{
				Value: []string{
					"爱",
					"爱情",
					"恋爱",
				},
				Key: "Love",
			},
			{
				Value: []string{
					"无尽的爱",
					"蓝色生死恋",
					"不了情",
				},
				Key: "Endless Love",
			},
			{
				Value: []string{
					"早恋",
					"青春期恋爱",
					"初恋",
				},
				Key: "puppy love",
			},
		},
	}
	data, err := json.Marshal(rest)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	json.Unmarshal(data, &rest)
	fmt.Println(rest)
}
