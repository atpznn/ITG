import type { Transaction } from "../transaction";

export interface Stock extends Transaction {
  symbol: string;
  price: number;
  shares: number;
  amount: number;
  type: "Sell" | "Buy";
  remark: string;
}
