# https://atcoder.jp/contests/abc273/tasks/abc273_a

N=int(input())

anser = 1
for i in range(N,0,-1):
  anser *= i
print(anser)
