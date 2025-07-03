# シミュレーション設定
param_list <- list(
  annual_return         = 0.06,
  annual_volatility     = 0.15,
  starting_balance      = 1e6,
  deposit_before_change = 5e4,
  change_after_years    = 10,
  deposit_after_change  = 8e4,
  total_years           = 20,
  n_simulations         = 2000
)

# ファイル出力先
home_dir    <- path.expand("~")
file_name   <- "accumulation_simulation.png"
output_path <- file.path(home_dir, file_name)

# パッケージインストール、読み込み
options(repos = c(CRAN = "https://cloud.r-project.org"))
pkgs <- c("ggplot2", "scales", "ragg")
for (p in pkgs) {
  if (!p %in% installed.packages()[, "Package"]) install.packages(p)
}
lapply(pkgs, require, character.only = TRUE)

# シミュレーション
simulate_portfolio <- function(params) {
  with(params, {
    # 月次リターン・ボラティリティ
    mu_m     <- annual_return   / 12
    sigma_m  <- annual_volatility / sqrt(12)
    total_m  <- total_years     * 12
    change_m <- change_after_years * 12

    # シミュレーション結果格納行列
    mat <- matrix(NA, nrow = total_m + 1, ncol = n_simulations)
    mat[1, ] <- starting_balance

    # モンテカルロループ
    for (sim in seq_len(n_simulations)) {
      for (t in 2:(total_m + 1)) {
        depo <- if ((t - 1) <= change_m) deposit_before_change else deposit_after_change
        r    <- rnorm(1, mu_m, sigma_m)
        mat[t, sim] <- mat[t - 1, sim] * (1 + r) + depo
      }
    }
    mat
  })
}

# シミュレーション実行
bal_mat <- simulate_portfolio(param_list)

# 年次サマリー作成
years_seq <- 0:param_list$total_years
rows      <- years_seq * 12 + 1
summary_df <- data.frame(
  year   = years_seq,
  pct_5  = apply(bal_mat[rows, ], 1, quantile, probs = 0.05),
  pct_50 = apply(bal_mat[rows, ], 1, quantile, probs = 0.50),
  pct_95 = apply(bal_mat[rows, ], 1, quantile, probs = 0.95)
)

# 画像作成
ragg::agg_png(
  filename = output_path,
  width    = 1000,  # 横
  height   = 1000,  # 縦
  units    = "px"
)

unit   <- 10000000  # 1千万
y_max  <- ceiling(max(summary_df$pct_95) / unit) * unit

ggplot(summary_df, aes(x = year)) +
  geom_ribbon(
    aes(ymin = pct_5, ymax = pct_95),
    fill      = "#3366CC",
    alpha     = 0.3,
    linewidth = 0
  ) +
  geom_line(
    aes(y = pct_50),
    color     = "#3366CC",
    linewidth = 1.2
  ) +
  scale_x_continuous(
    breaks = seq(0, param_list$total_years, by = 2),
    labels = ~ paste0(.x)
  ) +
  scale_y_continuous(
    breaks = seq(0, y_max, by = unit),
    labels = function(x) paste0(round(x / unit, 1), "千万"),
    expand = expansion(mult = c(0, 0.02))
  ) +
  labs(
    title   = paste0("Monte Carlo Simulation: Deposit Change at ", param_list$change_after_years, " Years"),
    x       = "年",
    y       = "円",
    caption = "バンド：5%–95%パーセンタイル、線：50%パーセンタイル"
  ) +
  theme_minimal(base_size = 16) +
  theme(
    axis.title         = element_text(size = 20),
    axis.text          = element_text(size = 18, color = "grey20"),
    panel.grid.major.y = element_line(linewidth = 0.6, color = "grey80"),
    panel.grid.major.x = element_line(linewidth = 0.6, color = "grey90"),
    panel.grid.minor   = element_line(linewidth = 0.3, color = "grey95")
  )
dev.off()
