#https://atcoder.jp/contests/abc081/tasks/abc081_b

N=int(input())
A_list = list(map(int, input().split()))
count = 0
flag = True

while flag:
  for i in range(N):
    if A_list[i] % 2 == 0: ##偶数
      A_list[i] = A_list[i]//2
    else: ##奇数
      flag = False
      break
  else:
    count += 1

print(count)
