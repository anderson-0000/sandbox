#https://atcoder.jp/contests/abc096/tasks/abc096_c

H,W = map(int,input().split())
map_list = [input() for _ in range(H)]
h = 0
ans = 'Yes'

for line in map_list: #hカウントアップ
  w = 0
  for column in line: #wカウントアップ
    # 上下左右を確認
    if column == '#':
      if 0<= h-1 < H and 0<= w < W and map_list[h-1][w] == '#': #上
        w += 1
        continue
      elif 0<= h+1 < H and 0<= w < W and map_list[h+1][w] == '#': #下
        w += 1
        continue
      elif 0<= h < H and 0<= w-1 < W and map_list[h][w-1] == '#': #左
        w += 1
        continue
      elif 0<= h < H and 0<= w+1 < W and map_list[h][w+1] == '#': #右
        w += 1
        continue
      else:
        ans = 'No'
        break
    else:
      w += 1
  else: #内部ループでbrakeしなかった時。正常終了時
    h += 1
    continue
  break #breakで内部ループを抜けた時

print(ans)
