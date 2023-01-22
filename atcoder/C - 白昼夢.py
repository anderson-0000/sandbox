#https://atcoder.jp/contests/abc049/tasks/arc065_a

beforeS=S=input()
ans='YES'

while len(S) != 0:
  if S.endswith('dreamer'):
    S = S[:-7]
  if S.endswith('eraser'):
    S = S[:-6]
  if S.endswith('dream') or S.endswith('erase'):
    S = S[:-5]
  if beforeS == S:
    ans = 'NO'
    break
  beforeS = S

print(ans)
