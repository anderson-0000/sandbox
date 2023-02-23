import urllib.request
import time
import os

os.makedirs("tmp_user_html", exist_ok=True)

with open("tmp_all_user_id/all_user_id", mode="r", encoding="utf-8") as all_user_id:
    numbers = all_user_id.read().splitlines()

for number in numbers:
  url = "https://job-draft.jp/users/" + str(number) + ".html"
  filename = "tmp_user_html/" + str(number) + ".html"
  try:
      response = urllib.request.urlopen(url)
      status_code = response.getcode()
      if status_code == 200:
          print("Status: " + str(status_code))
          urllib.request.urlretrieve(url, filename)
      else:
          print("Status: " + str(status_code))
  except urllib.error.HTTPError as e:
    print("Error: " + str(e))
    continue
  finally:
    time.sleep(1)
