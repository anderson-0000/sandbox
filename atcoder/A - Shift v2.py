# https://atcoder.jp/contests/abc278/tasks/abc278_a

N,K = map(int,input().split())
A = list(map(int,input().split()))

[(A.pop(0),A.append(0)) for i in range(K)]

print(*A)
