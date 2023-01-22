# https://atcoder.jp/contests/abc274/tasks/abc274_a

from decimal import Decimal, ROUND_HALF_UP

A,B = map(int, input().split())
AdivisionB= B/A

print(Decimal(str(AdivisionB)).quantize(Decimal('0.001'), rounding=ROUND_HALF_UP))
