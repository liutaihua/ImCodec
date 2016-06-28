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
    enc.add_item("url", "http://0163.com");
    enc.add_item("name", "\\SSS@@@@AA\0AA");

    encoder sub_enc;
    sub_enc.add_item("child_name", "Hama");
    sub_enc.add_item("child_age", 99);

    enc.add_item("child", sub_enc.to_string().c_str());

    string enc_str = enc.to_string();
    cout << "encode: " << enc_str << endl;
    cout << endl;

    cout << "decode res:" << endl;
    decoder dec;
    dec.parse(enc_str.c_str());

    string url = dec.get_item_as_string("url");
    string name = dec.get_item_as_string("name");

    cout << "url: " << url << ", name: " << name << endl;

    decoder child_dec;
    string child_str = dec.get_item_as_string("child");
    child_dec.parse(child_str.c_str());
    string child_name = child_dec.get_item_as_string("child_name");
    int child_age = child_dec.get_item_as_int("child_age");

    cout << "child info, name: " << child_name << " age: " << std::to_string(child_age) << endl;
}
