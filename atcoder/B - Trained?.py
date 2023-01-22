# https://atcoder.jp/contests/abc065/tasks/abc065_b

N = int(input())
a = [int(input()) for _ in range(N)]
checked_list = []
flag = True

button = a[0]
list_index = a.index(button)
count = 0

while True:
  count += 1
  button = a[list_index]
  if button == 2:
    print(count)
    break
  elif list_index in checked_list:
    print(-1)
    break
  checked_list.append(list_index)
  list_index=button-1
  if count == N:
    print(-1)
    break
