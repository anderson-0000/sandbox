# https://atcoder.jp/contests/abc104/tasks/abc104_b

S = input()
anser = 'WA'

if S[0:1] == 'A' and S.count('C', 2 , -1) == 1:
  if sum(map(str.isupper, S)) == 2:
    anser = 'AC'

print(anser)
