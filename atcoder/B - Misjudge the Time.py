# https://atcoder.jp/contests/abc278/tasks/abc278_b

import datetime

H,M = map(int,input().split())
H = int(format(H, '02'))
M = int(format(M, '02'))

dt = datetime.datetime(2000, 1, 1, H, M)

dt_start_end_list = [
  [datetime.datetime(2000, 1, 1, 0, 0), datetime.datetime(2000, 1, 1, 5, 59)],
  [datetime.datetime(2000, 1, 1, 10, 0), datetime.datetime(2000, 1, 1, 15, 59)],
  [datetime.datetime(2000, 1, 1, 20, 0), datetime.datetime(2000, 1, 1, 20, 39)],
  [datetime.datetime(2000, 1, 1, 21, 0), datetime.datetime(2000, 1, 1, 21, 39)],
  [datetime.datetime(2000, 1, 1, 22, 0), datetime.datetime(2000, 1, 1, 22, 39)],
  [datetime.datetime(2000, 1, 1, 23, 0), datetime.datetime(2000, 1, 1, 23, 39)]
]

for i in range(0,len(dt_start_end_list)):
  if dt_start_end_list[i][0] <= dt <= dt_start_end_list[i][1]:
    break
  elif i == len(dt_start_end_list) - 1:
    H = 0
    M = 0
    break
  elif dt_start_end_list[i][1] < dt < dt_start_end_list[i+1][0]:
    H = dt_start_end_list[i+1][0].strftime('%H')
    M = dt_start_end_list[i+1][0].strftime('%M')
    break
print(H,M)
