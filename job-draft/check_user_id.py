import os
import re

os.makedirs("tmp_all_user_id", exist_ok=True)
file_count = len(os.listdir("tmp_user_rank"))

with open("tmp_all_user_id/all_user_id", mode="w", encoding="utf-8") as all_user_id:
  for i in range(1, file_count+1):
    filename = "tmp_user_rank/" + str(i) + ".html"
    with open(filename, encoding="utf-8") as html_file:
      match_list = re.findall(r'href="/users/(\d+)"', html_file.read())
      for match in match_list:
        all_user_id.write(match + "\n")

with open("tmp_all_user_id/all_user_id", mode="r", encoding="utf-8") as tmp_all_user_id:
  sorted_list = sorted(tmp_all_user_id.read().splitlines(), key=int)

with open("tmp_all_user_id/all_user_id", mode="w", encoding="utf-8") as tmp_all_user_id:
  tmp_all_user_id.write("\n".join(sorted_list) + "\n")
