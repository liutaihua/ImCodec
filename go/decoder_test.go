package imcodec

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestEncode(t *testing.T) {
	Convey("Encode", t, func() {
		enc := NewEncoder()
		enc.AddItem("id", 1249004792)
		enc.AddItem("type", "chat")
		enc.AddItem("special", "@@@===@=@@S@@A@A@SSSAAA@S@A")
		enc.AddItem("time", "/Date(1467943242551+0800)/")

		sub_enc := NewEncoder()
		sub_enc.AddItem("username", "狗比")
		sub_enc.AddItem("grade", 6)
		sub_enc.AddItem("uid", 2323232)
		sub_enc.AddItem("content", "@@@===@=@@S@@A@A@S测试一下,狗比你大业的你以为你很牛逼吗")

		sub_enc2 := NewEncoder()
		sub_enc2.AddItem("content", "@@@===@=@@S@@A@A@S测试一下,狗比你大业的你以为你很牛逼吗")

		sub_enc2.AddItem("via", 2)
		sub_enc2.AddItem("user", sub_enc.ToString())

		enc.AddItem("msg", sub_enc2.ToString())

		s := enc.ToString()

		dec := NewDecoder()
		mp := dec.Decode(s)
		t.Log(mp)
		So(mp["time"], ShouldEqual, "/Date(1467943242551+0800)/")
		So(mp["special"], ShouldEqual, "@@@===@=@@S@@A@A@SSSAAA@S@A")
		So(mp["msg"].(map[string]interface{})["user"].(map[string]interface{})["content"], ShouldEqual, "@@@===@=@@S@@A@A@S测试一下,狗比你大业的你以为你很牛逼吗")

		dec2 := NewDecoder()
		mp2 := dec2.Decode("id@=24050743115806/type@=chat/msg@=content@A=我真是日了狗了@Suser@A=grade@AA=2@ASuid@AA=24637211@ASusername@AA=MyLord@AS@Svia@A=1@S/")
		t.Log(mp2)
		fmt.Println("-----------------", mp2)
	})
}

