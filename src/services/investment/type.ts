export interface BaseInvestment {
  shares: number;
  price: number;
  total: number;
  type: string;
  createdAt: Date;
  id: string;
  executedAt: Date;
}
