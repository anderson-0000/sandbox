# https://atcoder.jp/contests/abc055/tasks/abc055_b

import math
N = int(input())

power = math.factorial(N-1)*N #math.factorial()は階乗

print(power%(10**9+7))
