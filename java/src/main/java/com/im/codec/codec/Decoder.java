package com.im.codec.codec;

import java.util.HashMap;
import java.util.Map;
import java.lang.String;
import org.apache.commons.lang3.StringUtils;
import java.io.BufferedReader;
import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.zip.GZIPInputStream;
import java.util.zip.GZIPOutputStream;


public class Decoder
{
    private static String decompress(byte[] compressed) throws IOException {
        ByteArrayInputStream bis = new ByteArrayInputStream(compressed);
        GZIPInputStream gis = new GZIPInputStream(bis);
        BufferedReader br = new BufferedReader(new InputStreamReader(gis, "UTF-8"));
        StringBuilder sb = new StringBuilder();
        String line;
        while((line = br.readLine()) != null) {
            sb.append(line);
        }
        br.close();
        gis.close();
        bis.close();
        return sb.toString();
    }

    private static String Decompress(byte[] compressed) {
        try {
            return decompress(compressed);
        } catch(Exception e) {
            e.printStackTrace();
            return null;
        }
    }

    public static Map<String, Object> GzipDecode(byte[] compressed) {
        String origin = Decompress(compressed);
        if (origin != null) {
            return Decode(origin);
        }
        return new HashMap<String, Object>();
    }

    public static Map<String, Object> Decode(String data){
        Map<String, Object> rtnMsg = new HashMap<String, Object>();
        data = StringUtils.substringBeforeLast(data, "/");
        String[] buff = data.split("/");

        for(String tmp : buff){
            String key = StringUtils.substringBefore(tmp, "@=");
            Object value = StringUtils.substringAfter(tmp, "@=");

            if(StringUtils.contains((String)value, "@A=") && StringUtils.contains((String)value, "@S")){
                value = ((String)value).replaceAll("@S", "/").replaceAll("@A", "@");
                value = Decode((String)value);
            } else {
                value = ((String)value).replaceAll("@S", "/").replaceAll("@A", "@");
            }
            rtnMsg.put(key, value);
        }
        return rtnMsg;

    }
}
