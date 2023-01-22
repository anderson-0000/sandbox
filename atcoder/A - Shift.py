# https://atcoder.jp/contests/abc278/tasks/abc278_a

N,K = map(int,input().split())
A = list(map(int,input().split()))

for i in range(K):
  del A[:1]
  A.append(0)

print(*A)
