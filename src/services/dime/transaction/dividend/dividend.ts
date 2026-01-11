import type { Transaction } from "../transaction";

export type DividentType = "Tax" | "Income";
export interface Dividend extends Transaction {
  symbol: string;
  amount: number;
  type: "Tax" | "Income Dividend";
  remark?: string;
}
