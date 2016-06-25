package com.im.codec.codec;

public class Encoder
{
    private StringBuffer buf = new StringBuffer();
    public String toString()
    {
    	buf.append('\0');
        return buf.toString();
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
