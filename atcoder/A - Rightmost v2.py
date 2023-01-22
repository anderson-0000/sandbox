# https://atcoder.jp/contests/abc276/tasks/abc276_a

S = input()

for i in range(len(S), 0, -1):
  if S[i-1] == 'a':
    print(i)
    break
else: #breakしなかったら処理される
  print('-1')
