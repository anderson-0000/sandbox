# R --vanilla -f accumulation_simulation.R

# パラメータファイルをがあれば読み込む
param_file <- file.path(path.expand("~"), "parameter", "accumulation_simulation_paramater.R")
if (file.exists(param_file)) {
  source(param_file)   # param_list, param_list2, uninvested_cash を上書き
} else {
  # 金融商品1（デフォルト）
  param_list <- list(
    yearly_return_percent     = 8.50,    # 年間リターン (%)
    yearly_volatility_percent = 25.00,   # 年間リスク (%)
    initial_balance           = 3000000, # 初期資産額
    deposit_first_phase       = 60000,   # 月次入金額
    years_until_change        = 12,      # 入金額変更までの年数
    deposit_second_phase      = 40000,   # 変更後の月次入金額
    simulation_years          = 25,      # シミュレーション期間（年）
    num_simulations           = 10000,   # 試行回数
    start_age                 = 30       # 開始年齢
  )

  # 金融商品2（デフォルト）
  param_list2 <- list(
    yearly_return_percent     = 9.45,    # 年間リターン (%)
    yearly_volatility_percent = 30.21,   # 年間リスク (%)
    initial_balance           = 5000000, # 初期資産額
    deposit_first_phase       = 50000,   # 月次入金額
    years_until_change        = 10,      # 入金額変更までの年数
    deposit_second_phase      = 30000,   # 変更後の月次入金額
    simulation_years          = param_list$simulation_years,
    num_simulations           = param_list$num_simulations,
    start_age                 = param_list$start_age
  )

  # 投資していない貯金額
  uninvested_cash <- 1000000
}

if (!exists("uninvested_cash")) {
  uninvested_cash <- 1000000
}

# 出力先
home_dir    <- path.expand("~")
file_name   <- "accumulation_simulation.png"
output_path <- file.path(home_dir, file_name)

# パッケージ読み込み
options(repos = c(CRAN = "https://cloud.r-project.org"))
pkgs <- c("ggplot2", "scales", "ragg")
for (p in pkgs) {
  if (!p %in% installed.packages()[, "Package"]) install.packages(p)
}
lapply(pkgs, require, character.only = TRUE)

# シミュレーション関数
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

# 二商品を別々にシミュレーションして合算
bal1 <- simulate_portfolio(param_list)
bal2 <- simulate_portfolio(param_list2)
bal_balance_matrix <- bal1 + bal2

# 年次サマリー作成（uninvested_cash を加算）
years_seq    <- 0:param_list$simulation_years
summary_rows <- years_seq * 12 + 1
yearly_summary <- data.frame(
  year   = years_seq,
  pct_30 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.30) + uninvested_cash,
  pct_50 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.50) + uninvested_cash,
  pct_70 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.70) + uninvested_cash
)
yearly_summary$age <- param_list$start_age + yearly_summary$year

# 現在の年月と各年1月をメジャーティックに使う設定
current_year        <- as.integer(format(Sys.Date(), "%Y"))
current_month       <- as.integer(format(Sys.Date(), "%m"))
months_to_next_jan  <- 12 - current_month + 1
first_jan_offset    <- months_to_next_jan / 12
year_count          <- param_list$simulation_years
break_offsets       <- c(0, first_jan_offset + seq(0, year_count - 1))
x_breaks            <- param_list$start_age + break_offsets
age_labels          <- param_list$start_age + floor(break_offsets)
date_labels         <- c(
  paste0(current_year,      "年", sprintf("%02d", current_month),       "月"),
  paste0(current_year + seq_len(year_count), "年01月")
)
x_labels            <- paste0(age_labels, "\n", date_labels)

# プロット出力
ragg::agg_png(filename = output_path, width = 1000, height = 1000, units = "px")

unit  <- 10000000
y_max <- ceiling(max(yearly_summary$pct_70) / unit) * unit

ggplot(yearly_summary, aes(x = age)) +
  geom_ribbon(aes(ymin = pct_30, ymax = pct_70),
              fill = "#3366CC", alpha = 0.3, linewidth = 0) +
  geom_line(aes(y = pct_50), color = "#3366CC", linewidth = 1.2) +
  scale_x_continuous(
    breaks       = x_breaks,
    minor_breaks = NULL,
    labels       = x_labels,
    expand       = expansion(mult = c(0, 0))
  ) +
  scale_y_continuous(
    # メモリは変更せず、ラベルだけ調整
    breaks = local({
      b1 <- seq(0, min(y_max, unit * 10), by = unit)
      b2 <- if (y_max > unit * 10) seq(unit * 10, y_max, by = unit * 10) else numeric(0)
      c(b1, b2)
    }),
    labels = function(x) {
      labs <- rep("", length(x))
      labs[x == 0] <- "0"
      # 1億円未満：2千万ごとに「x.x千万」
      idx1 <- x > 0 & x < unit * 10 & (x %% (unit * 2) == 0)
      labs[idx1] <- paste0(round(x[idx1] / unit, 1), "千万")
      # 1億円以上：1億ごとに「x億」
      idx2 <- x >= unit * 10 & (x %% (unit * 10) == 0)
      labs[idx2] <- paste0(round(x[idx2] / (unit * 10)), "億")
      labs
    },
    limits = c(0, y_max),
    expand = expansion(mult = c(0, 0.02))
  ) +
  labs(
    title   = paste0("資産シミュレーション"),
    x       = "年齢／西暦・月",
    y       = "資産（円）",
    caption = paste0(
      "バンド：30%–70%パーセンタイル、線：50%パーセンタイル\n",
      "実行日時: ", format(Sys.time(), "%Y-%m-%d %H:%M:%S")
    )
  ) +
  theme_minimal(base_size = 16) +
  theme(
    plot.margin         = margin(t = 10, r = 30, b = 10, l = 10),
    axis.title          = element_text(size = 20),
    axis.text           = element_text(size = 18, color = "grey20"),
    axis.text.x         = element_text(angle = 45, hjust = 1),
    panel.grid.major.y  = element_line(linewidth = 0.6, color = "grey80"),
    panel.grid.major.x  = element_line(linewidth = 0.6, color = "grey90"),
    panel.grid.minor.x  = element_line(linewidth = 0.3, color = "grey90"),
    panel.grid.minor.y  = element_blank(),
    plot.caption        = element_text(hjust = 1, size = 12)
  )

dev.off()

