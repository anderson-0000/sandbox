options(repos = c(CRAN = "https://cloud.r-project.org"))


# パッケージリスト
required_packages <- c("mandelbrot", "ragg")

for (package_name in required_packages) {
    if (!(package_name %in% installed.packages())) {
        install.packages(package_name)
    }
}

# パッケージのロード
lapply(required_packages, require, character.only = TRUE)

ragg::agg_png("output_ragg.png", width = 8, height = 6, units = "in", res = 300)

# マンデルブロ集合を計算
mb <- mandelbrot()

# 結果をプロット
plot(mb)
