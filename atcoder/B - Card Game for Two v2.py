# https://atcoder.jp/contests/abc088/tasks/abc088_b

##前から順番にたし引きしていく
N = int(input())
a_list = list(map(int,input().split()))
a_list.sort(reverse=True)
anser = 0

for i in range(N):
  if i%2 == 0:
    anser += a_list[i]
  else:
    anser -= a_list[i]

print(anser)
