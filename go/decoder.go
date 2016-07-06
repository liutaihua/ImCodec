package imcodec

import (
	"strings"
)

type Decoder struct {
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (dc *Decoder) Decode(data string) (map[string]interface{}) {
	//idx := strings.Index(data, "/")
	//data = data[idx+1:]
	buff := strings.Split(data, "/")

	var (
		idx int
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
		if strings.Contains(val.(string), "@A") || strings.Contains(val.(string), "@S") {
			val = strings.Replace(strings.Replace(val.(string), "@S", "/", -1), "@A", "@", -1)
			if strings.Contains(val.(string), "@=") {
				val = dc.Decode(val.(string))
			}
		}
		rtnMsg[key] = val
	}
	return rtnMsg
}