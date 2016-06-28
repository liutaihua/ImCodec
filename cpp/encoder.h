#ifndef __DY_CODEC_H__
#define __DY_CODEC_H__

#include <string>
#include <vector>

using namespace std;

class encoder
{
private:
    string buf;

public:
    encoder();
    ~encoder();

    string to_string();

    void add_item(const char *key, const char *value);
    void add_item(const char *key, const int value);
};

struct key_value
{
    string key;
    string value;
};


class decoder
{
public:
    typedef vector<key_value> item_array;

private:
    item_array arr;

public:
    decoder();
    ~decoder();

    void parse(const char * data);

    string get_item_as_string(const char *key);
    int get_item_as_int(const char *key);
};

#endif      //__DY_CODEC_H__
