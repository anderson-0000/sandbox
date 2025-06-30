## r 起動後、以下コマンドで実行
# source("mandelbrot_set_2.R")

options(repos = c(CRAN = "https://cloud.r-project.org"))

# パッケージリスト
required_packages <- c("mandelbrot", "ragg", "devtools", "fractal")

for (package_name in required_packages) {
    if (!(package_name %in% installed.packages())) {
        install.packages(package_name)
    }
}

devtools::install_github("mariodosreis/fractal")

# パッケージのロード
lapply(required_packages, require, character.only = TRUE)

mb <- mandelbrot(nx = 2000, ny = 2000, iter = 300)
image(mb, col = c(heat.colors(49), "black"), 
      xlab = "Re", ylab = "Im", main = "Mandelbrot Set")

cat("拡大したい領域2点をクリックしズームします。\n")
cat("中断するには Ctrl-C を押してください。\n\n")

i <- 1
repeat {
  tryCatch({
    cat(sprintf("Zoom #%d: 拡大したい領域2点をクリック\n", i))
    zoom(
      fun   = "mandelbrot",
      col   = c(heat.colors(49), "black"),
      iter  = 300 + i*20,
      nx    = 2000, ny = 2000,
      plot  = TRUE
    )
    i <- i + 1
  }, interrupt = function(ex) {
    cat("\nユーザ割り込みを検知しました。ループを終了します。\n")
    break
  })
}
