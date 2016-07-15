package com.im.codec.example;

import com.im.codec.codec.Encoder;
import com.im.codec.codec.Decoder;
import org.apache.commons.lang3.StringUtils;

import java.util.HashMap;
import java.util.Map;
import java.lang.String;

public class ImCodec
{
	public static void main(String[] args)
	{
		String s = "id@=333223232/type@=userjoin/msg@=uid@A=332233@Susername@A=名字几个字@Sgrade@A=5@S/";

		Map<String, Object> l = Decoder.Decode(s);
		for (String key : l.keySet()) {
			System.out.println("Key: " + key + ", Value: " + l.get(key));
		}

		Encoder enc = new Encoder();
		enc.addItem("url", "http://0163.com");
		enc.addItem("name", "\\SSS@@@SSSAAA");
		enc.addItem("list", "[111, 2222, 333, 4444]");

		Encoder subEnc = new Encoder();
		subEnc.addItem("sname", "\\@S@S@A@A");
		subEnc.addItem("content", "@@@===@=@@S@@A@A@S测试一下,狗比你大业的你以为你很牛逼吗");

		enc.addItem("sub", subEnc.toString());
		enc.addItem("special", "@@@===@=@@S@@A@A@SSSAAA@S@A");


		String res = enc.toString();

		byte[] b = enc.Bytes();
		System.out.println("gzip encoded:" + b);
		System.out.println("encoded:" + res);

		l = Decoder.Decode(res);
		for (String key: l.keySet()) {
			Object v = l.get(key);
			System.out.println("Key: " + key + ", Value: " + v);
		}

		l = Decoder.GzipDecode(b);
		for (String key: l.keySet()) {
			Object v = l.get(key);
			System.out.println("Key: " + key + ", Value: " + v);
		}
	}
}
