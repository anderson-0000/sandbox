# https://atcoder.jp/contests/abc083/tasks/abc083_b

N,A,B = map(int,input().split())
anser = 0

for i in range(1,N+1):
  list_i=list(map(int,str(i)))
  if A <= sum(list_i)<= B:
    anser += i
print(anser)
