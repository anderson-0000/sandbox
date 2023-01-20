#!/bin/bash
echo -e "\n\n\n空いている窓の確認を開始します"
rm -rf tmp
mkdir tmp

token="Enter your switch bot token"
window_list=$(curl -sS -X GET "https://api.switch-bot.com/v1.0/devices" -H "Authorization:${token}" | jq -r '.body.deviceList[] | select(.deviceType == "Contact Sensor")' | jq -s)

check_window_status() {
  device_id=$(echo ${window_list} | jq -r ".[$1].deviceId")
  device_name=$(echo ${window_list} | jq -r ".[$1].deviceName")
  window_status=$(curl -sS -X GET "https://api.switch-bot.com/v1.0/devices/${device_id}/status" -H "Authorization:${token}" | jq -r '.body.openState')

  if [ ${window_status} != "close" ]
  then
    echo "${device_name}" | tee -a tmp/output.txt
  fi
}

window_list_length=$(echo ${window_list} | jq length)

#同時実行は5まで
for ((i=0; i<${window_list_length}; i++))                                                      
do
  mkfifo tmp/check_window_status${i}
  check_window_status ${i} &
  while [ $(jobs | wc -l) == 5 ]
  do
    sleep 0.5
  done
done

if [ -e tmp/output.txt ]
then
  message=$(cat tmp/output.txt)
else
  message="なし"
  echo "なし"
fi

termux-notification --title ">>>>>空いている窓<<<<<" --content "${message}"
rm -rf tmp
sleep 20
