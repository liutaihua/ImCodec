package imcodec

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEncode(t *testing.T) {
	Convey("Encode", t, func() {
		enc := NewEncoder()
		enc.AddItem("id", 24060406792204)
		enc.AddItem("type", "chat")

		user_enc := NewEncoder()
		user_enc.AddItem("username", "MyLord")
		user_enc.AddItem("grade", 2)
		user_enc.AddItem("uid", 24637211)

		msg_enc := NewEncoder()
		msg_enc.AddItem("content", "我真是日了狗了")

		msg_enc.AddItem("via", 1)
		msg_enc.AddItem("user", user_enc.ToString())

		enc.AddItem("msg", msg_enc.ToString())

		b := enc.Bytes()
		t.Logf("after gzip %v", b)

		dec := NewDecoder()
		mmp := dec.DecodeGzipBytes(b)
		t.Logf("gzip decoded: %v", mmp)
	})
}

