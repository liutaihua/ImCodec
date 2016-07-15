package imcodec

import (
	"strings"
	"fmt"
	"compress/gzip"
	"bytes"
)

type Encoder struct {
	buf []string
}

func NewEncoder() *Encoder {
	return &Encoder{
		buf: make([]string, 0),
	}
}

func (ed *Encoder) AddItem(k string, v interface{}) {
	key := strings.Replace(strings.Replace(k, "/", "@S", -1), "@", "@A", -1)
	ed.buf = append(ed.buf, key)
	ed.buf = append(ed.buf, "@=")
	var val string
	switch v.(type) {
	case int:
		val = strings.Replace(strings.Replace(fmt.Sprintf("%d", v.(int)), "@", "@A", -1), "/", "@S", -1)
	case int32:
		val = strings.Replace(strings.Replace(fmt.Sprintf("%d", v.(int32)), "@", "@A", -1), "/", "@S", -1)
	case int64:
		val = strings.Replace(strings.Replace(fmt.Sprintf("%d", v.(int64)), "@", "@A", -1), "/", "@S", -1)
	case string:
		val = strings.Replace(strings.Replace(v.(string), "@", "@A", -1), "/", "@S", -1)
	case []byte:
		val = strings.Replace(strings.Replace(string(v.([]byte)), "@", "@A", -1), "/", "@S", -1)
	default:
		fmt.Println("error type")
	}
	ed.buf = append(ed.buf, val)
	ed.buf = append(ed.buf, "/")
}

func (ed *Encoder) Bytes() []byte {
	//ed.buf = append(ed.buf, "\\0")
	var (
		s = strings.Join(ed.buf, "")
		buf bytes.Buffer
	)
	w := gzip.NewWriter(&buf)
	w.Write([]byte(s))
	w.Close()
	return buf.Bytes()
}

func (ed *Encoder) ToString() string {
	return strings.Join(ed.buf, "")
}