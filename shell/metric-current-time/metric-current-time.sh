#!/bin/bash
# サーバーを起動します
while true
do
 # 現在の時刻を秒とナノ秒で取得します
 now_seconds=$(gdate +%s)
 now_nanoseconds=$(gdate +%s%N)

 # HTTPレスポンスを作成します
 response="<html><body>"
 response+="<h1>Current Time</h1>"
 response+="<p>Seconds: ${now_seconds}</p>"
 response+="<p>Nanoseconds: ${now_nanoseconds}</p>"
 response+="</body></html>"

 # HTTPレスポンスを表示します gnu版ncを使ってください。
 echo -e "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: $(echo -n "${response}" | wc -c)\r\n\r\n${response}" | nc -l -p 50005
done

