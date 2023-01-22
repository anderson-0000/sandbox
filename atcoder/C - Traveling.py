#https://atcoder.jp/contests/abc086/tasks/arc089_a

#1移動で時間1を消費する
N = int(input())

plan = []
plan.append([0,0,0])
ans = 'Yes'

for i in range(N):
  t,xi,yi = map(int,input().split())
  plan.append([t,xi,yi])

for i in range(1,N+1):
  #時間の差を出す
  operation_count = plan[i][0] - plan[i-1][0]
  diff_x = abs(plan[i][1] - plan[i-1][1]) #絶対値をとって距離がマイナスにならないようにする
  diff_y = abs(plan[i][2] - plan[i-1][2]) #絶対値をとって距離がマイナスにならないようにする

  if diff_x + diff_y > operation_count:
    ans = 'No'
    break
  elif operation_count == diff_x + diff_y:
    i+=1
  elif (operation_count - diff_x - diff_y)%2 == 0: #調整回数が偶数ならOK
    i+=1
  elif  (operation_count - diff_x - diff_y)%2 != 0: #調整回数が奇数は戻ってくれないのでNG
    ans = 'No'
    break
print(ans)
