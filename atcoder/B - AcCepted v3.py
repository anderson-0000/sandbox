# https://atcoder.jp/contests/abc104/tasks/abc104_b

S = input()

print('AC' if S[0] == 'A' and S[2:-1].count('C') == 1 and sum(map(str.isupper, S)) == 2 else 'WA')
