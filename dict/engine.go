package dict

import (
	"strings"
)

const (
	IcibaEngine  = "iciba"
	YoudaoEngine = "youdao"
)

type Engine struct {
	cur string
	dic Dict
}

func NewEngine(name string) *Engine {
	e := &Engine{}
	e.ReNew(name)
	return e
}

func (e *Engine) ReNew(name string) {
	lname := strings.ToLower(name)
	switch lname {
	case e.cur:
		break
	case IcibaEngine:
		e.dic = NewIciba()
		break
	default:
		e.dic = NewYoudao()
	}
}

func (e *Engine) Translate(word string) {
	if e == nil {
		return
	}
	e.dic.Translate(word)
}
