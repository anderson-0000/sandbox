options(repos = c(CRAN = "https://cloud.r-project.org"))

px <- 10000
home_dir <- path.expand("~")
file_name <- "mandelbrot_high_resolution.png"
output_path <- file.path(home_dir, file_name)

# パッケージリスト
required_packages <- c("mandelbrot", "ragg")

for (package_name in required_packages) {
    if (!(package_name %in% installed.packages())) {
        install.packages(package_name)
    }
}

# パッケージのロード
lapply(required_packages, require, character.only = TRUE)

ragg::agg_png(
    filename = output_path,
    width    = px,
    height   = px,
    units    = "px",
)

# マンデルブロ集合を計算
mb <- mandelbrot(
    resolution = px,
    iterations = 500 
)

# 結果をプロット
plot(mb)
