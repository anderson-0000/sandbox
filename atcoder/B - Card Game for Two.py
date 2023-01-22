# https://atcoder.jp/contests/abc088/tasks/abc088_b

##Alice／Bobそれぞれの値だしてから差分出す
N = int(input())
a_list = list(map(int,input().split()))
a_list.sort(reverse=True)

Alice = Bob = count = 0

for i in a_list:
  count +=1
  if count%2 != 0:
    Alice += i
  else:
    Bob += i

print(Alice - Bob)
