package imcodec

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEncode(t *testing.T) {
	Convey("Encode", t, func() {
		enc := NewEncoder()
		enc.AddItem("id", 1249004792)
		enc.AddItem("type", "chat")

		sub_enc := NewEncoder()
		sub_enc.AddItem("username", "狗比")
		sub_enc.AddItem("grade", 6)
		sub_enc.AddItem("uid", 2323232)

		sub_enc2 := NewEncoder()
		sub_enc2.AddItem("content", "测试一下,狗比你大业的你以为你很牛逼吗")
		sub_enc2.AddItem("via", 2)
		sub_enc2.AddItem("user", sub_enc.ToString())

		enc.AddItem("msg", sub_enc2.ToString())

		s := enc.ToString()

		dec := NewDecoder()
		mp := dec.Decode(s)
		t.Log(mp)
		So(s, ShouldEqual, "id@=1249004792/type@=chat/msg@=content@A=测试一下,狗比你大业的你以为你很牛逼吗@Svia@A=2@Suser@A=username@AA=狗比@ASgrade@AA=6@ASuid@AA=2323232@AS@S/")
	})
}

