#https://api.slack.com/methods
#curl -sS "https://slack.com/api/users.profile.get" -H "Authorization: Bearer ${slack_get_user_profile_token}" -H "Content-Type: application/x-www-form-urlencoded" -d "user=${slack_userid}" | jq -r .profile.phone

import urllib.request
import json
import logging
import os

def slack_request(request):
  try:
    response = urllib.request.urlopen(request)
  except urllib.error.URLError as error:
    body = error.reason
    status = error.code
  else:
    body = json.loads(response.read())
    status = response.getcode()
  return body, status

# loggerの設定
logger = logging.getLogger()
logger.setLevel(logging.INFO)
handler = logging.StreamHandler()
handler.setLevel(logging.INFO)
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)

# メインの処理
users_profile_get_url = 'https://slack.com/api/users.profile.get'
slack_get_user_profile_token = os.getenv('slack_get_user_profile_token')
slack_userid = os.getenv('slack_userid')

header = {
  "Authorization": "Bearer " + str(slack_get_user_profile_token),
  "Content-Type": "application/x-www-form-urlencoded"
}
parameter_dics = {
  "user": slack_userid
}

parameter = urllib.parse.urlencode(parameter_dics)
request = urllib.request.Request(users_profile_get_url + '?' + parameter, headers=header, method='GET')
response = slack_request(request)

if response[1] != 200:
  logger.error(str(response[1]) + ' ' + response[0])
#else:
#  logger.info(response[0])

response_body = response[0]

logger.info(response_body['profile']['phone'])
