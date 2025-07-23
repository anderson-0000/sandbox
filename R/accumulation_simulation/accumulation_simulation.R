# R --vanilla -f accumulation_simulation.R

# シミュレーション設定
param_list <- list(
  yearly_return_percent     = 9.45,    # 年間リターン (%)
  yearly_volatility_percent = 30.21,   # 年間リスク (%)
  initial_balance           = 5000000,  # 初期資産額
  deposit_first_phase       = 50000,   # 月次入金額
  years_until_change        = 10,      # 入金額変更までの年数
  deposit_second_phase      = 30000,   # 入金額変更後の月次入金額
  simulation_years          = 20,      # シミュレーション期間
  num_simulations           = 5000,    # シミュレーション回数
  start_age                 = 32      # シミュレーション開始年齢
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
    avg_monthly_return <- (yearly_return_percent / 100) / 12
    monthly_return_sd  <- (yearly_volatility_percent / 100) / sqrt(12)
    total_months       <- simulation_years * 12
    change_month       <- years_until_change * 12
    balance_matrix     <- matrix(NA, nrow = total_months + 1, ncol = num_simulations)
    balance_matrix[1, ] <- initial_balance

    for (sim in seq_len(num_simulations)) {
      for (t in 2:(total_months + 1)) {
        depo <- if ((t - 1) <= change_month) deposit_first_phase else deposit_second_phase
        r    <- rnorm(1, avg_monthly_return, monthly_return_sd)
        balance_matrix[t, sim] <- balance_matrix[t - 1, sim] * (1 + r) + depo
      }
    }
    balance_matrix
  })
}

# シミュレーション実行
bal_balance_matrix <- simulate_portfolio(param_list)

# 年次サマリー作成
years_seq     <- 0:param_list$simulation_years
summary_rows  <- years_seq * 12 + 1
yearly_summary <- data.frame(
  year   = years_seq,
  pct_30 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.30),
  pct_50 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.50),
  pct_70 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.70)
)

# 年齢列を追加
yearly_summary$age <- param_list$start_age + yearly_summary$year

# 画像作成
ragg::agg_png(
  filename = output_path,
  width    = 1000,
  height   = 1000,
  units    = "px"
)

unit  <- 10000000  # 1千万
y_max <- ceiling(max(yearly_summary$pct_70) / unit) * unit

ggplot(yearly_summary, aes(x = age)) +
  geom_ribbon(
    aes(ymin = pct_30, ymax = pct_70),
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
    # 年齢メジャーブレイク：5歳ごとのみ
    breaks = seq(
      from = param_list$start_age,
      to   = param_list$start_age + param_list$simulation_years,
      by   = 5
    ),
    # 年齢マイナーブレイク：1年ごと（補助線用）
    minor_breaks = seq(
      from = param_list$start_age,
      to   = param_list$start_age + param_list$simulation_years,
      by   = 1
    ),
    # 年齢／西暦を二段で表示
    labels = function(x) {
      yrs      <- x - param_list$start_age
      cal_year <- as.integer(format(Sys.Date(), "%Y")) + yrs
      paste0(x, "\n", cal_year, "年")
    },
    expand = expansion(add = c(0, 1))
  ) +
  scale_y_continuous(
    breaks = local({
      b1 <- seq(0, min(y_max, unit * 10), by = unit)
      b2 <- if (y_max > unit * 10) seq(unit * 10, y_max, by = unit * 10) else numeric(0)
      c(b1, b2)
    }),
    labels = function(x) {
      ifelse(
        x == 0, "0",
        ifelse(
          x %% (unit * 10) == 0,
          paste0(round(x / (unit * 10)), "億"),
          paste0(round(x / unit, 1), "千万")
        )
      )
    },
    limits = c(0, y_max),  # 0円からスタート
    expand = expansion(mult = c(0, 0.02))
  ) +
  labs(
    title   = paste0("Monte Carlo Simulation: Deposit Change at ",
                     param_list$years_until_change, " Years"),
    x       = "年齢／西暦",
    y       = "円",
    caption = "バンド：30%–70%パーセンタイル、線：50%パーセンタイル"
  ) +
  theme_minimal(base_size = 16) +
  theme(
    plot.margin         = margin(t = 10, r = 30, b = 10, l = 10),
    axis.title          = element_text(size = 20),
    axis.text           = element_text(size = 18, color = "grey20"),
    panel.grid.major.y  = element_line(linewidth = 0.6, color = "grey80"),
    panel.grid.major.x  = element_line(linewidth = 0.6, color = "grey90"),
    panel.grid.minor.x  = element_line(linewidth = 0.3, color = "grey90"),  # 1年ごとの縦線
    panel.grid.minor.y  = element_blank()
  )

dev.off()

