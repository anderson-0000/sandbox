import urllib.request
import time
import os

os.makedirs("tmp_user_rank", exist_ok=True)

for i in range(1, 675):
  url = "https://job-draft.jp/users?page=" + str(i) + ".html"
  filename = "tmp_user_rank/" + str(i) + ".html"
  try:
    with urllib.request.urlopen(url) as response:
      status_code = response.getcode()
      if status_code == 200:
        print("Status: " + str(status_code))
        with open(filename, mode="wb") as f:
          f.write(response.read())
      else:
        print("Status: " + str(status_code))
  except urllib.error.HTTPError as e:
    print("Error: " + str(e))
    continue
  finally:
    time.sleep(1)
