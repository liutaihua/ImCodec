#include <iostream>
#include <string>
#include <stdint.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include "encoder.h"

using namespace std;


int main(int argc, char **args)
{
    encoder enc;
    enc.add_item("id", 9838939384);
    enc.add_item("type", "chat");
    enc.add_item("time", "\/Date(1467943242551+0800)\/");

    encoder sub_enc;
    sub_enc.add_item("username", "狗比");
    sub_enc.add_item("grade", 6);
    sub_enc.add_item("uid", 2323232);

    encoder sub_enc2;
    sub_enc2.add_item("content", "测试一下,狗比你大业的你以为你很牛逼吗");
    sub_enc2.add_item("via", 2);
    sub_enc2.add_item("user", sub_enc.to_string().c_str());

    enc.add_item("msg", sub_enc2.to_string().c_str());

    string enc_str = enc.to_string();
    cout << "encode: " << enc_str << "\nsize:" << sizeof(enc_str) << endl;
    cout << endl;

    cout << "decode res:" << endl;
    decoder dec;
    dec.parse(enc_str.c_str());

    string type = dec.get_item_as_string("type");
    int id = dec.get_item_as_int("id");

    cout << "type: " << type << ", id: " << id << endl;

    decoder child_dec;
    string child_str = dec.get_item_as_string("msg");
    child_dec.parse(child_str.c_str());
    string content = child_dec.get_item_as_string("content");

    string time = dec.get_item_as_string("time");

    cout << "content: " << content << endl;
    cout << "time: " << time << endl;
}
