import numpy as np
import matplotlib.pyplot as plt

# 角度を θ として用意（ラジアン）
theta = np.linspace(0, 2 * np.pi, 100) # np.linspace(開始、終了、プロットの多さ)
y     = np.sin(theta)

plt.figure(figsize=(8, 4))
plt.plot(theta, y, color='blue', linewidth=2)

# タイトルと軸ラベル
plt.title('y = sin(θ)')
plt.xlabel('x')
plt.ylabel('y')

# 目盛の設定
xticks = [0, np.pi/2, np.pi, 3*np.pi/2, 2*np.pi]
xlabels = ['0', 'π/2', 'π', '3π/2', '2π']
plt.xticks(xticks, xlabels)

plt.yticks([-1, -0.5, 0, 0.5, 1])
plt.ylim(-1.1, 1.1)

plt.grid(True)
plt.show()

