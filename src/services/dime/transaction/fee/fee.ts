import type { Transaction } from "../transaction";
export interface Fee extends Transaction {
  type: "TAF";
  amount: number;
  remark: string;
}
