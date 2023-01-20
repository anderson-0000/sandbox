#!/usr/bin/env python3

import urllib.request
import urllib.parse
import json
import logging

def open_weather_request(request):
  try:
    response = urllib.request.urlopen(request)
  except urllib.error.URLError as error:
    body = error.reason
    status = error.code
  else:
    body = json.loads(response.read())
    status = response.getcode()
  return body, status

# メインの処理
longitude = 'your longitude' #緯度
latitude = 'your latitude' #経度

api_key = 'your api key'
url = 'https://api.openweathermap.org/data/2.5/weather'
parameters = 'lon=' + longitude + '&lat=' + latitude + '&appid=' + api_key

# loggerの設定
logger = logging.getLogger()
logger.setLevel(logging.INFO)
handler = logging.StreamHandler()
handler.setLevel(logging.INFO)
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)

request = urllib.request.Request(url + '?' + parameters, method='GET')
response = open_weather_request(request)

if response[1] != 200:
  logger.error(str(response[1]) + ' ' + response[0])
#else:
#  logger.info(response[0])

response_body = response[0]

logger.info(response_body['weather'][0]['main'])
