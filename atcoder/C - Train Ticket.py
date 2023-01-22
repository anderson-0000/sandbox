#https://atcoder.jp/contests/abc079/tasks/abc079_c

A,B,C,D = list(map(int,(input())))

if A+B+C+D==7:
  print('{}+{}+{}+{}=7'.format(A,B,C,D))
elif A+B+C-D==7:
  print('{}+{}+{}-{}=7'.format(A,B,C,D))
elif A+B-C-D==7:
  print('{}+{}-{}-{}=7'.format(A,B,C,D))
elif A+B-C+D==7:
  print('{}+{}-{}+{}=7'.format(A,B,C,D))
elif A-B-C-D==7:
  print('{}-{}-{}-{}=7'.format(A,B,C,D))
elif A-B-C+D==7:
  print('{}-{}-{}+{}=7'.format(A,B,C,D))
elif A-B+C+D==7:
  print('{}-{}+{}+{}=7'.format(A,B,C,D))
elif A-B+C-D==7:
  print('{}-{}+{}-{}=7'.format(A,B,C,D))
