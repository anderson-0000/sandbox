# https://atcoder.jp/contests/abc070/tasks/abc070_b

A,B,C,D = map(int,input().split())
ans = 0

for alice in range(A,B):
  if C<= alice <D:
    ans += 1
print(ans)
