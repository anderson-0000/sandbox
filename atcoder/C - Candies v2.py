#https://atcoder.jp/contests/abc087/tasks/arc090_a

N = int(input())
i1 = list(map(int,input().split()))
i2 = list(map(int,input().split()))
i1_sum = i1[0]
i2_sum = i1[0]+i2[0] #↓

for i in range(1,N):
  i1_sum = i1_sum+i1[i] #→
  i2_sum = max(i1_sum+i2[i],i2_sum+i2[i])#→↓,→

print(i2_sum)
