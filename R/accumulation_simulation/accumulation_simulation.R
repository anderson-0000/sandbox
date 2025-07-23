# R --vanilla -f accumulation_simulation.R

# シミュレーション設定（金融商品1）
param_list <- list(
  yearly_return_percent     = 9.45,    # 年間リターン (%)
  yearly_volatility_percent = 30.21,   # 年間リスク (%)
  initial_balance           = 5000000, # 初期資産額
  deposit_first_phase       = 50000,   # 月次入金額
  years_until_change        = 10,      # 入金額変更までの年数
  deposit_second_phase      = 30000,   # 入金額変更後の月次入金額
  simulation_years          = 20,      # シミュレーション期間
  num_simulations           = 5000,    # シミュレーション回数
  start_age                 = 32       # シミュレーション開始年齢
)

# 金融商品2のパラメータ
param_list2 <- list(
  yearly_return_percent     = 9.45,    # 年間リターン (%)
  yearly_volatility_percent = 30.21,   # 年間リスク (%)
  initial_balance           = 5000000, # 初期資産額
  deposit_first_phase       = 50000,   # 月次入金額
  years_until_change        = 10,      # 入金額変更までの年数
  deposit_second_phase      = 30000,   # 入金額変更後の月次入金額
  simulation_years          = param_list$simulation_years,
  num_simulations           = param_list$num_simulations,
  start_age                 = param_list$start_age
)

# 出力先
home_dir    <- path.expand("~")
file_name   <- "accumulation_simulation.png"
output_path <- file.path(home_dir, file_name)

# パッケージ
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

# 年次サマリー
years_seq    <- 0:param_list$simulation_years
summary_rows <- years_seq * 12 + 1
yearly_summary <- data.frame(
  year   = years_seq,
  pct_30 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.30),
  pct_50 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.50),
  pct_70 = apply(bal_balance_matrix[summary_rows, ], 1, quantile, probs = 0.70)
)
yearly_summary$age <- param_list$start_age + yearly_summary$year

# プロット出力
ragg::agg_png(filename = output_path, width = 1000, height = 1000, units = "px")

unit  <- 10000000
y_max <- ceiling(max(yearly_summary$pct_70) / unit) * unit

ggplot(yearly_summary, aes(x = age)) +
  geom_ribbon(aes(ymin = pct_30, ymax = pct_70),
              fill = "#3366CC", alpha = 0.3, linewidth = 0) +
  geom_line(aes(y = pct_50),
            color = "#3366CC", linewidth = 1.2) +
  scale_x_continuous(
    breaks       = seq(param_list$start_age, param_list$start_age + param_list$simulation_years, by = 5),
    minor_breaks = seq(param_list$start_age, param_list$start_age + param_list$simulation_years, by = 1),
    labels       = function(x) {
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
      ifelse(x == 0, "0",
             ifelse(x %% (unit * 10) == 0,
                    paste0(round(x / (unit * 10)), "億"),
                    paste0(round(x / unit, 1), "千万")))
    },
    limits = c(0, y_max),
    expand = expansion(mult = c(0, 0.02))
  ) +
  labs(
    title   = paste0("Monte Carlo Simulation: Deposit Change at ",
                     param_list$years_until_change, " Years"),
    x       = "年齢／西暦",
    y       = "円",
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
    panel.grid.major.y  = element_line(linewidth = 0.6, color = "grey80"),
    panel.grid.major.x  = element_line(linewidth = 0.6, color = "grey90"),
    panel.grid.minor.x  = element_line(linewidth = 0.3, color = "grey90"),
    panel.grid.minor.y  = element_blank(),
    plot.caption        = element_text(hjust = 1, size = 12)
  )

dev.off()

