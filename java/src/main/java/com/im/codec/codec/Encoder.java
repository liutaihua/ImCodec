package com.im.codec.codec;

import java.io.BufferedReader;
import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.zip.GZIPInputStream;
import java.util.zip.GZIPOutputStream;

public class Encoder
{
    private StringBuffer buf = new StringBuffer();

    private byte[] compress(String data) throws IOException {
        ByteArrayOutputStream bos = new ByteArrayOutputStream(data.length());
        GZIPOutputStream gzip = new GZIPOutputStream(bos);
        gzip.write(data.getBytes());
        gzip.close();
        byte[] compressed = bos.toByteArray();
        bos.close();
        return compressed;
    }

    public String toString()
    {
    	buf.append('\0');
        return buf.toString();
    }

    public byte[] Bytes()
    {
        String s = toString();
        try {
            return compress(s);
        } catch(Exception e) {
            e.printStackTrace();
            return null;
        }
    }

    public void addItem(String key, Object value)
    {
        buf.append(key.replaceAll("/", "@S").replaceAll("@", "@A"));
        buf.append("@=");
        if(value instanceof String){
            buf.append(((String)value).replaceAll("@", "@A").replaceAll("/", "@S"));
        }else if(value instanceof Integer){
            buf.append(value);
        }
        buf.append("/");
    }
}
