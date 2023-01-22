#https://atcoder.jp/contests/abc085/tasks/abc085_c

## OK 全検索
N,Y = map(int,input().split())

an_x, an_y, an_z = -1, -1, -1

for y in range(N + 1):
  for z in range(N + 1):
    x = N-z-y
    if 0 <= x and 10000*x+5000*y+1000*z == Y:
      an_x, an_y, an_z = x, y, z
print(an_x, an_y, an_z)
