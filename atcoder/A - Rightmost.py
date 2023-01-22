# https://atcoder.jp/contests/abc276/tasks/abc276_a

S = input()

try:
  print([i+1 for i,alphabet in enumerate(S) if alphabet == 'a' ][-1])
except:
  print('-1')
