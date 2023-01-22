#https://atcoder.jp/contests/abc087/tasks/arc090_a

N = int(input())
i1 = list(map(int,input().split()))
i2 = list(map(int,input().split()))
i1_sum=[]
i2_sum=[]

i1_sum.append(i1[0])
i2_sum.append(i1[0]+i2[0]) #↓

for i in range(1,N):
  i1_sum.append(i1_sum[i-1]+i1[i]) #→
  i1_i2 = i1_sum[i]+i2[i] #→↓
  i2_right = i2_sum[i-1]+i2[i] #→
  i2_sum.append(max(i1_i2,i2_right))

print(i2_sum[-1])
