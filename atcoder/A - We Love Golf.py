# https://atcoder.jp/contests/abc165/tasks/abc165_a

K = int(input())
A,B = map(int,input().split())

flag = False

for i in range(A, B+1):
  if i%K == 0:
    flag = True
    break

print('OK' if flag else 'NG')
