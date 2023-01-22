#https://atcoder.hp/contests/abc075/tasks/abc075_b

H,W = map(int,input().split())
map_list = [input() for _ in range(H)]
h = 0
ans = []

for line in map_list: #hカウントアップ
  w = 0
  ans_line = ''
  for column in line: #iカウントアップ
    # 上下左右斜めを確認
    if column == '.':
      count = 0
      if 0<= h-1 < H and 0<= w-1 < W and map_list[h-1][w-1] == '#': #左上
        count += 1
      if 0<= h-1 < H and 0<= w < W and map_list[h-1][w] == '#': #上
        count += 1
      if 0<= h-1 < H and 0<= w+1 < W and map_list[h-1][w+1] == '#': #右上
        count += 1
      if 0<= h < H and 0<= w-1 < W and map_list[h][w-1] == '#': #左
        count += 1
      if 0<= h < H and 0<= w+1 < W and map_list[h][w+1] == '#': #右
        count += 1
      if 0<= h+1 < H and 0<= w-1 < W and map_list[h+1][w-1] == '#': #左下
        count += 1
      if 0<= h+1 < H and 0<= w < W and map_list[h+1][w] == '#': #下
        count += 1
      if 0<= h+1 < H and 0<= w+1 < W and map_list[h+1][w+1] == '#': #右下
        count += 1
      ans_line += str(count)
    else:
      ans_line += '#'
    w += 1
  ans.append(ans_line)
  h += 1

print(*ans, sep='\n')
