options(repos = c(CRAN = "https://cloud.r-project.org"))

#install.packages("showtext")
#install.packages("Cairo")
#install.packages("ragg")
#install.packages("sysfonts")
#install.packages("curl")

# パッケージリスト
required_packages <- c("showtext", "Cairo", "ragg", "sysfonts", "curl")

for (package_name in required_packages) {
    if (!(package_name %in% installed.packages())) {
        install.packages(package_name)
    }
}

# パッケージのロード
lapply(required_packages, require, character.only = TRUE)

#library(showtext)
#library(ragg)
#library(sysfonts)

# フォント指定
font_add_google("Noto Sans JP", "notosans-jp")

## Rの対話モードの時だけ
#showtext_auto()

## PNGで図形を出力したいとき1
# library(Cairo)
# CairoPNG("output_cairo.png", width=480, height=480)

## PNGで図形を出力したいとき2
library(ragg)
ragg::agg_png("output_ragg.png", width = 8, height = 6, units = "in", res = 300)

# キャンバス作成
plot(1, type ="n", xlim = c(0, 10), ylim = c(0, 10), xlab = "", ylab = "", main = "図形test")

# 線
lines(c(1, 9), c(9, 4))
lines(c(1, 9), c(9, 3), col = "red")
lines(c(1, 9), c(9, 2), col = "red", lwd = 1)
lines(c(1, 9), c(9, 1), col = "red", lwd = 2)

# 四角形
rect(2, 2, 4, 4)
rect(3, 3, 5, 5, lwd = 2)
rect(5, 5, 6, 6, col = "blue")
rect(5, 5, 7, 7, border = "yellow")
rect(2, 2, 4, 4, col = "lightblue", border = "blue")

# 円
symbols(4, 8, circles=0.5, inches = FALSE, add = TRUE)
symbols(7, 7, circles = 1, inches = FALSE, add = TRUE, bg = "pink")

# 三角形
lines(c(2, 8, 5, 2), c(2, 2, 8, 2), col="red", lwd=2)

# y=4
abline(h = 4, col = "gray", lty = 2)

# y=sin(x)
lines(seq(0, 10), sin(seq(0, 10)), col = "blue", lwd = 1)
lines(seq(0, 10, length.out = 100), sin(seq(0, 10, length.out = 100)), col = "blue", lwd = 1)
# y=sin(x)+4
lines(seq(0, 10, length.out = 100), sin(seq(0, 10, length.out = 100)) + 4, col = "blue", lwd = 1)

# y=cos(x)
lines(seq(0, 10, length.out = 100), cos(seq(0, 10, length.out = 100)), col = "forestgreen", lwd = 2)
# y=cos(x)+4
lines(seq(0, 10, length.out = 100), cos(seq(0, 10, length.out = 100)) + 4, col = "forestgreen", lwd = 2)

# y=tan(x)
lines( seq(0, 10, length.out = 100), ifelse(abs(tan(seq(0, 10, length.out = 100))) > 10, NA, tan(seq(0, 10, length.out = 100))), col = "red", lwd = 2 )
# y=tan(x)+4
lines( seq(0, 10, length.out = 2000), ifelse(abs(tan(seq(0, 10, length.out = 2000)) + 4) > 10, NA, tan(seq(0, 10, length.out = 2000)) + 4), col = "red", lwd = 2 )


# グリッド線を追加
grid()
