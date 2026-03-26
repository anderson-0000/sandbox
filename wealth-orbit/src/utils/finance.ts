export interface FinanceParams {
  initialInvestment: number;
  monthlySavings: number;
  annualRate: number; // パーセンテージ (例: 5)
  targetAmount: number;
}

/**
 * 指定した年数後の資産額を計算する (月複利)
 */
export const calculateFutureValue = (params: FinanceParams, years: number): number => {
  const { initialInvestment, monthlySavings, annualRate } = params;
  const monthlyRate = annualRate / 100 / 12;
  const totalMonths = years * 12;

  if (monthlyRate === 0) {
    return initialInvestment + monthlySavings * totalMonths;
  }

  // 複利計算: FV = P(1+r)^n + PMT * [((1+r)^n - 1) / r]
  const futureValue = 
    initialInvestment * Math.pow(1 + monthlyRate, totalMonths) +
    monthlySavings * ((Math.pow(1 + monthlyRate, totalMonths) - 1) / monthlyRate);

  return futureValue;
};

/**
 * 目標金額に到達するまでの年数を計算する
 */
export const calculateYearsToTarget = (params: FinanceParams): number => {
  const { initialInvestment, monthlySavings, annualRate, targetAmount } = params;
  const r = annualRate / 100 / 12;
  const P = initialInvestment;
  const PMT = monthlySavings;
  const FV = targetAmount;

  if (P >= FV) return 0;

  if (r === 0) {
    if (PMT <= 0) return Infinity;
    return (FV - P) / (PMT * 12);
  }

  // FV = P(1+r)^n + PMT * [((1+r)^n - 1) / r]
  // FV * r = P * r * (1+r)^n + PMT * (1+r)^n - PMT
  // FV * r + PMT = (P * r + PMT) * (1+r)^n
  // (1+r)^n = (FV * r + PMT) / (P * r + PMT)
  // n = log((FV * r + PMT) / (P * r + PMT)) / log(1+r)
  const numerator = FV * r + PMT;
  const denominator = P * r + PMT;

  if (denominator <= 0) return Infinity;

  const n = Math.log(numerator / denominator) / Math.log(1 + r);
  return n / 12;
};
