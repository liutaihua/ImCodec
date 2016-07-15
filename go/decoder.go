package imcodec

import (
	"strings"
	"bytes"
	"io/ioutil"
	"compress/gzip"
)

type Decoder struct {
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (dc *Decoder) Decode(data string) (map[string]interface{}) {
	var idx int
	//idx := strings.Index(data, "/")
	//data = data[idx+1:]
	buff := strings.Split(data, "/")

	var (
		//idx int
		key string
		val interface{}
		rtnMsg = make(map[string]interface{})
	)
	for _, tmp := range buff {
		if idx = strings.Index(tmp, "@="); idx == -1 {
			continue
		}
		key = tmp[:idx]
		val = tmp[idx+2:]

		if strings.Contains(val.(string), "@A=") && strings.Contains(val.(string), "@S") {
			val = strings.Replace(strings.Replace(val.(string), "@S", "/", -1), "@A", "@", -1)
			val = dc.Decode(val.(string))
		} else {
			val = strings.Replace(strings.Replace(val.(string), "@S", "/", -1), "@A", "@", -1)
		}
		rtnMsg[key] = val
	}
	return rtnMsg
}

func (dc *Decoder) DecodeGzipBytes(b []byte) (map[string]interface{}) {
	var (
		err error
		bb []byte
		r *gzip.Reader
		buf = bytes.NewBuffer(b)
	)
	if r, err = gzip.NewReader(buf); err != nil {
		panic(err)
	}
	if bb, err = ioutil.ReadAll(r); err != nil {
		panic(err)
	}
	return dc.Decode(string(bb))
}