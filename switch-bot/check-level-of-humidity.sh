#!/bin/bash
echo -e "\n\n\n---絶対湿度を確認します 8~10g/m3 が適正です---"
rm -rf tmp
mkdir tmp

token="Enter your switch bot token"
meter_list=$(curl -sS -X GET "https://api.switch-bot.com/v1.0/devices" -H "Authorization:${token}" | jq -r '.body.deviceList[] | select(.deviceType == "Meter")' | jq -s)

check_level_of_humidity() {
  device_id=$(echo ${meter_list} | jq -r ".[$1].deviceId")
  device_name=$(echo ${meter_list} | jq -r ".[$1].deviceName")
  meter_status=$(curl -sS -X GET "https://api.switch-bot.com/v1.0/devices/${device_id}/status" -H "Authorization:${token}" | jq -r '.body')
  temperature=$(echo ${meter_status} | jq -r ".temperature")
  relative_humidity=$(echo ${meter_status} | jq -r ".humidity")
  absolute_humidity=$(python3 -c "print('{:.2f}'.format(217 * (6.1078 * 10 ** (7.5 * ${temperature} / (${temperature} + 237.3 ))) / ( ${temperature} + 273.15 ) * (${relative_humidity} / 100)))")

  echo "${device_name} : ${absolute_humidity}" | tee -a tmp/output.txt

  if [ $(echo "${absolute_humidity} < 8" | bc) == 1 ]  # 湿度が低い
  then
    if [ ${device_id} != "対象から外したいdevice_id" ]
    then
      echo "${device_name} : ${absolute_humidity}" >> tmp/low_humidity.txt
    fi
  elif [ $(echo "10 < ${absolute_humidity}" | bc) == 1 ]  # 湿度が高い
  then
    if [ ${device_id} != "対象から外したいdevice_id" ]
    then
      echo "${device_name} : ${absolute_humidity}" >> tmp/high_humidity.txt
    fi
  fi
}

meter_list_length=$(echo ${meter_list} | jq length)

#同時実行は5まで
for ((i=0; i<${meter_list_length}; i++))                                                      
do
  mkfifo tmp/check_level_of_humidity${i}
  check_level_of_humidity ${i} &
  while [ $(jobs | wc -l) == 5 ]
  do
    sleep 0.5
  done
done

sleep 3
if [ -e tmp/output.txt ]
then
  message=$(cat tmp/output.txt)
else
  message="ファイルが見つかりません"
  echo "ファイルが見つかりません"
fi

if [ -e tmp/low_humidity.txt ]
then
  echo -e "\n---湿度が低い部屋---"
  cat tmp/low_humidity.txt
fi

if [ -e tmp/high_humidity.txt ]
then
  echo -e "\n---湿度が高い部屋---"
  cat tmp/high_humidity.txt
fi

termux-notification --title ">>>>>絶対湿度<<<<<" --content "${message}"

sleep 10
rm -rf tmp
sleep 10
