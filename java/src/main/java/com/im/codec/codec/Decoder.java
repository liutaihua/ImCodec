package com.im.codec.codec;

import java.util.HashMap;
import java.util.Map;
import java.lang.String;
import org.apache.commons.lang3.StringUtils;

public class Decoder
{
    public static Map<String, Object> Decode(String data){
        Map<String, Object> rtnMsg = new HashMap<String, Object>();

        data = StringUtils.substringBeforeLast(data, "/");
        String[] buff = data.split("/");

        for(String tmp : buff){
            String key = StringUtils.substringBefore(tmp, "@=");
            Object value = StringUtils.substringAfter(tmp, "@=");
            if(StringUtils.contains((String)value, "@A") || StringUtils.contains((String)value, "@S")){
                value = ((String)value).replaceAll("@S", "/").replaceAll("@A", "@");
                if (StringUtils.contains((String)value, "@=")) {
                    value = Decode((String)value);
                }
            }
            rtnMsg.put(key, value);
        }

        return rtnMsg;

    }
}
